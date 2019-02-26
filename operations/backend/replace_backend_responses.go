// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// ReplaceBackendOKCode is the HTTP code returned for type ReplaceBackendOK
const ReplaceBackendOKCode int = 200

/*ReplaceBackendOK Backend replaced

swagger:response replaceBackendOK
*/
type ReplaceBackendOK struct {

	/*
	  In: Body
	*/
	Payload *models.Backend `json:"body,omitempty"`
}

// NewReplaceBackendOK creates ReplaceBackendOK with default headers values
func NewReplaceBackendOK() *ReplaceBackendOK {

	return &ReplaceBackendOK{}
}

// WithPayload adds the payload to the replace backend o k response
func (o *ReplaceBackendOK) WithPayload(payload *models.Backend) *ReplaceBackendOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace backend o k response
func (o *ReplaceBackendOK) SetPayload(payload *models.Backend) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBackendOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceBackendAcceptedCode is the HTTP code returned for type ReplaceBackendAccepted
const ReplaceBackendAcceptedCode int = 202

/*ReplaceBackendAccepted Configuration change accepted and reload requested

swagger:response replaceBackendAccepted
*/
type ReplaceBackendAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Backend `json:"body,omitempty"`
}

// NewReplaceBackendAccepted creates ReplaceBackendAccepted with default headers values
func NewReplaceBackendAccepted() *ReplaceBackendAccepted {

	return &ReplaceBackendAccepted{}
}

// WithReloadID adds the reloadId to the replace backend accepted response
func (o *ReplaceBackendAccepted) WithReloadID(reloadID string) *ReplaceBackendAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace backend accepted response
func (o *ReplaceBackendAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace backend accepted response
func (o *ReplaceBackendAccepted) WithPayload(payload *models.Backend) *ReplaceBackendAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace backend accepted response
func (o *ReplaceBackendAccepted) SetPayload(payload *models.Backend) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBackendAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceBackendBadRequestCode is the HTTP code returned for type ReplaceBackendBadRequest
const ReplaceBackendBadRequestCode int = 400

/*ReplaceBackendBadRequest Bad request

swagger:response replaceBackendBadRequest
*/
type ReplaceBackendBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceBackendBadRequest creates ReplaceBackendBadRequest with default headers values
func NewReplaceBackendBadRequest() *ReplaceBackendBadRequest {

	return &ReplaceBackendBadRequest{}
}

// WithPayload adds the payload to the replace backend bad request response
func (o *ReplaceBackendBadRequest) WithPayload(payload *models.Error) *ReplaceBackendBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace backend bad request response
func (o *ReplaceBackendBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBackendBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceBackendNotFoundCode is the HTTP code returned for type ReplaceBackendNotFound
const ReplaceBackendNotFoundCode int = 404

/*ReplaceBackendNotFound The specified resource was not found

swagger:response replaceBackendNotFound
*/
type ReplaceBackendNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceBackendNotFound creates ReplaceBackendNotFound with default headers values
func NewReplaceBackendNotFound() *ReplaceBackendNotFound {

	return &ReplaceBackendNotFound{}
}

// WithPayload adds the payload to the replace backend not found response
func (o *ReplaceBackendNotFound) WithPayload(payload *models.Error) *ReplaceBackendNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace backend not found response
func (o *ReplaceBackendNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBackendNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ReplaceBackendDefault General Error

swagger:response replaceBackendDefault
*/
type ReplaceBackendDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceBackendDefault creates ReplaceBackendDefault with default headers values
func NewReplaceBackendDefault(code int) *ReplaceBackendDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceBackendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace backend default response
func (o *ReplaceBackendDefault) WithStatusCode(code int) *ReplaceBackendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace backend default response
func (o *ReplaceBackendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the replace backend default response
func (o *ReplaceBackendDefault) WithPayload(payload *models.Error) *ReplaceBackendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace backend default response
func (o *ReplaceBackendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBackendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
