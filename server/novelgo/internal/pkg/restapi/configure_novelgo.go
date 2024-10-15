// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"novelgo/internal/pkg/handlers"
	"novelgo/internal/pkg/models"
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

	api.CreateGameHandler = operations.CreateGameHandlerFunc(func(params operations.CreateGameParams) middleware.Responder {
		newGame := &models.Game{
			ID:       params.Body.ID,
			Name:     params.Body.Name,
			Settings: params.Body.Settings,
			Gameplay: params.Body.Gameplay,
		}
		newGame, err := handlers.CreateGame(newGame)
		if err != nil {
			return middleware.ResponderFunc(func(w http.ResponseWriter, _ runtime.Producer) {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"message": "failed to creat new game"}`))
			})
		}
		return operations.NewCreateGameCreated().WithPayload(newGame)
	})
	api.DeleteGameHandler = operations.DeleteGameHandlerFunc(func(params operations.DeleteGameParams) middleware.Responder {
		err := handlers.DeleteGame(params.GameID)
		if err != nil {
			return operations.NewDeleteGameNotFound()
		}
		return operations.NewDeleteGameNoContent()
	})
	api.GetGameByIDHandler = operations.GetGameByIDHandlerFunc(func(params operations.GetGameByIDParams) middleware.Responder {
		game, err := handlers.GetGameByID(params.GameID)
		if err != nil {
			return operations.NewGetGameByIDNotFound()
		}
		return operations.NewGetGameByIDOK().WithPayload(game)
	})
	api.ListGamesHandler = operations.ListGamesHandlerFunc(func(params operations.ListGamesParams) middleware.Responder {
		games := handlers.ListGames()
		return operations.NewListGamesOK().WithPayload(games)
	})
	api.UpdateGameHandler = operations.UpdateGameHandlerFunc(func(params operations.UpdateGameParams) middleware.Responder {
		updatedGame := &models.Game{
			ID:       &params.GameID,
			Name:     params.Body.Name,
			Settings: params.Body.Settings,
			Gameplay: params.Body.Gameplay,
		}
		err := handlers.UpdateGame(params.GameID, updatedGame)
		if err != nil {
			return operations.NewUpdateGameNotFound()
		}
		return operations.NewUpdateGameOK().WithPayload(updatedGame)
	})

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
