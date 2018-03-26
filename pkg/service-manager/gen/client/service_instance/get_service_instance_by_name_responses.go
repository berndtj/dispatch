// Code generated by go-swagger; DO NOT EDIT.

package service_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/service-manager/gen/models"
)

// GetServiceInstanceByNameReader is a Reader for the GetServiceInstanceByName structure.
type GetServiceInstanceByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceInstanceByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServiceInstanceByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetServiceInstanceByNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetServiceInstanceByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetServiceInstanceByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServiceInstanceByNameOK creates a GetServiceInstanceByNameOK with default headers values
func NewGetServiceInstanceByNameOK() *GetServiceInstanceByNameOK {
	return &GetServiceInstanceByNameOK{}
}

/*GetServiceInstanceByNameOK handles this case with default header values.

successful operation
*/
type GetServiceInstanceByNameOK struct {
	Payload *models.ServiceInstance
}

func (o *GetServiceInstanceByNameOK) Error() string {
	return fmt.Sprintf("[GET /serviceinstance/{serviceInstanceName}][%d] getServiceInstanceByNameOK  %+v", 200, o.Payload)
}

func (o *GetServiceInstanceByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceInstanceByNameBadRequest creates a GetServiceInstanceByNameBadRequest with default headers values
func NewGetServiceInstanceByNameBadRequest() *GetServiceInstanceByNameBadRequest {
	return &GetServiceInstanceByNameBadRequest{}
}

/*GetServiceInstanceByNameBadRequest handles this case with default header values.

Invalid ID supplied
*/
type GetServiceInstanceByNameBadRequest struct {
	Payload *models.Error
}

func (o *GetServiceInstanceByNameBadRequest) Error() string {
	return fmt.Sprintf("[GET /serviceinstance/{serviceInstanceName}][%d] getServiceInstanceByNameBadRequest  %+v", 400, o.Payload)
}

func (o *GetServiceInstanceByNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceInstanceByNameNotFound creates a GetServiceInstanceByNameNotFound with default headers values
func NewGetServiceInstanceByNameNotFound() *GetServiceInstanceByNameNotFound {
	return &GetServiceInstanceByNameNotFound{}
}

/*GetServiceInstanceByNameNotFound handles this case with default header values.

Service instance not found
*/
type GetServiceInstanceByNameNotFound struct {
	Payload *models.Error
}

func (o *GetServiceInstanceByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /serviceinstance/{serviceInstanceName}][%d] getServiceInstanceByNameNotFound  %+v", 404, o.Payload)
}

func (o *GetServiceInstanceByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceInstanceByNameDefault creates a GetServiceInstanceByNameDefault with default headers values
func NewGetServiceInstanceByNameDefault(code int) *GetServiceInstanceByNameDefault {
	return &GetServiceInstanceByNameDefault{
		_statusCode: code,
	}
}

/*GetServiceInstanceByNameDefault handles this case with default header values.

Generic error response
*/
type GetServiceInstanceByNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get service instance by name default response
func (o *GetServiceInstanceByNameDefault) Code() int {
	return o._statusCode
}

func (o *GetServiceInstanceByNameDefault) Error() string {
	return fmt.Sprintf("[GET /serviceinstance/{serviceInstanceName}][%d] getServiceInstanceByName default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceInstanceByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
