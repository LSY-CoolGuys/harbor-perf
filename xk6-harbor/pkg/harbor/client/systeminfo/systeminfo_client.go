// Code generated by go-swagger; DO NOT EDIT.

package systeminfo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery --name API --keeptree --with-expecter --case underscore

// API is the interface of the systeminfo client
type API interface {
	/*
	   GetCert gets default root certificate

	   This endpoint is for downloading a default root certificate.
	*/
	GetCert(ctx context.Context, params *GetCertParams, writer io.Writer) (*GetCertOK, error)
	/*
	   GetSystemInfo gets general system info

	   This API is for retrieving general system info, this can be called by anonymous request.  Some attributes will be omitted in the response when this API is called by anonymous request.
	*/
	GetSystemInfo(ctx context.Context, params *GetSystemInfoParams) (*GetSystemInfoOK, error)
	/*
	   GetVolumes gets system volume info total free size

	   This endpoint is for retrieving system volume info that only provides for admin user.  Note that the response only reflects the storage status of local disk.
	*/
	GetVolumes(ctx context.Context, params *GetVolumesParams) (*GetVolumesOK, error)
}

// New creates a new systeminfo API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for systeminfo API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
GetCert gets default root certificate

This endpoint is for downloading a default root certificate.
*/
func (a *Client) GetCert(ctx context.Context, params *GetCertParams, writer io.Writer) (*GetCertOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCert",
		Method:             "GET",
		PathPattern:        "/systeminfo/getcert",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCertReader{formats: a.formats, writer: writer},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetCertOK:
		return value, nil
	case *GetCertNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetCertInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCert: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSystemInfo gets general system info

This API is for retrieving general system info, this can be called by anonymous request.  Some attributes will be omitted in the response when this API is called by anonymous request.
*/
func (a *Client) GetSystemInfo(ctx context.Context, params *GetSystemInfoParams) (*GetSystemInfoOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSystemInfo",
		Method:             "GET",
		PathPattern:        "/systeminfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSystemInfoReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetSystemInfoOK:
		return value, nil
	case *GetSystemInfoInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSystemInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetVolumes gets system volume info total free size

This endpoint is for retrieving system volume info that only provides for admin user.  Note that the response only reflects the storage status of local disk.
*/
func (a *Client) GetVolumes(ctx context.Context, params *GetVolumesParams) (*GetVolumesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVolumes",
		Method:             "GET",
		PathPattern:        "/systeminfo/volumes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetVolumesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetVolumesOK:
		return value, nil
	case *GetVolumesUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetVolumesForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetVolumesNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetVolumesInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVolumes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}
