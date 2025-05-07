package pkg

import (
	"context"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/hashicorp/go-metrics"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	api2 "github.com/influxdata/influxdb-client-go/v2/api"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
	"time"
)

type module struct {
	container           service.Container
	backendEventHandler backend_event_handler.BackendEventHandler
	api                 api.Interface
	disabled            bool
	options             map[string]string
	serviceId           string
	influxQueryApi      api2.QueryAPI
}

func (m module) Init() {
	if m.disabled {
		return
	}
	log.Println("Initializing module metrics")
	//m.ensureNamespace()
	//m.ensureResources()

	//oauth2ConfigRepository := api.NewRepository[*model2.TestResource](m.api, model2.TestResourceMapperInstance)

	var hostUrl = m.options["hostUrl"]
	var token = m.options["token"]
	var organization = m.options["organization"]
	var bucket = m.options["bucket"]

	influxClient := influxdb2.NewClient(hostUrl, token)

	writer := influxClient.WriteAPIBlocking(organization, bucket)
	writer.EnableBatching()

	encoder := &influxEncoder{writer: writer}
	m.serviceId = m.serviceId

	m.influxQueryApi = influxClient.QueryAPI(organization)

	var retention = time.Hour * 24 * 7
	var interval = 10 * time.Second

	//if config.Metrics.Retention != nil {
	//	retention = time.Duration(*config.Metrics.Retention) * time.Millisecond
	//}
	//
	//if config.Metrics.Interval != nil {
	//	interval = time.Duration(*config.Metrics.Interval) * time.Millisecond
	//}

	var inm = metrics.NewInmemSink(interval, retention)
	_, err := metrics.NewGlobal(metrics.DefaultConfig(m.serviceId), inm)

	if err != nil {
		log.Fatal(err)
	}

	go inm.Stream(context.TODO(), encoder)

	m.registerMetricListener()
}

func (m module) GetMetrics(req MetricsRequest) ([]MetricsResponseItem, error) {
	var query = `from(bucket:"apibrew")
					|> range(start: -100m)
					|> filter(fn: (r) => r._measurement == "` + m.serviceId + `.RecordService" and r._field == "count")
					`
	it, err := m.influxQueryApi.Query(context.Background(), query)

	var result []MetricsResponseItem

	if err != nil {
		return nil, errors.InternalError.WithDetails(err.Error())
	}

	for it.Next() {
		var item MetricsResponseItem
		var rec = it.Record()
		var values = rec.Values()

		item.Namespace = values["namespace"].(string)
		item.Resource = values["resource"].(string)
		item.Operation = MetricsOperation(values["operation"].(string))
		item.Count = uint64(values["_value"].(int64))

		item.Time = values["_time"].(time.Time)

		result = append(result, item)
	}

	return result, nil
}

func (m module) ensureNamespace() {
	_, err := m.container.GetRecordService().Apply(util.SystemContext, service.RecordUpdateParams{
		Namespace: resources.NamespaceResource.Namespace,
		Resource:  resources.NamespaceResource.Name,
		Records: []*model.Record{
			{
				Properties: map[string]*structpb.Value{
					"name": structpb.NewStringValue("template"),
				},
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (m module) registerMetricListener() {
	m.backendEventHandler.RegisterHandler(backend_event_handler.Handler{
		Id:   "metrics-listener",
		Name: "metrics-listener",
		Fn: func(ctx context.Context, event *model.Event) (*model.Event, error) {
			// begin metrics
			var count = len(event.Records)
			if count == 0 {
				count = 1
			}

			metrics.IncrCounterWithLabels([]string{"RecordMetrics"}, float32(count), []metrics.Label{
				{Name: "namespace", Value: event.Resource.Namespace},
				{Name: "resource", Value: event.Resource.Name},
				{Name: "action", Value: event.Action.String()},
			})

			metrics.IncrCounterWithLabels([]string{"RecordMetricsNA"}, float32(count), []metrics.Label{
				{Name: "namespace", Value: event.Resource.Namespace},
				{Name: "action", Value: event.Action.String()},
			})

			metrics.IncrCounterWithLabels([]string{"RecordMetricsN"}, float32(count), []metrics.Label{
				{Name: "namespace", Value: event.Resource.Namespace},
			})

			metrics.IncrCounterWithLabels([]string{"RecordMetricsA"}, float32(count), []metrics.Label{
				{Name: "action", Value: event.Action.String()},
			})

			metrics.IncrCounterWithLabels([]string{"RecordMetricsT"}, float32(count), []metrics.Label{})

			// end metrics
			return event, nil
		},
		Order:    1,
		Sync:     true,
		Internal: true,
	})
}

func NewModule(container service.Container) service.Module {
	a := api.NewInterface(container)

	var config = container.GetAppConfig().Modules["metrics"]

	if config == nil {
		config = &model.ModuleConfig{
			Disabled: true,
		}
	}

	backendEventHandler := container.GetBackendEventHandler().(backend_event_handler.BackendEventHandler)
	return &module{container: container,
		api:                 a,
		disabled:            config.Disabled,
		options:             config.Options,
		serviceId:           container.GetAppConfig().ServiceId,
		backendEventHandler: backendEventHandler}
}
