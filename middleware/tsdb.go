package middleware

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/m3db/prometheus_remote_client_golang/promremote"
)

var (
	writeurl          = "http://172.18.0.101:7201/api/v1/prom/remote/write"
	overrideUserAgent = "overrideUserAgent"
	customHeaders     = map[string]string{
		"M3-Metrics-Type": "unaggregated",
		"User-Agent":      overrideUserAgent,
	}
)

func init() {
	conf := Config{}
	conf.New()
	brokers = conf.GetStringSlice("kafka.brokers")
}

//TSDB is warpper for m3db or similar
type TSDB struct {
	URL    string
	client promremote.Client
	cfg    promremote.Config
	err    error
}

//New Intializes singleton instance of TSDB
func (ts *TSDB) New() {
	ts.URL = writeurl
	ts.cfg = promremote.NewConfig(
		promremote.WriteURLOption(ts.URL),
		promremote.HTTPClientTimeoutOption(60*time.Second),
	)

	ts.client, ts.err = promremote.NewClient(ts.cfg)
	if ts.err != nil {
		log.Fatal(fmt.Errorf("unable to construct client: %v", ts.err))
	}
}

func (ts *TSDB) Write(Labels []promremote.Label, datapoint promremote.Datapoint) {
	timeSeriesList := []promremote.TimeSeries{
		{
			Labels:    Labels,
			Datapoint: datapoint,
		},
	}

	var ctx = context.Background()
	_, err := ts.client.WriteTimeSeries(ctx, timeSeriesList, promremote.WriteOptions{Headers: customHeaders})
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("Status code: %d\n", result.StatusCode)
}
