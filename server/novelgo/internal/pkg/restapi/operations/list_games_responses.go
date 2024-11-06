// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"novelgo/internal/pkg/models"
)

// ListGamesOKCode is the HTTP code returned for type ListGamesOK
const ListGamesOKCode int = 200

/*
ListGamesOK An array of games

swagger:response listGamesOK
*/
type ListGamesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Game `json:"body,omitempty"`
}

// NewListGamesOK creates ListGamesOK with default headers values
func NewListGamesOK() *ListGamesOK {

	return &ListGamesOK{}
}

// WithPayload adds the payload to the list games o k response
func (o *ListGamesOK) WithPayload(payload []*models.Game) *ListGamesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list games o k response
func (o *ListGamesOK) SetPayload(payload []*models.Game) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListGamesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Game, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
