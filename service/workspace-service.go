package service

import (
	"context"
	mapping "data-handler/service/mapping"
	"data-handler/service/security"
	"data-handler/service/system"
	"data-handler/stub"
	"data-handler/stub/model"
	log "github.com/sirupsen/logrus"
)

type WorkSpaceService interface {
	stub.WorkSpaceServiceServer
	InjectRecordService(service RecordService)
	InjectAuthenticationService(service AuthenticationService)
	Init(data *model.InitData)
	InjectResourceService(service ResourceService)
}

type workSpaceService struct {
	stub.WorkSpaceServiceServer
	recordService         RecordService
	authenticationService AuthenticationService
	serviceName           string
	resourceService       ResourceService
}

func (u *workSpaceService) InjectResourceService(service ResourceService) {
	u.resourceService = service
}

func (u *workSpaceService) InjectAuthenticationService(service AuthenticationService) {
	u.authenticationService = service
}

func (u *workSpaceService) InjectRecordService(service RecordService) {
	u.recordService = service
}

func (u *workSpaceService) Create(ctx context.Context, request *stub.CreateWorkSpaceRequest) (*stub.CreateWorkSpaceResponse, error) {
	// insert records via resource service
	records := mapping.MapToRecord(request.WorkSpaces, mapping.WorkSpaceToRecord)
	systemCtx := security.WithSystemContext(ctx)

	result, err := u.recordService.Create(systemCtx, &stub.CreateRecordRequest{
		Token:   request.Token,
		Records: records,
	})

	if err != nil {
		return nil, err
	}

	return &stub.CreateWorkSpaceResponse{
		WorkSpaces: mapping.MapFromRecord(result.Records, mapping.WorkSpaceFromRecord),
		Error:      result.Error,
	}, err
}

func (u *workSpaceService) Update(ctx context.Context, request *stub.UpdateWorkSpaceRequest) (*stub.UpdateWorkSpaceResponse, error) {
	// insert records via resource service
	records := mapping.MapToRecord(request.WorkSpaces, mapping.WorkSpaceToRecord)
	systemCtx := security.WithSystemContext(ctx)
	result, err := u.recordService.Update(systemCtx, &stub.UpdateRecordRequest{
		Token:   request.Token,
		Records: records,
	})

	if err != nil {
		return nil, err
	}

	return &stub.UpdateWorkSpaceResponse{
		WorkSpaces: mapping.MapFromRecord(result.Records, mapping.WorkSpaceFromRecord),
		Error:      result.Error,
	}, err
}

func (u *workSpaceService) Delete(ctx context.Context, request *stub.DeleteWorkSpaceRequest) (*stub.DeleteWorkSpaceResponse, error) {
	systemCtx := security.WithSystemContext(ctx)

	record, err := u.recordService.Delete(systemCtx, &stub.DeleteRecordRequest{
		Token:    request.Token,
		Resource: system.WorkSpaceResource.Name,
		Ids:      request.Ids,
	})

	if err != nil {
		return nil, err
	}

	return &stub.DeleteWorkSpaceResponse{
		Error: record.Error,
	}, nil
}

func (u *workSpaceService) Get(ctx context.Context, request *stub.GetWorkSpaceRequest) (*stub.GetWorkSpaceResponse, error) {
	systemCtx := security.WithSystemContext(ctx)
	record, err := u.recordService.Get(systemCtx, &stub.GetRecordRequest{
		Token:    request.Token,
		Resource: system.WorkSpaceResource.Name,
		Id:       request.Id,
	})

	if err != nil {
		return nil, err
	}

	return &stub.GetWorkSpaceResponse{
		WorkSpace: mapping.WorkSpaceFromRecord(record.Record),
		Error:     record.Error,
	}, nil
}

func (u *workSpaceService) List(ctx context.Context, request *stub.ListWorkSpaceRequest) (*stub.ListWorkSpaceResponse, error) {
	systemCtx := security.WithSystemContext(ctx)
	result, err := u.recordService.List(systemCtx, &stub.ListRecordRequest{
		Resource: system.WorkSpaceResource.Name,
		Token:    request.Token,
	})

	if err != nil {
		return nil, err
	}

	return &stub.ListWorkSpaceResponse{
		Content: mapping.MapFromRecord(result.Content, mapping.WorkSpaceFromRecord),
		Error:   result.Error,
	}, err
}

func (d *workSpaceService) Init(data *model.InitData) {
	d.resourceService.InitResource(system.WorkSpaceResource)

	if len(data.InitWorkSpaces) > 0 {
		res, err := d.recordService.Create(security.SystemContext, &stub.CreateRecordRequest{
			Records:        mapping.MapToRecord(data.InitWorkSpaces, mapping.WorkSpaceToRecord),
			IgnoreIfExists: true,
		})

		if err != nil || res.Error != nil {
			log.Error(err, res.Error)
		}
	}
}

func NewWorkSpaceService() WorkSpaceService {
	return &workSpaceService{
		serviceName: "WorkSpaceService",
	}
}
