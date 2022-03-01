// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"websvc/auth/models"
	"websvc/auth/restapi/operations"
	"websvc/impl"
)

//go:generate swagger generate server --target ../../auth --name Auth --spec ../../swaggers/auth.yaml --principal models.Principal

func configureFlags(api *operations.AuthAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.AuthAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the Authorization header is set with the Basic scheme
	if api.BasicAuthAuth == nil {
		api.BasicAuthAuth = func(user string, pass string) (*models.Principal, error) {
			return nil, errors.NotImplemented("basic auth  (basicAuth) has not yet been implemented")
		}
	}

	impl.Auth(api)

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.PostLoginHandler == nil {
		api.PostLoginHandler = operations.PostLoginHandlerFunc(func(params operations.PostLoginParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostLogin has not yet been implemented")
		})
	}
	if api.PostLogoutHandler == nil {
		api.PostLogoutHandler = operations.PostLogoutHandlerFunc(func(params operations.PostLogoutParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostLogout has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
