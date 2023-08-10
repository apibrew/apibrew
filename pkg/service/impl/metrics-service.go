package impl

import (
	"context"
	"github.com/InfluxCommunity/influxdb3-go/influx"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/hashicorp/go-metrics"
	log "github.com/sirupsen/logrus"
	"time"
)

type metricService struct {
	recordService   service.RecordService
	resourceService service.ResourceService
}

type metricsServiceInfluxEncoder struct {
	config *model.InfluxDBConfig
	client *influx.Client
}

func (m *metricsServiceInfluxEncoder) Init() {
	// Create a new client using an InfluxDB server base URL and an authentication token
	client, err := influx.New(influx.Configs{
		HostURL:      m.config.HostUrl,
		AuthToken:    m.config.Token,
		Organization: "apibrew",
	})

	if err != nil {
		log.Fatal(err)
	}

	m.client = client
}

func (m *metricsServiceInfluxEncoder) Encode(i interface{}) error {
	if metricsSummary, ok := i.(metrics.MetricsSummary); ok {
		var points []*influx.Point

		// Timestamp
		timestamp, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", metricsSummary.Timestamp)
		if err != nil {
			log.Fatalf("Failed to parse timestamp: %s", err)
		}

		// Gauges
		for _, gauge := range metricsSummary.Gauges {
			point := influx.NewPoint(gauge.Name,
				gauge.DisplayLabels,
				map[string]interface{}{"value": gauge.Value},
				timestamp)
			points = append(points, point)
		}

		// Precision Gauges
		for _, pgauge := range metricsSummary.PrecisionGauges {
			point := influx.NewPoint(pgauge.Name,
				pgauge.DisplayLabels,
				map[string]interface{}{"value": pgauge.Value},
				timestamp)
			points = append(points, point)
		}

		// Points
		for _, pointValue := range metricsSummary.Points {
			point := influx.NewPoint(pointValue.Name,
				nil,
				map[string]interface{}{"points": pointValue.Points},
				timestamp)
			points = append(points, point)
		}

		// Counters
		for _, counter := range metricsSummary.Counters {
			point := influx.NewPoint(counter.Name,
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
			points = append(points, point)
		}

		// Samples
		for _, sample := range metricsSummary.Samples {
			point := influx.NewPoint(sample.Name,
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
			points = append(points, point)
		}

		err = m.client.WritePoints(context.TODO(), m.config.Database, points...)

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
			encoder := &metricsServiceInfluxEncoder{config: config.Metrics.Influxdb}

			encoder.Init()

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
