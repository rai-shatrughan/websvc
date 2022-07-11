package test

import (
	"net/http"
	// "fmt"
	"testing"
	// "encoding/json"
	"time"
	// "log"
	// "bytes"
	// "io"
	// "strings"
	// "strconv"

	"github.com/gavv/httpexpect/v2"
)

var tsBaseURL = "http://localhost:8080/api/v1/timeseries/"

type TimeSeries struct {
    Timestamp	string	`json:"timestamp"`
    Property	string	`json:"property"`
    Unit	string	`json:"unit"`
	Value	float32	`json:"value"`
}

func TestPostTimeSeries(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	tsa := []TimeSeries{ 
		{time.Now().In(loc).Format("2006-01-02T15:04:05Z"), "temperature", "celcius", 100.01},
		{time.Now().In(loc).Format("2006-01-02T15:04:05Z"), "temperature", "celcius", 105.01},
	}

	e := httpexpect.New(t, tsBaseURL)

	obj := e.PUT("/6fdae6af-226d-48bd-8b61-699758137eb3").
		WithHeader("X-API-Key", "srkey12345").
		WithHeader("Content-Type", "application/json").
		WithJSON(tsa).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	obj.Value("TimeseriesUpload").String().Equal("ok")
}

func TestGetTimeSeries(t *testing.T) {
	e := httpexpect.New(t, tsBaseURL)
	obj := e.GET("/6fdae6af-226d-48bd-8b61-699758137eb3?duration=1d").
		WithHeader("X-API-Key", "srkey12345").
		Expect().
		Status(http.StatusOK).
		JSON().
		Array()

	obj.First().Object().Value("unit").String().Contains("celcius")
}

