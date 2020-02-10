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

// NewGetAssignmentLogParams creates a new GetAssignmentLogParams object
// with the default values initialized.
func NewGetAssignmentLogParams() *GetAssignmentLogParams {
	var ()
	return &GetAssignmentLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAssignmentLogParamsWithTimeout creates a new GetAssignmentLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAssignmentLogParamsWithTimeout(timeout time.Duration) *GetAssignmentLogParams {
	var ()
	return &GetAssignmentLogParams{

		timeout: timeout,
	}
}

// NewGetAssignmentLogParamsWithContext creates a new GetAssignmentLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAssignmentLogParamsWithContext(ctx context.Context) *GetAssignmentLogParams {
	var ()
	return &GetAssignmentLogParams{

		Context: ctx,
	}
}

// NewGetAssignmentLogParamsWithHTTPClient creates a new GetAssignmentLogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAssignmentLogParamsWithHTTPClient(client *http.Client) *GetAssignmentLogParams {
	var ()
	return &GetAssignmentLogParams{
		HTTPClient: client,
	}
}

/*GetAssignmentLogParams contains all the parameters to send to the API endpoint
for the get assignment log operation typically these are written to a http.Request
*/
type GetAssignmentLogParams struct {

	/*ID
	  The assignment's id

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get assignment log params
func (o *GetAssignmentLogParams) WithTimeout(timeout time.Duration) *GetAssignmentLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get assignment log params
func (o *GetAssignmentLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get assignment log params
func (o *GetAssignmentLogParams) WithContext(ctx context.Context) *GetAssignmentLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get assignment log params
func (o *GetAssignmentLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get assignment log params
func (o *GetAssignmentLogParams) WithHTTPClient(client *http.Client) *GetAssignmentLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get assignment log params
func (o *GetAssignmentLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get assignment log params
func (o *GetAssignmentLogParams) WithID(id string) *GetAssignmentLogParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get assignment log params
func (o *GetAssignmentLogParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAssignmentLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
