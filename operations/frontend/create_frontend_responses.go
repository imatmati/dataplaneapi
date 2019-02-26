// Code generated by go-swagger; DO NOT EDIT.

package frontend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// CreateFrontendCreatedCode is the HTTP code returned for type CreateFrontendCreated
const CreateFrontendCreatedCode int = 201

/*CreateFrontendCreated Frontend created

swagger:response createFrontendCreated
*/
type CreateFrontendCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Frontend `json:"body,omitempty"`
}

// NewCreateFrontendCreated creates CreateFrontendCreated with default headers values
func NewCreateFrontendCreated() *CreateFrontendCreated {

	return &CreateFrontendCreated{}
}

// WithPayload adds the payload to the create frontend created response
func (o *CreateFrontendCreated) WithPayload(payload *models.Frontend) *CreateFrontendCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create frontend created response
func (o *CreateFrontendCreated) SetPayload(payload *models.Frontend) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFrontendCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateFrontendAcceptedCode is the HTTP code returned for type CreateFrontendAccepted
const CreateFrontendAcceptedCode int = 202

/*CreateFrontendAccepted Configuration change accepted and reload requested

swagger:response createFrontendAccepted
*/
type CreateFrontendAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Frontend `json:"body,omitempty"`
}

// NewCreateFrontendAccepted creates CreateFrontendAccepted with default headers values
func NewCreateFrontendAccepted() *CreateFrontendAccepted {

	return &CreateFrontendAccepted{}
}

// WithReloadID adds the reloadId to the create frontend accepted response
func (o *CreateFrontendAccepted) WithReloadID(reloadID string) *CreateFrontendAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create frontend accepted response
func (o *CreateFrontendAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create frontend accepted response
func (o *CreateFrontendAccepted) WithPayload(payload *models.Frontend) *CreateFrontendAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create frontend accepted response
func (o *CreateFrontendAccepted) SetPayload(payload *models.Frontend) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFrontendAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateFrontendBadRequestCode is the HTTP code returned for type CreateFrontendBadRequest
const CreateFrontendBadRequestCode int = 400

/*CreateFrontendBadRequest Bad request

swagger:response createFrontendBadRequest
*/
type CreateFrontendBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateFrontendBadRequest creates CreateFrontendBadRequest with default headers values
func NewCreateFrontendBadRequest() *CreateFrontendBadRequest {

	return &CreateFrontendBadRequest{}
}

// WithPayload adds the payload to the create frontend bad request response
func (o *CreateFrontendBadRequest) WithPayload(payload *models.Error) *CreateFrontendBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create frontend bad request response
func (o *CreateFrontendBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFrontendBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateFrontendConflictCode is the HTTP code returned for type CreateFrontendConflict
const CreateFrontendConflictCode int = 409

/*CreateFrontendConflict The specified resource already exists

swagger:response createFrontendConflict
*/
type CreateFrontendConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateFrontendConflict creates CreateFrontendConflict with default headers values
func NewCreateFrontendConflict() *CreateFrontendConflict {

	return &CreateFrontendConflict{}
}

// WithPayload adds the payload to the create frontend conflict response
func (o *CreateFrontendConflict) WithPayload(payload *models.Error) *CreateFrontendConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create frontend conflict response
func (o *CreateFrontendConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFrontendConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateFrontendDefault General Error

swagger:response createFrontendDefault
*/
type CreateFrontendDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateFrontendDefault creates CreateFrontendDefault with default headers values
func NewCreateFrontendDefault(code int) *CreateFrontendDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateFrontendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create frontend default response
func (o *CreateFrontendDefault) WithStatusCode(code int) *CreateFrontendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create frontend default response
func (o *CreateFrontendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create frontend default response
func (o *CreateFrontendDefault) WithPayload(payload *models.Error) *CreateFrontendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create frontend default response
func (o *CreateFrontendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFrontendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
