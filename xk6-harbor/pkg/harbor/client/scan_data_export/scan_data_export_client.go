// Code generated by go-swagger; DO NOT EDIT.

package scan_data_export

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

// API is the interface of the scan data export client
type API interface {
	/*
	   DownloadScanData downloads the scan data export file

	   Download the scan data report. Default format is CSV*/
	DownloadScanData(ctx context.Context, params *DownloadScanDataParams, writer io.Writer) (*DownloadScanDataOK, error)
	/*
	   ExportScanData exports scan data for selected projects

	   Export scan data for selected projects*/
	ExportScanData(ctx context.Context, params *ExportScanDataParams) (*ExportScanDataOK, error)
	/*
	   GetScanDataExportExecution gets the specific scan data export execution

	   Get the scan data export execution specified by ID*/
	GetScanDataExportExecution(ctx context.Context, params *GetScanDataExportExecutionParams) (*GetScanDataExportExecutionOK, error)
	/*
	   GetScanDataExportExecutionList gets a list of specific scan data export execution jobs for a specified user

	   Get a list of specific scan data export execution jobs for a specified user*/
	GetScanDataExportExecutionList(ctx context.Context, params *GetScanDataExportExecutionListParams) (*GetScanDataExportExecutionListOK, error)
}

// New creates a new scan data export API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for scan data export API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DownloadScanData downloads the scan data export file

Download the scan data report. Default format is CSV
*/
func (a *Client) DownloadScanData(ctx context.Context, params *DownloadScanDataParams, writer io.Writer) (*DownloadScanDataOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "downloadScanData",
		Method:             "GET",
		PathPattern:        "/export/cve/download/{execution_id}",
		ProducesMediaTypes: []string{"text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DownloadScanDataReader{formats: a.formats, writer: writer},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *DownloadScanDataOK:
		return value, nil
	case *DownloadScanDataUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *DownloadScanDataForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *DownloadScanDataNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *DownloadScanDataInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for downloadScanData: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExportScanData exports scan data for selected projects

Export scan data for selected projects
*/
func (a *Client) ExportScanData(ctx context.Context, params *ExportScanDataParams) (*ExportScanDataOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "exportScanData",
		Method:             "POST",
		PathPattern:        "/export/cve",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ExportScanDataReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *ExportScanDataOK:
		return value, nil
	case *ExportScanDataBadRequest:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ExportScanDataUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ExportScanDataForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ExportScanDataNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ExportScanDataMethodNotAllowed:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ExportScanDataConflict:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ExportScanDataInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for exportScanData: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetScanDataExportExecution gets the specific scan data export execution

Get the scan data export execution specified by ID
*/
func (a *Client) GetScanDataExportExecution(ctx context.Context, params *GetScanDataExportExecutionParams) (*GetScanDataExportExecutionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScanDataExportExecution",
		Method:             "GET",
		PathPattern:        "/export/cve/execution/{execution_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetScanDataExportExecutionReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetScanDataExportExecutionOK:
		return value, nil
	case *GetScanDataExportExecutionUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetScanDataExportExecutionForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetScanDataExportExecutionNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetScanDataExportExecutionInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getScanDataExportExecution: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetScanDataExportExecutionList gets a list of specific scan data export execution jobs for a specified user

Get a list of specific scan data export execution jobs for a specified user
*/
func (a *Client) GetScanDataExportExecutionList(ctx context.Context, params *GetScanDataExportExecutionListParams) (*GetScanDataExportExecutionListOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScanDataExportExecutionList",
		Method:             "GET",
		PathPattern:        "/export/cve/executions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetScanDataExportExecutionListReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetScanDataExportExecutionListOK:
		return value, nil
	case *GetScanDataExportExecutionListUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetScanDataExportExecutionListForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetScanDataExportExecutionListNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetScanDataExportExecutionListInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getScanDataExportExecutionList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}
