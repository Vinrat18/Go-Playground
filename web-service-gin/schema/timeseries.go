package schema

import "time"

type Timeseries struct {
	Id        string
	Timestamp time.Time
	Value     float64
}
