package impl

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/hashicorp/go-metrics"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	log "github.com/sirupsen/logrus"
	"time"
)

type metricService struct {
	recordService    service.RecordService
	resourceService  service.ResourceService
	influxBucketName string
	influxQueryApi   api.QueryAPI
	serviceId        string
}

func (m *metricService) GetMetrics(req service.MetricsRequest) ([]service.MetricsResponseItem, error) {
	var query = `from(bucket:"apibrew")
					|> range(start: -100m)
					|> filter(fn: (r) => r._measurement == "` + m.serviceId + `.RecordService" and r._field == "count")
					`
	it, err := m.influxQueryApi.Query(context.Background(), query)

	var result []service.MetricsResponseItem

	if err != nil {
		return nil, errors.InternalError.WithDetails(err.Error())
	}

	for it.Next() {
		var item service.MetricsResponseItem
		var rec = it.Record()
		var values = rec.Values()

		item.Namespace = values["namespace"].(string)
		item.Resource = values["resource"].(string)
		item.Operation = service.MetricsOperation(values["operation"].(string))
		item.Count = uint64(values["_value"].(int64))

		item.Time = values["_time"].(time.Time)

		result = append(result, item)
	}

	return result, nil
}

type metricsServiceInfluxEncoder struct {
	writer api.WriteAPIBlocking
}

func (m *metricsServiceInfluxEncoder) Encode(i interface{}) error {
	if metricsSummary, ok := i.(metrics.MetricsSummary); ok {

		// Timestamp
		timestamp, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", metricsSummary.Timestamp)
		if err != nil {
			log.Fatalf("Failed to parse timestamp: %s", err)
		}

		// Gauges
		for _, gauge := range metricsSummary.Gauges {
			point := write.NewPoint(gauge.Name,
				gauge.DisplayLabels,
				map[string]interface{}{"value": gauge.Value},
				timestamp)

			err = m.writer.WritePoint(context.TODO(), point)

			if err != nil {
				log.Fatal(err)
			}
		}

		// Precision Gauges
		for _, pgauge := range metricsSummary.PrecisionGauges {
			point := write.NewPoint(pgauge.Name,
				pgauge.DisplayLabels,
				map[string]interface{}{"value": pgauge.Value},
				timestamp)

			err = m.writer.WritePoint(context.TODO(), point)

			if err != nil {
				log.Fatal(err)
			}
		}

		// Points
		for _, pointValue := range metricsSummary.Points {
			point := write.NewPoint(pointValue.Name,
				nil,
				map[string]interface{}{"points": pointValue.Points},
				timestamp)

			err = m.writer.WritePoint(context.TODO(), point)

			if err != nil {
				log.Fatal(err)
			}
		}

		// Counters
		for _, counter := range metricsSummary.Counters {
			point := write.NewPoint(counter.Name,
				counter.DisplayLabels,
				map[string]interface{}{
					"count":  counter.Count,
					"max":    counter.Max,
					"min":    counter.Min,
					"rate":   counter.Rate,
					"sum":    counter.Sum,
					"mean":   counter.Mean,
					"stddev": counter.Stddev,
				},
				timestamp)

			err = m.writer.WritePoint(context.TODO(), point)

			if err != nil {
				log.Fatal(err)
			}
		}

		// Samples
		for _, sample := range metricsSummary.Samples {
			point := write.NewPoint(sample.Name,
				sample.DisplayLabels,
				map[string]interface{}{
					"count":  sample.Count,
					"max":    sample.Max,
					"min":    sample.Min,
					"rate":   sample.Rate,
					"sum":    sample.Sum,
					"mean":   sample.Mean,
					"stddev": sample.Stddev,
				},
				timestamp)

			err = m.writer.WritePoint(context.TODO(), point)

			if err != nil {
				log.Fatal(err)
			}
		}

		err = m.writer.Flush(context.TODO())

		if err != nil {
			log.Fatal(err)
		}

	} else {
		log.Fatal("Unknown type")
	}

	return nil
}

func (m *metricService) Init(config *model.AppConfig) {
	if config.Metrics == nil || !config.Metrics.Enabled {
		return
	}
	var retention = time.Hour * 24 * 7
	var interval = 10 * time.Second

	if config.Metrics.Retention != nil {
		retention = time.Duration(*config.Metrics.Retention) * time.Millisecond
	}

	if config.Metrics.Interval != nil {
		interval = time.Duration(*config.Metrics.Interval) * time.Millisecond
	}

	var inm = metrics.NewInmemSink(interval, retention)
	_, err := metrics.NewGlobal(metrics.DefaultConfig(config.ServiceId), inm)

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if config.Metrics.Influxdb != nil {
			influxClient := influxdb2.NewClient(config.Metrics.Influxdb.HostUrl, config.Metrics.Influxdb.Token)

			writer := influxClient.WriteAPIBlocking(config.Metrics.Influxdb.Organization, config.Metrics.Influxdb.Bucket)
			writer.EnableBatching()

			encoder := &metricsServiceInfluxEncoder{writer: writer}
			m.influxBucketName = config.Metrics.Influxdb.Bucket
			m.serviceId = config.ServiceId

			m.influxQueryApi = influxClient.QueryAPI(config.Metrics.Influxdb.Organization)

			inm.Stream(context.TODO(), encoder)
		}
	}()
}

func NewMetricService(recordService service.RecordService, resourceService service.ResourceService) service.MetricsService {
	return &metricService{
		recordService:   recordService,
		resourceService: resourceService,
	}
}
