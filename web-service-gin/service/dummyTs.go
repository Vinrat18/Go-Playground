package service

import (
	"time"
	"web-service-gin/schema"
)

func GetDummtTSData(n int) []schema.Timeseries {
	ts := []schema.Timeseries{}

	for i := 0; i < n; i++ {
		point := schema.Timeseries{
			Id:        "TestId",
			Timestamp: time.Now().Add(-(time.Duration(i) * time.Minute)),
			Value:     10,
		}
		ts = append(ts, point)
	}

	return ts
}
