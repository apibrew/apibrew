package pkg

import (
	"context"
	model2 "github.com/apibrew/apibrew/modules/storage/pkg/model"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"time"
)

type module struct {
	container           service.Container
	backendEventHandler backend_event_handler.BackendEventHandler
	api                 api.Interface
	minioClient         *minio.Client
	fileRepository      api.Repository[*model2.File]
	fileProcessor       *fileProcessor
}

func (m *module) Init() {
	config := m.container.GetAppConfig().Modules["cloud-storage"]

	if config == nil || config.Options == nil || len(config.Options) == 0 {
		log.Println("cloud-storage module is not enabled or configured properly")
		return
	}

	if err := m.setupBucket(config.Options); err != nil {
		log.Fatal(err)
	}

	m.ensureNamespace()
	m.ensureResources()

	fileRepository := api.NewRepository[*model2.File](m.api, model2.FileMapperInstance)
	m.fileRepository = fileRepository
	m.fileProcessor = &fileProcessor{
		api:            m.api,
		client:         m.minioClient,
		bucketName:     config.Options["bucketName"],
		fileRepository: fileRepository,
	}

	if err := RegisterResourceProcessor[*model2.File](
		"storage-file-listener",
		m.fileProcessor,
		m.backendEventHandler,
		m.container,
		model2.FileResource,
	); err != nil {
		log.Fatal(err)
	}

	go m.presignUrls()
}

func (m *module) ensureNamespace() {
	_, err := m.container.GetRecordService().Apply(util.SystemContext, service.RecordUpdateParams{
		Namespace: resources.NamespaceResource.Namespace,
		Resource:  resources.NamespaceResource.Name,
		Records: []*model.Record{
			{
				Properties: map[string]*structpb.Value{
					"name": structpb.NewStringValue("storage"),
				},
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (m *module) ensureResources() {
	var list = []*model.Resource{
		model2.FileResource,
	}

	for _, resource := range list {
		existingResource, err := m.container.GetResourceService().GetResourceByName(util.SystemContext, resource.Namespace, resource.Name)

		if err == nil {
			resource.Id = existingResource.Id
			err = m.container.GetResourceService().Update(util.SystemContext, resource, true, true)

			if err != nil {
				log.Fatal(err)
			}
		} else if errors.ResourceNotFoundError.Is(err) {
			_, err = m.container.GetResourceService().Create(util.SystemContext, resource, true, true)

			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}
}

func (m *module) setupBucket(options map[string]string) error {
	minioClient, err := minio.New("storage.apibrew.io:443", &minio.Options{
		Creds:  credentials.NewStaticV4(options["accessKey"], options["secretKey"], ""),
		Secure: true,
	})
	if err != nil {
		return err
	}

	m.minioClient = minioClient

	bucketName := options["bucketName"]
	location := "us-east-1"

	exists, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		log.Fatalln("Error checking if bucket exists:", err)
	}

	if exists {
		log.Printf("Bucket %s exists\n", bucketName)
		return nil
	}

	err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully created bucket %s\n", bucketName)

	return nil
}

func (m *module) presignUrls() {
	for {
		log.Println("Begin Presigning URLs")
		var files, err = m.fileRepository.List(util.SystemContext, api.ListParams{
			Type:  "storage/File",
			Limit: 100000,
		})

		if err != nil {
			log.Error(err)
		}

		for _, file := range files.Content {
			err = m.fileProcessor.presignUrls(file)

			if err != nil {
				log.Error(err)
			}

			_, err = m.fileRepository.Update(util.SystemContext, file)

			if err != nil {
				log.Error(err)
			}
		}

		log.Println("End Presigning URLs")

		time.Sleep(12 * time.Hour)
	}
}

func NewModule(container service.Container) service.Module {
	a := api.NewInterface(container)

	backendEventHandler := container.GetBackendEventHandler().(backend_event_handler.BackendEventHandler)
	return &module{container: container,
		api:                 a,
		backendEventHandler: backendEventHandler}
}
