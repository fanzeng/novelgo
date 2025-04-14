// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GameSettings game settings
//
// swagger:model GameSettings
type GameSettings struct {

	// board height
	// Required: true
	BoardHeight *int64 `json:"BoardHeight"`

	// board width
	// Required: true
	BoardWidth *int64 `json:"BoardWidth"`

	// cyclic logic
	// Required: true
	CyclicLogic *bool `json:"CyclicLogic"`
}

// Validate validates this game settings
func (m *GameSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoardHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBoardWidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCyclicLogic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GameSettings) validateBoardHeight(formats strfmt.Registry) error {

	if err := validate.Required("BoardHeight", "body", m.BoardHeight); err != nil {
		return err
	}

	return nil
}

func (m *GameSettings) validateBoardWidth(formats strfmt.Registry) error {

	if err := validate.Required("BoardWidth", "body", m.BoardWidth); err != nil {
		return err
	}

	return nil
}

func (m *GameSettings) validateCyclicLogic(formats strfmt.Registry) error {

	if err := validate.Required("CyclicLogic", "body", m.CyclicLogic); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this game settings based on context it is used
func (m *GameSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GameSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GameSettings) UnmarshalBinary(b []byte) error {
	var res GameSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
