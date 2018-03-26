// Code generated by go-swagger; DO NOT EDIT.

package service_class

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/service-manager/gen/models"
)

// GetServiceClassesReader is a Reader for the GetServiceClasses structure.
type GetServiceClassesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceClassesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServiceClassesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetServiceClassesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServiceClassesOK creates a GetServiceClassesOK with default headers values
func NewGetServiceClassesOK() *GetServiceClassesOK {
	return &GetServiceClassesOK{}
}

/*GetServiceClassesOK handles this case with default header values.

successful operation
*/
type GetServiceClassesOK struct {
	Payload models.GetServiceClassesOKBody
}

func (o *GetServiceClassesOK) Error() string {
	return fmt.Sprintf("[GET /serviceclass][%d] getServiceClassesOK  %+v", 200, o.Payload)
}

func (o *GetServiceClassesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceClassesDefault creates a GetServiceClassesDefault with default headers values
func NewGetServiceClassesDefault(code int) *GetServiceClassesDefault {
	return &GetServiceClassesDefault{
		_statusCode: code,
	}
}

/*GetServiceClassesDefault handles this case with default header values.

Generic error response
*/
type GetServiceClassesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get service classes default response
func (o *GetServiceClassesDefault) Code() int {
	return o._statusCode
}

func (o *GetServiceClassesDefault) Error() string {
	return fmt.Sprintf("[GET /serviceclass][%d] getServiceClasses default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceClassesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
