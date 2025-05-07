package pkg

import (
	"context"
	"github.com/hashicorp/go-metrics"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	log "github.com/sirupsen/logrus"
	"time"
)

type influxEncoder struct {
	writer api.WriteAPIBlocking
}

func (m *influxEncoder) Encode(i interface{}) error {
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
