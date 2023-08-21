package dbinflux

import influxdb2 "github.com/influxdata/influxdb-client-go/v2"

type InfluxStore struct {
	client *influxdb2.Client
}

func NewInfluxStore(client *influxdb2.Client) *InfluxStore {
	return &InfluxStore{
		client: client,
	}
}
