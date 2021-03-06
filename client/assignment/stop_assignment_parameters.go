// Code generated by go-swagger; DO NOT EDIT.

package assignment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStopAssignmentParams creates a new StopAssignmentParams object
// with the default values initialized.
func NewStopAssignmentParams() *StopAssignmentParams {
	var ()
	return &StopAssignmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopAssignmentParamsWithTimeout creates a new StopAssignmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopAssignmentParamsWithTimeout(timeout time.Duration) *StopAssignmentParams {
	var ()
	return &StopAssignmentParams{

		timeout: timeout,
	}
}

// NewStopAssignmentParamsWithContext creates a new StopAssignmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopAssignmentParamsWithContext(ctx context.Context) *StopAssignmentParams {
	var ()
	return &StopAssignmentParams{

		Context: ctx,
	}
}

// NewStopAssignmentParamsWithHTTPClient creates a new StopAssignmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopAssignmentParamsWithHTTPClient(client *http.Client) *StopAssignmentParams {
	var ()
	return &StopAssignmentParams{
		HTTPClient: client,
	}
}

/*StopAssignmentParams contains all the parameters to send to the API endpoint
for the stop assignment operation typically these are written to a http.Request
*/
type StopAssignmentParams struct {

	/*ID
	  The assignment's id

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop assignment params
func (o *StopAssignmentParams) WithTimeout(timeout time.Duration) *StopAssignmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop assignment params
func (o *StopAssignmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop assignment params
func (o *StopAssignmentParams) WithContext(ctx context.Context) *StopAssignmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop assignment params
func (o *StopAssignmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop assignment params
func (o *StopAssignmentParams) WithHTTPClient(client *http.Client) *StopAssignmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop assignment params
func (o *StopAssignmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the stop assignment params
func (o *StopAssignmentParams) WithID(id string) *StopAssignmentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the stop assignment params
func (o *StopAssignmentParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StopAssignmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Id
	if err := r.SetPathParam("Id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
