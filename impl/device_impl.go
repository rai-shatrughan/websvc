package impl

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"websvc/device/models"
	"websvc/device/restapi/operations"
	mw "websvc/middleware"
)

//Device implements device API middleware
func Device(api *operations.DeviceAPI) {
	api.Logger = log.Printf

	// kv := mw.KV{}
	sdb := mw.SessionDB{}
	sdb.New()

	sdb.Set("key1", "srkey12345")
	tok := sdb.Get("key1")
	api.Logger("Token--", tok)

	api.APIKeyAuth = func(token string) (*models.Principal, error) {
		if token == tok {
			prin := models.Principal(token)
			return &prin, nil
		}
		api.Logger("Access attempt with incorrect api key auth: %s", token)
		return nil, errors.New(401, "incorrect api key auth")
	}

	api.GetDeviceDeviceIDHandler = operations.GetDeviceDeviceIDHandlerFunc(func(params operations.GetDeviceDeviceIDParams, principal *models.Principal) middleware.Responder {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			str := `{
					"id": "5ca0f5e4-ac05-4480-9ee0-e896a22b827a",
					"location": "pune",
					"name": "1st device",
					"status": "active"
				}`
			rw.Write([]byte(str))
		})
	})

	api.GetDeviceVersionHandler = operations.GetDeviceVersionHandlerFunc(func(params operations.GetDeviceVersionParams, principal *models.Principal) middleware.Responder {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			rw.Write([]byte("{\"version\":\"v1\"}"))
		})

	})
}
