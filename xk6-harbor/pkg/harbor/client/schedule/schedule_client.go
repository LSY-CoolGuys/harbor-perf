// Code generated by go-swagger; DO NOT EDIT.

package schedule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery --name API --keeptree --with-expecter --case underscore

// API is the interface of the schedule client
type API interface {
	/*
	   GetSchedulePaused Get scheduler paused status*/
	GetSchedulePaused(ctx context.Context, params *GetSchedulePausedParams) (*GetSchedulePausedOK, error)
	/*
	   ListSchedules List schedules*/
	ListSchedules(ctx context.Context, params *ListSchedulesParams) (*ListSchedulesOK, error)
}

// New creates a new schedule API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for schedule API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
GetSchedulePaused Get scheduler paused status
*/
func (a *Client) GetSchedulePaused(ctx context.Context, params *GetSchedulePausedParams) (*GetSchedulePausedOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSchedulePaused",
		Method:             "GET",
		PathPattern:        "/schedules/{job_type}/paused",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSchedulePausedReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetSchedulePausedOK:
		return value, nil
	case *GetSchedulePausedUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetSchedulePausedForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetSchedulePausedNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetSchedulePausedInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSchedulePaused: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListSchedules List schedules
*/
func (a *Client) ListSchedules(ctx context.Context, params *ListSchedulesParams) (*ListSchedulesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSchedules",
		Method:             "GET",
		PathPattern:        "/schedules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListSchedulesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *ListSchedulesOK:
		return value, nil
	case *ListSchedulesUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ListSchedulesForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ListSchedulesNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ListSchedulesInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listSchedules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}
