// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewAgentParams creates a new AgentParams object
// with the default values initialized.
func NewAgentParams() *AgentParams {

	return &AgentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAgentParamsWithTimeout creates a new AgentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAgentParamsWithTimeout(timeout time.Duration) *AgentParams {

	return &AgentParams{

		timeout: timeout,
	}
}

// NewAgentParamsWithContext creates a new AgentParams object
// with the default values initialized, and the ability to set a context for a request
func NewAgentParamsWithContext(ctx context.Context) *AgentParams {

	return &AgentParams{

		Context: ctx,
	}
}

// NewAgentParamsWithHTTPClient creates a new AgentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAgentParamsWithHTTPClient(client *http.Client) *AgentParams {

	return &AgentParams{
		HTTPClient: client,
	}
}

/*AgentParams contains all the parameters to send to the API endpoint
for the agent operation typically these are written to a http.Request
*/
type AgentParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the agent params
func (o *AgentParams) WithTimeout(timeout time.Duration) *AgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the agent params
func (o *AgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the agent params
func (o *AgentParams) WithContext(ctx context.Context) *AgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the agent params
func (o *AgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the agent params
func (o *AgentParams) WithHTTPClient(client *http.Client) *AgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the agent params
func (o *AgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}