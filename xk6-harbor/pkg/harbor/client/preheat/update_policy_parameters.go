// Code generated by go-swagger; DO NOT EDIT.

package preheat

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

// NewUpdatePolicyParams creates a new UpdatePolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePolicyParams() *UpdatePolicyParams {
	return &UpdatePolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePolicyParamsWithTimeout creates a new UpdatePolicyParams object
// with the ability to set a timeout on a request.
func NewUpdatePolicyParamsWithTimeout(timeout time.Duration) *UpdatePolicyParams {
	return &UpdatePolicyParams{
		timeout: timeout,
	}
}

// NewUpdatePolicyParamsWithContext creates a new UpdatePolicyParams object
// with the ability to set a context for a request.
func NewUpdatePolicyParamsWithContext(ctx context.Context) *UpdatePolicyParams {
	return &UpdatePolicyParams{
		Context: ctx,
	}
}

// NewUpdatePolicyParamsWithHTTPClient creates a new UpdatePolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePolicyParamsWithHTTPClient(client *http.Client) *UpdatePolicyParams {
	return &UpdatePolicyParams{
		HTTPClient: client,
	}
}

/*
UpdatePolicyParams contains all the parameters to send to the API endpoint

	for the update policy operation.

	Typically these are written to a http.Request.
*/
type UpdatePolicyParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* Policy.

	   The policy schema info
	*/
	Policy *models.PreheatPolicy `js:"policy"`

	/* PreheatPolicyName.

	   Preheat Policy Name
	*/
	PreheatPolicyName string `js:"preheatPolicyName"`

	/* ProjectName.

	   The name of the project
	*/
	ProjectName string `js:"projectName"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the update policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePolicyParams) WithDefaults() *UpdatePolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update policy params
func (o *UpdatePolicyParams) WithTimeout(timeout time.Duration) *UpdatePolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update policy params
func (o *UpdatePolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update policy params
func (o *UpdatePolicyParams) WithContext(ctx context.Context) *UpdatePolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update policy params
func (o *UpdatePolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update policy params
func (o *UpdatePolicyParams) WithHTTPClient(client *http.Client) *UpdatePolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update policy params
func (o *UpdatePolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update policy params
func (o *UpdatePolicyParams) WithXRequestID(xRequestID *string) *UpdatePolicyParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update policy params
func (o *UpdatePolicyParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPolicy adds the policy to the update policy params
func (o *UpdatePolicyParams) WithPolicy(policy *models.PreheatPolicy) *UpdatePolicyParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the update policy params
func (o *UpdatePolicyParams) SetPolicy(policy *models.PreheatPolicy) {
	o.Policy = policy
}

// WithPreheatPolicyName adds the preheatPolicyName to the update policy params
func (o *UpdatePolicyParams) WithPreheatPolicyName(preheatPolicyName string) *UpdatePolicyParams {
	o.SetPreheatPolicyName(preheatPolicyName)
	return o
}

// SetPreheatPolicyName adds the preheatPolicyName to the update policy params
func (o *UpdatePolicyParams) SetPreheatPolicyName(preheatPolicyName string) {
	o.PreheatPolicyName = preheatPolicyName
}

// WithProjectName adds the projectName to the update policy params
func (o *UpdatePolicyParams) WithProjectName(projectName string) *UpdatePolicyParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the update policy params
func (o *UpdatePolicyParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Policy != nil {
		if err := r.SetBodyParam(o.Policy); err != nil {
			return err
		}
	}

	// path param preheat_policy_name
	if err := r.SetPathParam("preheat_policy_name", o.PreheatPolicyName); err != nil {
		return err
	}

	// path param project_name
	if err := r.SetPathParam("project_name", o.ProjectName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
