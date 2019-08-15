// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/charliekenney23/bitrisectl/bitrise/models"
)

// UserPlanReader is a Reader for the UserPlan structure.
type UserPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUserPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUserPlanUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUserPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUserPlanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserPlanOK creates a UserPlanOK with default headers values
func NewUserPlanOK() *UserPlanOK {
	return &UserPlanOK{}
}

/*UserPlanOK handles this case with default header values.

OK
*/
type UserPlanOK struct {
	Payload *models.V0UserPlanRespModel
}

func (o *UserPlanOK) Error() string {
	return fmt.Sprintf("[GET /me/plan][%d] userPlanOK  %+v", 200, o.Payload)
}

func (o *UserPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0UserPlanRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserPlanBadRequest creates a UserPlanBadRequest with default headers values
func NewUserPlanBadRequest() *UserPlanBadRequest {
	return &UserPlanBadRequest{}
}

/*UserPlanBadRequest handles this case with default header values.

Bad Request
*/
type UserPlanBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *UserPlanBadRequest) Error() string {
	return fmt.Sprintf("[GET /me/plan][%d] userPlanBadRequest  %+v", 400, o.Payload)
}

func (o *UserPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserPlanUnauthorized creates a UserPlanUnauthorized with default headers values
func NewUserPlanUnauthorized() *UserPlanUnauthorized {
	return &UserPlanUnauthorized{}
}

/*UserPlanUnauthorized handles this case with default header values.

Unauthorized
*/
type UserPlanUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *UserPlanUnauthorized) Error() string {
	return fmt.Sprintf("[GET /me/plan][%d] userPlanUnauthorized  %+v", 401, o.Payload)
}

func (o *UserPlanUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserPlanNotFound creates a UserPlanNotFound with default headers values
func NewUserPlanNotFound() *UserPlanNotFound {
	return &UserPlanNotFound{}
}

/*UserPlanNotFound handles this case with default header values.

Not Found
*/
type UserPlanNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *UserPlanNotFound) Error() string {
	return fmt.Sprintf("[GET /me/plan][%d] userPlanNotFound  %+v", 404, o.Payload)
}

func (o *UserPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserPlanInternalServerError creates a UserPlanInternalServerError with default headers values
func NewUserPlanInternalServerError() *UserPlanInternalServerError {
	return &UserPlanInternalServerError{}
}

/*UserPlanInternalServerError handles this case with default header values.

Internal Server Error
*/
type UserPlanInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *UserPlanInternalServerError) Error() string {
	return fmt.Sprintf("[GET /me/plan][%d] userPlanInternalServerError  %+v", 500, o.Payload)
}

func (o *UserPlanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
