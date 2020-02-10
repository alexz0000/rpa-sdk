// Code generated by go-swagger; DO NOT EDIT.

package assignment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new assignment API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for assignment API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAssignment(params *GetAssignmentParams, authInfo runtime.ClientAuthInfoWriter) (*GetAssignmentOK, error)

	GetAssignmentLog(params *GetAssignmentLogParams, authInfo runtime.ClientAuthInfoWriter) (*GetAssignmentLogOK, error)

	GetAssignmentResult(params *GetAssignmentResultParams, authInfo runtime.ClientAuthInfoWriter) (*GetAssignmentResultOK, error)

	GetAssignments(params *GetAssignmentsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAssignmentsOK, error)

	PostAssignment(params *PostAssignmentParams, authInfo runtime.ClientAuthInfoWriter) (*PostAssignmentOK, error)

	ReportAssignmentLog(params *ReportAssignmentLogParams, authInfo runtime.ClientAuthInfoWriter) (*ReportAssignmentLogOK, error)

	ReportAssignmentResult(params *ReportAssignmentResultParams, authInfo runtime.ClientAuthInfoWriter) (*ReportAssignmentResultOK, error)

	StopAssignment(params *StopAssignmentParams, authInfo runtime.ClientAuthInfoWriter) (*StopAssignmentOK, error)

	UpdateAssignment(params *UpdateAssignmentParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAssignmentOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAssignment gets an assignment

  This will get an assignment
*/
func (a *Client) GetAssignment(params *GetAssignmentParams, authInfo runtime.ClientAuthInfoWriter) (*GetAssignmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAssignmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAssignment",
		Method:             "GET",
		PathPattern:        "/assignments/{Id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAssignmentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAssignmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAssignment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAssignmentLog gets assignment log

  This will get the assignment log
*/
func (a *Client) GetAssignmentLog(params *GetAssignmentLogParams, authInfo runtime.ClientAuthInfoWriter) (*GetAssignmentLogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAssignmentLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAssignmentLog",
		Method:             "GET",
		PathPattern:        "/assignments/{Id}/log",
		ProducesMediaTypes: []string{"application/text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAssignmentLogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAssignmentLogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAssignmentLog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAssignmentResult gets assignment result

  This will get the assignment result
*/
func (a *Client) GetAssignmentResult(params *GetAssignmentResultParams, authInfo runtime.ClientAuthInfoWriter) (*GetAssignmentResultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAssignmentResultParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAssignmentResult",
		Method:             "GET",
		PathPattern:        "/assignments/{Id}/result",
		ProducesMediaTypes: []string{"application/text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAssignmentResultReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAssignmentResultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAssignmentResult: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAssignments gets all assignments by default

  This will get all assignments
*/
func (a *Client) GetAssignments(params *GetAssignmentsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAssignmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAssignmentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAssignments",
		Method:             "GET",
		PathPattern:        "/assignments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAssignmentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAssignmentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAssignments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAssignment posts an assignment to the server

  This will post an assignment
*/
func (a *Client) PostAssignment(params *PostAssignmentParams, authInfo runtime.ClientAuthInfoWriter) (*PostAssignmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAssignmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postAssignment",
		Method:             "POST",
		PathPattern:        "/assignments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostAssignmentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAssignmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postAssignment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReportAssignmentLog reports assignment log

  This will report assignment log
*/
func (a *Client) ReportAssignmentLog(params *ReportAssignmentLogParams, authInfo runtime.ClientAuthInfoWriter) (*ReportAssignmentLogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReportAssignmentLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "reportAssignmentLog",
		Method:             "POST",
		PathPattern:        "/assignments/{Id}/log",
		ProducesMediaTypes: []string{"application/text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReportAssignmentLogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReportAssignmentLogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for reportAssignmentLog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReportAssignmentResult reports assignment result

  This will report assignment result
*/
func (a *Client) ReportAssignmentResult(params *ReportAssignmentResultParams, authInfo runtime.ClientAuthInfoWriter) (*ReportAssignmentResultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReportAssignmentResultParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "reportAssignmentResult",
		Method:             "POST",
		PathPattern:        "/assignments/{Id}/result",
		ProducesMediaTypes: []string{"application/text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReportAssignmentResultReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReportAssignmentResultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for reportAssignmentResult: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StopAssignment stops an assignment

  This will stop an assignment
*/
func (a *Client) StopAssignment(params *StopAssignmentParams, authInfo runtime.ClientAuthInfoWriter) (*StopAssignmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopAssignmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stopAssignment",
		Method:             "DELETE",
		PathPattern:        "/assignments/{Id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopAssignmentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StopAssignmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for stopAssignment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateAssignment updates an assignment

  This will update an assignment
*/
func (a *Client) UpdateAssignment(params *UpdateAssignmentParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAssignmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAssignmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateAssignment",
		Method:             "PATCH",
		PathPattern:        "/assignments/{Id}",
		ProducesMediaTypes: []string{"application/text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateAssignmentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateAssignmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAssignment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
