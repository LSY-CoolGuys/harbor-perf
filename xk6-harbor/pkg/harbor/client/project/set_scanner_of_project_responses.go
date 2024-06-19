// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
)

// SetScannerOfProjectReader is a Reader for the SetScannerOfProject structure.
type SetScannerOfProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetScannerOfProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetScannerOfProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetScannerOfProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSetScannerOfProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetScannerOfProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetScannerOfProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetScannerOfProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /projects/{project_name_or_id}/scanner] setScannerOfProject", response, response.Code())
	}
}

// NewSetScannerOfProjectOK creates a SetScannerOfProjectOK with default headers values
func NewSetScannerOfProjectOK() *SetScannerOfProjectOK {
	return &SetScannerOfProjectOK{}
}

/*
SetScannerOfProjectOK describes a response with status code 200, with default header values.

Success
*/
type SetScannerOfProjectOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this set scanner of project o k response has a 2xx status code
func (o *SetScannerOfProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set scanner of project o k response has a 3xx status code
func (o *SetScannerOfProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set scanner of project o k response has a 4xx status code
func (o *SetScannerOfProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set scanner of project o k response has a 5xx status code
func (o *SetScannerOfProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set scanner of project o k response a status code equal to that given
func (o *SetScannerOfProjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set scanner of project o k response
func (o *SetScannerOfProjectOK) Code() int {
	return 200
}

func (o *SetScannerOfProjectOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/scanner][%d] setScannerOfProjectOK ", 200)
}

func (o *SetScannerOfProjectOK) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/scanner][%d] setScannerOfProjectOK ", 200)
}

func (o *SetScannerOfProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewSetScannerOfProjectBadRequest creates a SetScannerOfProjectBadRequest with default headers values
func NewSetScannerOfProjectBadRequest() *SetScannerOfProjectBadRequest {
	return &SetScannerOfProjectBadRequest{}
}

/*
SetScannerOfProjectBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SetScannerOfProjectBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this set scanner of project bad request response has a 2xx status code
func (o *SetScannerOfProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set scanner of project bad request response has a 3xx status code
func (o *SetScannerOfProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set scanner of project bad request response has a 4xx status code
func (o *SetScannerOfProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this set scanner of project bad request response has a 5xx status code
func (o *SetScannerOfProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this set scanner of project bad request response a status code equal to that given
func (o *SetScannerOfProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the set scanner of project bad request response
func (o *SetScannerOfProjectBadRequest) Code() int {
	return 400
}

func (o *SetScannerOfProjectBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/scanner][%d] setScannerOfProjectBadRequest  %+v", 400, o.Payload)
}

func (o *SetScannerOfProjectBadRequest) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/scanner][%d] setScannerOfProjectBadRequest  %+v", 400, o.Payload)
}

func (o *SetScannerOfProjectBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SetScannerOfProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetScannerOfProjectUnauthorized creates a SetScannerOfProjectUnauthorized with default headers values
func NewSetScannerOfProjectUnauthorized() *SetScannerOfProjectUnauthorized {
	return &SetScannerOfProjectUnauthorized{}
}

/*
SetScannerOfProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SetScannerOfProjectUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this set scanner of project unauthorized response has a 2xx status code
func (o *SetScannerOfProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set scanner of project unauthorized response has a 3xx status code
func (o *SetScannerOfProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set scanner of project unauthorized response has a 4xx status code
func (o *SetScannerOfProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this set scanner of project unauthorized response has a 5xx status code
func (o *SetScannerOfProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this set scanner of project unauthorized response a status code equal to that given
func (o *SetScannerOfProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the set scanner of project unauthorized response
func (o *SetScannerOfProjectUnauthorized) Code() int {
	return 401
}

func (o *SetScannerOfProjectUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/scanner][%d] setScannerOfProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *SetScannerOfProjectUnauthorized) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/scanner][%d] setScannerOfProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *SetScannerOfProjectUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SetScannerOfProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetScannerOfProjectForbidden creates a SetScannerOfProjectForbidden with default headers values
func NewSetScannerOfProjectForbidden() *SetScannerOfProjectForbidden {
	return &SetScannerOfProjectForbidden{}
}

/*
SetScannerOfProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SetScannerOfProjectForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this set scanner of project forbidden response has a 2xx status code
func (o *SetScannerOfProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set scanner of project forbidden response has a 3xx status code
func (o *SetScannerOfProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set scanner of project forbidden response has a 4xx status code
func (o *SetScannerOfProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this set scanner of project forbidden response has a 5xx status code
func (o *SetScannerOfProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this set scanner of project forbidden response a status code equal to that given
func (o *SetScannerOfProjectForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the set scanner of project forbidden response
func (o *SetScannerOfProjectForbidden) Code() int {
	return 403
}

func (o *SetScannerOfProjectForbidden) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/scanner][%d] setScannerOfProjectForbidden  %+v", 403, o.Payload)
}

func (o *SetScannerOfProjectForbidden) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/scanner][%d] setScannerOfProjectForbidden  %+v", 403, o.Payload)
}

func (o *SetScannerOfProjectForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SetScannerOfProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetScannerOfProjectNotFound creates a SetScannerOfProjectNotFound with default headers values
func NewSetScannerOfProjectNotFound() *SetScannerOfProjectNotFound {
	return &SetScannerOfProjectNotFound{}
}

/*
SetScannerOfProjectNotFound describes a response with status code 404, with default header values.

Not found
*/
type SetScannerOfProjectNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this set scanner of project not found response has a 2xx status code
func (o *SetScannerOfProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set scanner of project not found response has a 3xx status code
func (o *SetScannerOfProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set scanner of project not found response has a 4xx status code
func (o *SetScannerOfProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this set scanner of project not found response has a 5xx status code
func (o *SetScannerOfProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this set scanner of project not found response a status code equal to that given
func (o *SetScannerOfProjectNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the set scanner of project not found response
func (o *SetScannerOfProjectNotFound) Code() int {
	return 404
}

func (o *SetScannerOfProjectNotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/scanner][%d] setScannerOfProjectNotFound  %+v", 404, o.Payload)
}

func (o *SetScannerOfProjectNotFound) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/scanner][%d] setScannerOfProjectNotFound  %+v", 404, o.Payload)
}

func (o *SetScannerOfProjectNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SetScannerOfProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetScannerOfProjectInternalServerError creates a SetScannerOfProjectInternalServerError with default headers values
func NewSetScannerOfProjectInternalServerError() *SetScannerOfProjectInternalServerError {
	return &SetScannerOfProjectInternalServerError{}
}

/*
SetScannerOfProjectInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SetScannerOfProjectInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this set scanner of project internal server error response has a 2xx status code
func (o *SetScannerOfProjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set scanner of project internal server error response has a 3xx status code
func (o *SetScannerOfProjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set scanner of project internal server error response has a 4xx status code
func (o *SetScannerOfProjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this set scanner of project internal server error response has a 5xx status code
func (o *SetScannerOfProjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this set scanner of project internal server error response a status code equal to that given
func (o *SetScannerOfProjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the set scanner of project internal server error response
func (o *SetScannerOfProjectInternalServerError) Code() int {
	return 500
}

func (o *SetScannerOfProjectInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/scanner][%d] setScannerOfProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *SetScannerOfProjectInternalServerError) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/scanner][%d] setScannerOfProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *SetScannerOfProjectInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SetScannerOfProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
