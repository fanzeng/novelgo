// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Game game
//
// swagger:model Game
type Game struct {

	// gameplay
	// Required: true
	Gameplay *GameGameplay `json:"Gameplay"`

	// Id
	// Required: true
	ID *string `json:"Id"`

	// name
	// Required: true
	Name *string `json:"Name"`

	// settings
	// Required: true
	Settings *GameSettings `json:"Settings"`
}

// Validate validates this game
func (m *Game) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGameplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Game) validateGameplay(formats strfmt.Registry) error {

	if err := validate.Required("Gameplay", "body", m.Gameplay); err != nil {
		return err
	}

	if m.Gameplay != nil {
		if err := m.Gameplay.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Gameplay")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Gameplay")
			}
			return err
		}
	}

	return nil
}

func (m *Game) validateID(formats strfmt.Registry) error {

	if err := validate.Required("Id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Game) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Game) validateSettings(formats strfmt.Registry) error {

	if err := validate.Required("Settings", "body", m.Settings); err != nil {
		return err
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Settings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this game based on the context it is used
func (m *Game) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGameplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Game) contextValidateGameplay(ctx context.Context, formats strfmt.Registry) error {

	if m.Gameplay != nil {

		if err := m.Gameplay.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Gameplay")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Gameplay")
			}
			return err
		}
	}

	return nil
}

func (m *Game) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.Settings != nil {

		if err := m.Settings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Game) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Game) UnmarshalBinary(b []byte) error {
	var res Game
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GameGameplay game gameplay
//
// swagger:model GameGameplay
type GameGameplay struct {

	// player moves
	PlayerMoves []*GameGameplayPlayerMovesItems0 `json:"PlayerMoves"`
}

// Validate validates this game gameplay
func (m *GameGameplay) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlayerMoves(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GameGameplay) validatePlayerMoves(formats strfmt.Registry) error {
	if swag.IsZero(m.PlayerMoves) { // not required
		return nil
	}

	for i := 0; i < len(m.PlayerMoves); i++ {
		if swag.IsZero(m.PlayerMoves[i]) { // not required
			continue
		}

		if m.PlayerMoves[i] != nil {
			if err := m.PlayerMoves[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Gameplay" + "." + "PlayerMoves" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Gameplay" + "." + "PlayerMoves" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this game gameplay based on the context it is used
func (m *GameGameplay) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlayerMoves(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GameGameplay) contextValidatePlayerMoves(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PlayerMoves); i++ {

		if m.PlayerMoves[i] != nil {

			if swag.IsZero(m.PlayerMoves[i]) { // not required
				return nil
			}

			if err := m.PlayerMoves[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Gameplay" + "." + "PlayerMoves" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Gameplay" + "." + "PlayerMoves" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GameGameplay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GameGameplay) UnmarshalBinary(b []byte) error {
	var res GameGameplay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GameGameplayPlayerMovesItems0 game gameplay player moves items0
//
// swagger:model GameGameplayPlayerMovesItems0
type GameGameplayPlayerMovesItems0 struct {

	// col
	Col int64 `json:"Col,omitempty"`

	// row
	Row int64 `json:"Row,omitempty"`
}

// Validate validates this game gameplay player moves items0
func (m *GameGameplayPlayerMovesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this game gameplay player moves items0 based on context it is used
func (m *GameGameplayPlayerMovesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GameGameplayPlayerMovesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GameGameplayPlayerMovesItems0) UnmarshalBinary(b []byte) error {
	var res GameGameplayPlayerMovesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GameSettings game settings
//
// swagger:model GameSettings
type GameSettings struct {

	// board height
	BoardHeight int64 `json:"BoardHeight,omitempty"`

	// board width
	BoardWidth int64 `json:"BoardWidth,omitempty"`
}

// Validate validates this game settings
func (m *GameSettings) Validate(formats strfmt.Registry) error {
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
