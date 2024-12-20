// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteGameHandlerFunc turns a function with the right signature into a delete game handler
type DeleteGameHandlerFunc func(DeleteGameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteGameHandlerFunc) Handle(params DeleteGameParams) middleware.Responder {
	return fn(params)
}

// DeleteGameHandler interface for that can handle valid delete game params
type DeleteGameHandler interface {
	Handle(DeleteGameParams) middleware.Responder
}

// NewDeleteGame creates a new http.Handler for the delete game operation
func NewDeleteGame(ctx *middleware.Context, handler DeleteGameHandler) *DeleteGame {
	return &DeleteGame{Context: ctx, Handler: handler}
}

/*
	DeleteGame swagger:route DELETE /games/{gameId} deleteGame

Delete a game by ID
*/
type DeleteGame struct {
	Context *middleware.Context
	Handler DeleteGameHandler
}

func (o *DeleteGame) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteGameParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
