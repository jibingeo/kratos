// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/kratos/internal/httpclient/models"
)

// InitializeSelfServiceRecoveryFlowReader is a Reader for the InitializeSelfServiceRecoveryFlow structure.
type InitializeSelfServiceRecoveryFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InitializeSelfServiceRecoveryFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewInitializeSelfServiceRecoveryFlowFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInitializeSelfServiceRecoveryFlowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInitializeSelfServiceRecoveryFlowFound creates a InitializeSelfServiceRecoveryFlowFound with default headers values
func NewInitializeSelfServiceRecoveryFlowFound() *InitializeSelfServiceRecoveryFlowFound {
	return &InitializeSelfServiceRecoveryFlowFound{}
}

/*InitializeSelfServiceRecoveryFlowFound handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type InitializeSelfServiceRecoveryFlowFound struct {
}

func (o *InitializeSelfServiceRecoveryFlowFound) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/recovery][%d] initializeSelfServiceRecoveryFlowFound ", 302)
}

func (o *InitializeSelfServiceRecoveryFlowFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitializeSelfServiceRecoveryFlowInternalServerError creates a InitializeSelfServiceRecoveryFlowInternalServerError with default headers values
func NewInitializeSelfServiceRecoveryFlowInternalServerError() *InitializeSelfServiceRecoveryFlowInternalServerError {
	return &InitializeSelfServiceRecoveryFlowInternalServerError{}
}

/*InitializeSelfServiceRecoveryFlowInternalServerError handles this case with default header values.

genericError
*/
type InitializeSelfServiceRecoveryFlowInternalServerError struct {
	Payload *models.GenericError
}

func (o *InitializeSelfServiceRecoveryFlowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/recovery][%d] initializeSelfServiceRecoveryFlowInternalServerError  %+v", 500, o.Payload)
}

func (o *InitializeSelfServiceRecoveryFlowInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *InitializeSelfServiceRecoveryFlowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}