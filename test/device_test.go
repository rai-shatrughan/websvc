package test

import (
	"net/http"
	// "net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

var baseURL = "http://localhost:8080/api/v1/device"

func TestDeviceVersion(t *testing.T) {
	e := httpexpect.New(t, baseURL)

	obj := e.GET("/version").
		WithHeader("X-API-Key", "srkey12345").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	obj.ContainsKey("version")

	obj.Value("version").String().Equal("v1")

}

func TestDeviceId(t *testing.T) {
	e := httpexpect.New(t, baseURL)

	obj := e.GET("/6fdae6af-226d-48bd-8b61-699758137eb3").
		WithHeader("X-API-Key", "srkey12345").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	obj.Keys().Contains("id", "location", "name", "status")

}
