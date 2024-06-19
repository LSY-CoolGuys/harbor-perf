// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
)

// GetRepositoryReader is a Reader for the GetRepository structure.
type GetRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRepositoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRepositoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRepositoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /projects/{project_name}/repositories/{repository_name}] getRepository", response, response.Code())
	}
}

// NewGetRepositoryOK creates a GetRepositoryOK with default headers values
func NewGetRepositoryOK() *GetRepositoryOK {
	return &GetRepositoryOK{}
}

/*
GetRepositoryOK describes a response with status code 200, with default header values.

Success
*/
type GetRepositoryOK struct {
	Payload *models.Repository
}

// IsSuccess returns true when this get repository o k response has a 2xx status code
func (o *GetRepositoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repository o k response has a 3xx status code
func (o *GetRepositoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository o k response has a 4xx status code
func (o *GetRepositoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository o k response has a 5xx status code
func (o *GetRepositoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository o k response a status code equal to that given
func (o *GetRepositoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repository o k response
func (o *GetRepositoryOK) Code() int {
	return 200
}

func (o *GetRepositoryOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}][%d] getRepositoryOK  %+v", 200, o.Payload)
}

func (o *GetRepositoryOK) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}][%d] getRepositoryOK  %+v", 200, o.Payload)
}

func (o *GetRepositoryOK) GetPayload() *models.Repository {
	return o.Payload
}

func (o *GetRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Repository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoryBadRequest creates a GetRepositoryBadRequest with default headers values
func NewGetRepositoryBadRequest() *GetRepositoryBadRequest {
	return &GetRepositoryBadRequest{}
}

/*
GetRepositoryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetRepositoryBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get repository bad request response has a 2xx status code
func (o *GetRepositoryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get repository bad request response has a 3xx status code
func (o *GetRepositoryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository bad request response has a 4xx status code
func (o *GetRepositoryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get repository bad request response has a 5xx status code
func (o *GetRepositoryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository bad request response a status code equal to that given
func (o *GetRepositoryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get repository bad request response
func (o *GetRepositoryBadRequest) Code() int {
	return 400
}

func (o *GetRepositoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}][%d] getRepositoryBadRequest  %+v", 400, o.Payload)
}

func (o *GetRepositoryBadRequest) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}][%d] getRepositoryBadRequest  %+v", 400, o.Payload)
}

func (o *GetRepositoryBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRepositoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoryUnauthorized creates a GetRepositoryUnauthorized with default headers values
func NewGetRepositoryUnauthorized() *GetRepositoryUnauthorized {
	return &GetRepositoryUnauthorized{}
}

/*
GetRepositoryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRepositoryUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get repository unauthorized response has a 2xx status code
func (o *GetRepositoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get repository unauthorized response has a 3xx status code
func (o *GetRepositoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository unauthorized response has a 4xx status code
func (o *GetRepositoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get repository unauthorized response has a 5xx status code
func (o *GetRepositoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository unauthorized response a status code equal to that given
func (o *GetRepositoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get repository unauthorized response
func (o *GetRepositoryUnauthorized) Code() int {
	return 401
}

func (o *GetRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}][%d] getRepositoryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRepositoryUnauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}][%d] getRepositoryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRepositoryUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoryForbidden creates a GetRepositoryForbidden with default headers values
func NewGetRepositoryForbidden() *GetRepositoryForbidden {
	return &GetRepositoryForbidden{}
}

/*
GetRepositoryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRepositoryForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get repository forbidden response has a 2xx status code
func (o *GetRepositoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get repository forbidden response has a 3xx status code
func (o *GetRepositoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository forbidden response has a 4xx status code
func (o *GetRepositoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get repository forbidden response has a 5xx status code
func (o *GetRepositoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository forbidden response a status code equal to that given
func (o *GetRepositoryForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get repository forbidden response
func (o *GetRepositoryForbidden) Code() int {
	return 403
}

func (o *GetRepositoryForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}][%d] getRepositoryForbidden  %+v", 403, o.Payload)
}

func (o *GetRepositoryForbidden) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}][%d] getRepositoryForbidden  %+v", 403, o.Payload)
}

func (o *GetRepositoryForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoryNotFound creates a GetRepositoryNotFound with default headers values
func NewGetRepositoryNotFound() *GetRepositoryNotFound {
	return &GetRepositoryNotFound{}
}

/*
GetRepositoryNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetRepositoryNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get repository not found response has a 2xx status code
func (o *GetRepositoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get repository not found response has a 3xx status code
func (o *GetRepositoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository not found response has a 4xx status code
func (o *GetRepositoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get repository not found response has a 5xx status code
func (o *GetRepositoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get repository not found response a status code equal to that given
func (o *GetRepositoryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get repository not found response
func (o *GetRepositoryNotFound) Code() int {
	return 404
}

func (o *GetRepositoryNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}][%d] getRepositoryNotFound  %+v", 404, o.Payload)
}

func (o *GetRepositoryNotFound) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}][%d] getRepositoryNotFound  %+v", 404, o.Payload)
}

func (o *GetRepositoryNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRepositoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoryInternalServerError creates a GetRepositoryInternalServerError with default headers values
func NewGetRepositoryInternalServerError() *GetRepositoryInternalServerError {
	return &GetRepositoryInternalServerError{}
}

/*
GetRepositoryInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetRepositoryInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get repository internal server error response has a 2xx status code
func (o *GetRepositoryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get repository internal server error response has a 3xx status code
func (o *GetRepositoryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repository internal server error response has a 4xx status code
func (o *GetRepositoryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repository internal server error response has a 5xx status code
func (o *GetRepositoryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get repository internal server error response a status code equal to that given
func (o *GetRepositoryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get repository internal server error response
func (o *GetRepositoryInternalServerError) Code() int {
	return 500
}

func (o *GetRepositoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}][%d] getRepositoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRepositoryInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}][%d] getRepositoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRepositoryInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetRepositoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
