// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ActionRequest The request to stop, pause or resume
//
// swagger:model ActionRequest
type ActionRequest struct {

	// The action of the request, should be stop, pause or resume
	// Enum: [stop pause resume]
	Action string `json:"action,omitempty" js:"action"`
}

// Validate validates this action request
func (m *ActionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var actionRequestTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["stop","pause","resume"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		actionRequestTypeActionPropEnum = append(actionRequestTypeActionPropEnum, v)
	}
}

const (

	// ActionRequestActionStop captures enum value "stop"
	ActionRequestActionStop string = "stop"

	// ActionRequestActionPause captures enum value "pause"
	ActionRequestActionPause string = "pause"

	// ActionRequestActionResume captures enum value "resume"
	ActionRequestActionResume string = "resume"
)

// prop value enum
func (m *ActionRequest) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, actionRequestTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ActionRequest) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this action request based on context it is used
func (m *ActionRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ActionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionRequest) UnmarshalBinary(b []byte) error {
	var res ActionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
