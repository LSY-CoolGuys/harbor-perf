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

// WorkerPool the worker pool of job service
//
// swagger:model WorkerPool
type WorkerPool struct {

	// The concurrency of the work pool
	Concurrency int64 `json:"concurrency,omitempty" js:"concurrency"`

	// The heartbeat time of the work pool
	// Format: date-time
	HeartbeatAt strfmt.DateTime `json:"heartbeat_at,omitempty" js:"heartbeatAt"`

	// The host of the work pool
	Host string `json:"host,omitempty" js:"host"`

	// the process id of jobservice
	Pid int64 `json:"pid,omitempty" js:"pid"`

	// The start time of the work pool
	// Format: date-time
	StartAt strfmt.DateTime `json:"start_at,omitempty" js:"startAt"`

	// the id of the worker pool
	WorkerPoolID string `json:"worker_pool_id,omitempty" js:"workerPoolID"`
}

// Validate validates this worker pool
func (m *WorkerPool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeartbeatAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkerPool) validateHeartbeatAt(formats strfmt.Registry) error {
	if swag.IsZero(m.HeartbeatAt) { // not required
		return nil
	}

	if err := validate.FormatOf("heartbeat_at", "body", "date-time", m.HeartbeatAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WorkerPool) validateStartAt(formats strfmt.Registry) error {
	if swag.IsZero(m.StartAt) { // not required
		return nil
	}

	if err := validate.FormatOf("start_at", "body", "date-time", m.StartAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this worker pool based on context it is used
func (m *WorkerPool) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkerPool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkerPool) UnmarshalBinary(b []byte) error {
	var res WorkerPool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
