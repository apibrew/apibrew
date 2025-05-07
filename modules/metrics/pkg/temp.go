package pkg

import "time"

type MetricsOperation string

const (
	MetricsOperationRead   MetricsOperation = "read"
	MetricsOperationWrite  MetricsOperation = "write"
	MetricsOperationDelete MetricsOperation = "delete"
)

type MetricsInterval string

const (
	MetricsIntervalMinute MetricsInterval = "minute"
	MetricsIntervalHour   MetricsInterval = "hour"
	MetricsIntervalDay    MetricsInterval = "day"
	MetricsIntervalWeek   MetricsInterval = "week"
	MetricsIntervalMonth  MetricsInterval = "month"
)

type MetricsRequest struct {
	Namespace *string           `json:"namespace"`
	Resource  *string           `json:"resource"`
	Operation *MetricsOperation `json:"operation"`
	Interval  *MetricsInterval  `json:"interval"`
	From      *time.Time
	To        *time.Time
}

type MetricsResponseItem struct {
	Namespace string           `json:"namespace"`
	Resource  string           `json:"resource"`
	Interval  MetricsInterval  `json:"interval"`
	Operation MetricsOperation `json:"operation"`
	Time      time.Time        `json:"time"`
	Count     uint64           `json:"count"`
}
