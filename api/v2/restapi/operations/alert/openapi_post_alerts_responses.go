// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package alert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// OpenapiPostAlertsOKCode is the HTTP code returned for type OpenapiPostAlertsOK
const OpenapiPostAlertsOKCode int = 200

/*OpenapiPostAlertsOK Create alerts response

swagger:response openapiPostAlertsOK
*/
type OpenapiPostAlertsOK struct {
}

// NewOpenapiPostAlertsOK creates OpenapiPostAlertsOK with default headers values
func NewOpenapiPostAlertsOK() *OpenapiPostAlertsOK {

	return &OpenapiPostAlertsOK{}
}

// WriteResponse to the client
func (o *OpenapiPostAlertsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// OpenapiPostAlertsBadRequestCode is the HTTP code returned for type OpenapiPostAlertsBadRequest
const OpenapiPostAlertsBadRequestCode int = 400

/*OpenapiPostAlertsBadRequest Bad request

swagger:response openapiPostAlertsBadRequest
*/
type OpenapiPostAlertsBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewOpenapiPostAlertsBadRequest creates OpenapiPostAlertsBadRequest with default headers values
func NewOpenapiPostAlertsBadRequest() *OpenapiPostAlertsBadRequest {

	return &OpenapiPostAlertsBadRequest{}
}

// WithPayload adds the payload to the openapi post alerts bad request response
func (o *OpenapiPostAlertsBadRequest) WithPayload(payload string) *OpenapiPostAlertsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the openapi post alerts bad request response
func (o *OpenapiPostAlertsBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *OpenapiPostAlertsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// OpenapiPostAlertsInternalServerErrorCode is the HTTP code returned for type OpenapiPostAlertsInternalServerError
const OpenapiPostAlertsInternalServerErrorCode int = 500

/*OpenapiPostAlertsInternalServerError Internal server error

swagger:response openapiPostAlertsInternalServerError
*/
type OpenapiPostAlertsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewOpenapiPostAlertsInternalServerError creates OpenapiPostAlertsInternalServerError with default headers values
func NewOpenapiPostAlertsInternalServerError() *OpenapiPostAlertsInternalServerError {

	return &OpenapiPostAlertsInternalServerError{}
}

// WithPayload adds the payload to the openapi post alerts internal server error response
func (o *OpenapiPostAlertsInternalServerError) WithPayload(payload string) *OpenapiPostAlertsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the openapi post alerts internal server error response
func (o *OpenapiPostAlertsInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *OpenapiPostAlertsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
