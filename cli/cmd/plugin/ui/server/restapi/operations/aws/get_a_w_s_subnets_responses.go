// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware-tanzu/community-edition/cli/cmd/plugin/ui/server/models"
)

// GetAWSSubnetsOKCode is the HTTP code returned for type GetAWSSubnetsOK
const GetAWSSubnetsOKCode int = 200

/*GetAWSSubnetsOK Successful retrieval of AWS subnets

swagger:response getAWSSubnetsOK
*/
type GetAWSSubnetsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.AWSSubnet `json:"body,omitempty"`
}

// NewGetAWSSubnetsOK creates GetAWSSubnetsOK with default headers values
func NewGetAWSSubnetsOK() *GetAWSSubnetsOK {

	return &GetAWSSubnetsOK{}
}

// WithPayload adds the payload to the get a w s subnets o k response
func (o *GetAWSSubnetsOK) WithPayload(payload []*models.AWSSubnet) *GetAWSSubnetsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get a w s subnets o k response
func (o *GetAWSSubnetsOK) SetPayload(payload []*models.AWSSubnet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAWSSubnetsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.AWSSubnet, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAWSSubnetsBadRequestCode is the HTTP code returned for type GetAWSSubnetsBadRequest
const GetAWSSubnetsBadRequestCode int = 400

/*GetAWSSubnetsBadRequest Bad request

swagger:response getAWSSubnetsBadRequest
*/
type GetAWSSubnetsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAWSSubnetsBadRequest creates GetAWSSubnetsBadRequest with default headers values
func NewGetAWSSubnetsBadRequest() *GetAWSSubnetsBadRequest {

	return &GetAWSSubnetsBadRequest{}
}

// WithPayload adds the payload to the get a w s subnets bad request response
func (o *GetAWSSubnetsBadRequest) WithPayload(payload *models.Error) *GetAWSSubnetsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get a w s subnets bad request response
func (o *GetAWSSubnetsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAWSSubnetsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAWSSubnetsUnauthorizedCode is the HTTP code returned for type GetAWSSubnetsUnauthorized
const GetAWSSubnetsUnauthorizedCode int = 401

/*GetAWSSubnetsUnauthorized Incorrect credentials

swagger:response getAWSSubnetsUnauthorized
*/
type GetAWSSubnetsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAWSSubnetsUnauthorized creates GetAWSSubnetsUnauthorized with default headers values
func NewGetAWSSubnetsUnauthorized() *GetAWSSubnetsUnauthorized {

	return &GetAWSSubnetsUnauthorized{}
}

// WithPayload adds the payload to the get a w s subnets unauthorized response
func (o *GetAWSSubnetsUnauthorized) WithPayload(payload *models.Error) *GetAWSSubnetsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get a w s subnets unauthorized response
func (o *GetAWSSubnetsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAWSSubnetsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAWSSubnetsInternalServerErrorCode is the HTTP code returned for type GetAWSSubnetsInternalServerError
const GetAWSSubnetsInternalServerErrorCode int = 500

/*GetAWSSubnetsInternalServerError Internal server error

swagger:response getAWSSubnetsInternalServerError
*/
type GetAWSSubnetsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAWSSubnetsInternalServerError creates GetAWSSubnetsInternalServerError with default headers values
func NewGetAWSSubnetsInternalServerError() *GetAWSSubnetsInternalServerError {

	return &GetAWSSubnetsInternalServerError{}
}

// WithPayload adds the payload to the get a w s subnets internal server error response
func (o *GetAWSSubnetsInternalServerError) WithPayload(payload *models.Error) *GetAWSSubnetsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get a w s subnets internal server error response
func (o *GetAWSSubnetsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAWSSubnetsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
