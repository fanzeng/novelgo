// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"novelgo/internal/pkg/restapi/operations"
)

//go:generate swagger generate server --target ../../../api --name Novelgo --spec ../../../api/swagger.yml --model-package ../internal/pkg/models --server-package ../internal/pkg/restapi --principal interface{}

func configureFlags(api *operations.NovelgoAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.NovelgoAPI) http.Handler {
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

	if api.CreateGameHandler == nil {
		api.CreateGameHandler = operations.CreateGameHandlerFunc(func(params operations.CreateGameParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateGame has not yet been implemented")
		})
	}
	if api.DeleteGameHandler == nil {
		api.DeleteGameHandler = operations.DeleteGameHandlerFunc(func(params operations.DeleteGameParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteGame has not yet been implemented")
		})
	}
	if api.GetGameByIDHandler == nil {
		api.GetGameByIDHandler = operations.GetGameByIDHandlerFunc(func(params operations.GetGameByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetGameByID has not yet been implemented")
		})
	}
	if api.ListGamesHandler == nil {
		api.ListGamesHandler = operations.ListGamesHandlerFunc(func(params operations.ListGamesParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ListGames has not yet been implemented")
		})
	}
	if api.UpdateGameHandler == nil {
		api.UpdateGameHandler = operations.UpdateGameHandlerFunc(func(params operations.UpdateGameParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateGame has not yet been implemented")
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
