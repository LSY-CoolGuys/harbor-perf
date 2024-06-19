// Code generated by go-swagger; DO NOT EDIT.

package scan_all

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

	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
)

// NewUpdateScanAllScheduleParams creates a new UpdateScanAllScheduleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateScanAllScheduleParams() *UpdateScanAllScheduleParams {
	return &UpdateScanAllScheduleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateScanAllScheduleParamsWithTimeout creates a new UpdateScanAllScheduleParams object
// with the ability to set a timeout on a request.
func NewUpdateScanAllScheduleParamsWithTimeout(timeout time.Duration) *UpdateScanAllScheduleParams {
	return &UpdateScanAllScheduleParams{
		timeout: timeout,
	}
}

// NewUpdateScanAllScheduleParamsWithContext creates a new UpdateScanAllScheduleParams object
// with the ability to set a context for a request.
func NewUpdateScanAllScheduleParamsWithContext(ctx context.Context) *UpdateScanAllScheduleParams {
	return &UpdateScanAllScheduleParams{
		Context: ctx,
	}
}

// NewUpdateScanAllScheduleParamsWithHTTPClient creates a new UpdateScanAllScheduleParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateScanAllScheduleParamsWithHTTPClient(client *http.Client) *UpdateScanAllScheduleParams {
	return &UpdateScanAllScheduleParams{
		HTTPClient: client,
	}
}

/*
UpdateScanAllScheduleParams contains all the parameters to send to the API endpoint

	for the update scan all schedule operation.

	Typically these are written to a http.Request.
*/
type UpdateScanAllScheduleParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* Schedule.

	   Updates the schedule of scan all job, which scans all of images in Harbor.
	*/
	Schedule *models.Schedule `js:"schedule"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the update scan all schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateScanAllScheduleParams) WithDefaults() *UpdateScanAllScheduleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update scan all schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateScanAllScheduleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) WithTimeout(timeout time.Duration) *UpdateScanAllScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) WithContext(ctx context.Context) *UpdateScanAllScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) WithHTTPClient(client *http.Client) *UpdateScanAllScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) WithXRequestID(xRequestID *string) *UpdateScanAllScheduleParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithSchedule adds the schedule to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) WithSchedule(schedule *models.Schedule) *UpdateScanAllScheduleParams {
	o.SetSchedule(schedule)
	return o
}

// SetSchedule adds the schedule to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) SetSchedule(schedule *models.Schedule) {
	o.Schedule = schedule
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateScanAllScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Schedule != nil {
		if err := r.SetBodyParam(o.Schedule); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
