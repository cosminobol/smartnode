// Code generated by go-swagger; DO NOT EDIT.

package general

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rocket-pool/smartnode/shared/services/alerting/alertmanager/models"
)

// GetStatusReader is a Reader for the GetStatus structure.
type GetStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /status] getStatus", response, response.Code())
	}
}

// NewGetStatusOK creates a GetStatusOK with default headers values
func NewGetStatusOK() *GetStatusOK {
	return &GetStatusOK{}
}

/*
GetStatusOK describes a response with status code 200, with default header values.

Get status response
*/
type GetStatusOK struct {
	Payload *models.AlertmanagerStatus
}

// IsSuccess returns true when this get status o k response has a 2xx status code
func (o *GetStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get status o k response has a 3xx status code
func (o *GetStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get status o k response has a 4xx status code
func (o *GetStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get status o k response has a 5xx status code
func (o *GetStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get status o k response a status code equal to that given
func (o *GetStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get status o k response
func (o *GetStatusOK) Code() int {
	return 200
}

func (o *GetStatusOK) Error() string {
	return fmt.Sprintf("[GET /status][%d] getStatusOK  %+v", 200, o.Payload)
}

func (o *GetStatusOK) String() string {
	return fmt.Sprintf("[GET /status][%d] getStatusOK  %+v", 200, o.Payload)
}

func (o *GetStatusOK) GetPayload() *models.AlertmanagerStatus {
	return o.Payload
}

func (o *GetStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertmanagerStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
