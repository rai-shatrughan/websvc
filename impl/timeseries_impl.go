package impl

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	mw "websvc/middleware"
	"websvc/timeseries/models"
	"websvc/timeseries/restapi/operations"
)

var (
	topic = "ts"
	kf    = mw.KafkaWriter{}
	kv    = mw.KV{}
)

//Timeseries implements timeseries API middlewares
func Timeseries(api *operations.TimeseriesAPI) {
	api.Logger = log.Printf

	kf.Topic = &topic
	kv.New()
	kf.New()

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

	api.GetTimeseriesDeviceIDHandler = operations.GetTimeseriesDeviceIDHandlerFunc(func(params operations.GetTimeseriesDeviceIDParams, principal *models.Principal) middleware.Responder {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			key := "/" + string(params.DeviceID)
			response := kv.GetFromKeyWithLimit(key, 1000)
			rw.Write([]byte(response))
		})
	})

	api.PutTimeseriesDeviceIDHandler = operations.PutTimeseriesDeviceIDHandlerFunc(func(params operations.PutTimeseriesDeviceIDParams, principal *models.Principal) middleware.Responder {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			func(params operations.PutTimeseriesDeviceIDParams) {
				var response string
				ts, err := json.Marshal(params.Timeseries)
				if err != nil {
					log.Fatalf("error parsing TS : %s", err)
					response = `{ "TimeseriesUpload": "failed"}`
				}
				log.Println("TS -", string(ts))
				kf.Write([]byte(params.DeviceID), ts)
				response = `{ "TimeseriesUpload": "ok"}`
				rw.Write([]byte(response))
			}(params)
		})
	})

	api.GetTimeseriesVersionHandler = operations.GetTimeseriesVersionHandlerFunc(func(params operations.GetTimeseriesVersionParams, principal *models.Principal) middleware.Responder {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			str := `{
				"sessionId": "5ca0f5e4-ac05-4480-9ee0-e896a22b827a",
				"status": "deactive"
			}`
			rw.Write([]byte(str))
		})
	})

}
