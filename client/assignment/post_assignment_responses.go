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

// PostAssignmentReader is a Reader for the PostAssignment structure.
type PostAssignmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAssignmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAssignmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAssignmentOK creates a PostAssignmentOK with default headers values
func NewPostAssignmentOK() *PostAssignmentOK {
	return &PostAssignmentOK{}
}

/*PostAssignmentOK handles this case with default header values.

Assignment meta info
*/
type PostAssignmentOK struct {
	Payload *models.Meta
}

func (o *PostAssignmentOK) Error() string {
	return fmt.Sprintf("[POST /assignments][%d] postAssignmentOK  %+v", 200, o.Payload)
}

func (o *PostAssignmentOK) GetPayload() *models.Meta {
	return o.Payload
}

func (o *PostAssignmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Meta)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
