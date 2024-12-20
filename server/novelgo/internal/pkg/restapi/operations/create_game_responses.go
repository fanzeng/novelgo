// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"novelgo/internal/pkg/models"
)

// CreateGameCreatedCode is the HTTP code returned for type CreateGameCreated
const CreateGameCreatedCode int = 201

/*
CreateGameCreated Game created

swagger:response createGameCreated
*/
type CreateGameCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Game `json:"body,omitempty"`
}

// NewCreateGameCreated creates CreateGameCreated with default headers values
func NewCreateGameCreated() *CreateGameCreated {

	return &CreateGameCreated{}
}

// WithPayload adds the payload to the create game created response
func (o *CreateGameCreated) WithPayload(payload *models.Game) *CreateGameCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create game created response
func (o *CreateGameCreated) SetPayload(payload *models.Game) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateGameCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
