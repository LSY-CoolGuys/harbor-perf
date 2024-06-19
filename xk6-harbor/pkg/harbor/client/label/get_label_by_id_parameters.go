// Code generated by go-swagger; DO NOT EDIT.

package label

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

// NewGetLabelByIDParams creates a new GetLabelByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLabelByIDParams() *GetLabelByIDParams {
	return &GetLabelByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLabelByIDParamsWithTimeout creates a new GetLabelByIDParams object
// with the ability to set a timeout on a request.
func NewGetLabelByIDParamsWithTimeout(timeout time.Duration) *GetLabelByIDParams {
	return &GetLabelByIDParams{
		timeout: timeout,
	}
}

// NewGetLabelByIDParamsWithContext creates a new GetLabelByIDParams object
// with the ability to set a context for a request.
func NewGetLabelByIDParamsWithContext(ctx context.Context) *GetLabelByIDParams {
	return &GetLabelByIDParams{
		Context: ctx,
	}
}

// NewGetLabelByIDParamsWithHTTPClient creates a new GetLabelByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLabelByIDParamsWithHTTPClient(client *http.Client) *GetLabelByIDParams {
	return &GetLabelByIDParams{
		HTTPClient: client,
	}
}

/*
GetLabelByIDParams contains all the parameters to send to the API endpoint

	for the get label by ID operation.

	Typically these are written to a http.Request.
*/
type GetLabelByIDParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* LabelID.

	   Label ID

	   Format: int64
	*/
	LabelID int64 `js:"labelID"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the get label by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLabelByIDParams) WithDefaults() *GetLabelByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get label by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLabelByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get label by ID params
func (o *GetLabelByIDParams) WithTimeout(timeout time.Duration) *GetLabelByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get label by ID params
func (o *GetLabelByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get label by ID params
func (o *GetLabelByIDParams) WithContext(ctx context.Context) *GetLabelByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get label by ID params
func (o *GetLabelByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get label by ID params
func (o *GetLabelByIDParams) WithHTTPClient(client *http.Client) *GetLabelByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get label by ID params
func (o *GetLabelByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get label by ID params
func (o *GetLabelByIDParams) WithXRequestID(xRequestID *string) *GetLabelByIDParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get label by ID params
func (o *GetLabelByIDParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithLabelID adds the labelID to the get label by ID params
func (o *GetLabelByIDParams) WithLabelID(labelID int64) *GetLabelByIDParams {
	o.SetLabelID(labelID)
	return o
}

// SetLabelID adds the labelId to the get label by ID params
func (o *GetLabelByIDParams) SetLabelID(labelID int64) {
	o.LabelID = labelID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLabelByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param label_id
	if err := r.SetPathParam("label_id", swag.FormatInt64(o.LabelID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
