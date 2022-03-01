package impl

import (
	"log"
	"net/http"

	"websvc/auth/models"
	"websvc/auth/restapi/operations"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

//Auth implements auth API middlewares
func Auth(api *operations.AuthAPI) {
	api.Logger = log.Printf

	api.BasicAuthAuth = func(user string, pass string) (*models.Principal, error) {
		if user == "sr" && pass == "sr123" {
			prin := models.Principal(user)
			return &prin, nil
		}
		api.Logger("Access attempt with incorrect api key auth: %s", user)
		return nil, errors.New(401, "incorrect api key auth")
	}

	api.PostLoginHandler = operations.PostLoginHandlerFunc(func(params operations.PostLoginParams, principal *models.Principal) middleware.Responder {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			str := `{
					"sessionId": "5ca0f5e4-ac05-4480-9ee0-e896a22b827a",
					"status": "active",
					"expiry": "30m"
				}`
			rw.Write([]byte(str))
		})
	})

	api.PostLogoutHandler = operations.PostLogoutHandlerFunc(func(params operations.PostLogoutParams) middleware.Responder {
		return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
			str := `{
				"sessionId": "5ca0f5e4-ac05-4480-9ee0-e896a22b827a",
				"status": "deactive"
			}`
			rw.Write([]byte(str))
		})
	})

}
