// Code generated by go-swagger; DO NOT EDIT.

package robotv1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteRobotV1Params creates a new DeleteRobotV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRobotV1Params() *DeleteRobotV1Params {
	return &DeleteRobotV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRobotV1ParamsWithTimeout creates a new DeleteRobotV1Params object
// with the ability to set a timeout on a request.
func NewDeleteRobotV1ParamsWithTimeout(timeout time.Duration) *DeleteRobotV1Params {
	return &DeleteRobotV1Params{
		timeout: timeout,
	}
}

// NewDeleteRobotV1ParamsWithContext creates a new DeleteRobotV1Params object
// with the ability to set a context for a request.
func NewDeleteRobotV1ParamsWithContext(ctx context.Context) *DeleteRobotV1Params {
	return &DeleteRobotV1Params{
		Context: ctx,
	}
}

// NewDeleteRobotV1ParamsWithHTTPClient creates a new DeleteRobotV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRobotV1ParamsWithHTTPClient(client *http.Client) *DeleteRobotV1Params {
	return &DeleteRobotV1Params{
		HTTPClient: client,
	}
}

/*
DeleteRobotV1Params contains all the parameters to send to the API endpoint

	for the delete robot v1 operation.

	Typically these are written to a http.Request.
*/
type DeleteRobotV1Params struct {

	/* XIsResourceName.

	   The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.
	*/
	XIsResourceName *bool `js:"xIsResourceName"`

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* ProjectNameOrID.

	   The name or id of the project
	*/
	ProjectNameOrID string `js:"projectNameOrID"`

	/* RobotID.

	   Robot ID
	*/
	RobotID int64 `js:"robotID"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the delete robot v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRobotV1Params) WithDefaults() *DeleteRobotV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete robot v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRobotV1Params) SetDefaults() {
	var (
		xIsResourceNameDefault = bool(false)
	)

	val := DeleteRobotV1Params{
		XIsResourceName: &xIsResourceNameDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete robot v1 params
func (o *DeleteRobotV1Params) WithTimeout(timeout time.Duration) *DeleteRobotV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete robot v1 params
func (o *DeleteRobotV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete robot v1 params
func (o *DeleteRobotV1Params) WithContext(ctx context.Context) *DeleteRobotV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete robot v1 params
func (o *DeleteRobotV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete robot v1 params
func (o *DeleteRobotV1Params) WithHTTPClient(client *http.Client) *DeleteRobotV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete robot v1 params
func (o *DeleteRobotV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the delete robot v1 params
func (o *DeleteRobotV1Params) WithXIsResourceName(xIsResourceName *bool) *DeleteRobotV1Params {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the delete robot v1 params
func (o *DeleteRobotV1Params) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the delete robot v1 params
func (o *DeleteRobotV1Params) WithXRequestID(xRequestID *string) *DeleteRobotV1Params {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete robot v1 params
func (o *DeleteRobotV1Params) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectNameOrID adds the projectNameOrID to the delete robot v1 params
func (o *DeleteRobotV1Params) WithProjectNameOrID(projectNameOrID string) *DeleteRobotV1Params {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the delete robot v1 params
func (o *DeleteRobotV1Params) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WithRobotID adds the robotID to the delete robot v1 params
func (o *DeleteRobotV1Params) WithRobotID(robotID int64) *DeleteRobotV1Params {
	o.SetRobotID(robotID)
	return o
}

// SetRobotID adds the robotId to the delete robot v1 params
func (o *DeleteRobotV1Params) SetRobotID(robotID int64) {
	o.RobotID = robotID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRobotV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XIsResourceName != nil {

		// header param X-Is-Resource-Name
		if err := r.SetHeaderParam("X-Is-Resource-Name", swag.FormatBool(*o.XIsResourceName)); err != nil {
			return err
		}
	}

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param project_name_or_id
	if err := r.SetPathParam("project_name_or_id", o.ProjectNameOrID); err != nil {
		return err
	}

	// path param robot_id
	if err := r.SetPathParam("robot_id", swag.FormatInt64(o.RobotID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
