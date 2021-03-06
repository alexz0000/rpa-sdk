// Code generated by go-swagger; DO NOT EDIT.

package assignment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"rpa-sdk/models"
)

// StopAssignmentReader is a Reader for the StopAssignment structure.
type StopAssignmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopAssignmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopAssignmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStopAssignmentOK creates a StopAssignmentOK with default headers values
func NewStopAssignmentOK() *StopAssignmentOK {
	return &StopAssignmentOK{}
}

/*StopAssignmentOK handles this case with default header values.

Assignment meta info
*/
type StopAssignmentOK struct {
	Payload *models.Meta
}

func (o *StopAssignmentOK) Error() string {
	return fmt.Sprintf("[DELETE /assignments/{Id}][%d] stopAssignmentOK  %+v", 200, o.Payload)
}

func (o *StopAssignmentOK) GetPayload() *models.Meta {
	return o.Payload
}

func (o *StopAssignmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Meta)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
