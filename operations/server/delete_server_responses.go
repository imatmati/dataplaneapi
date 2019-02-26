// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// DeleteServerAcceptedCode is the HTTP code returned for type DeleteServerAccepted
const DeleteServerAcceptedCode int = 202

/*DeleteServerAccepted Configuration change accepted and reload requested

swagger:response deleteServerAccepted
*/
type DeleteServerAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteServerAccepted creates DeleteServerAccepted with default headers values
func NewDeleteServerAccepted() *DeleteServerAccepted {

	return &DeleteServerAccepted{}
}

// WithReloadID adds the reloadId to the delete server accepted response
func (o *DeleteServerAccepted) WithReloadID(reloadID string) *DeleteServerAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete server accepted response
func (o *DeleteServerAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteServerAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteServerNoContentCode is the HTTP code returned for type DeleteServerNoContent
const DeleteServerNoContentCode int = 204

/*DeleteServerNoContent Server deleted

swagger:response deleteServerNoContent
*/
type DeleteServerNoContent struct {
}

// NewDeleteServerNoContent creates DeleteServerNoContent with default headers values
func NewDeleteServerNoContent() *DeleteServerNoContent {

	return &DeleteServerNoContent{}
}

// WriteResponse to the client
func (o *DeleteServerNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteServerNotFoundCode is the HTTP code returned for type DeleteServerNotFound
const DeleteServerNotFoundCode int = 404

/*DeleteServerNotFound The specified resource was not found

swagger:response deleteServerNotFound
*/
type DeleteServerNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteServerNotFound creates DeleteServerNotFound with default headers values
func NewDeleteServerNotFound() *DeleteServerNotFound {

	return &DeleteServerNotFound{}
}

// WithPayload adds the payload to the delete server not found response
func (o *DeleteServerNotFound) WithPayload(payload *models.Error) *DeleteServerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete server not found response
func (o *DeleteServerNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteServerDefault General Error

swagger:response deleteServerDefault
*/
type DeleteServerDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteServerDefault creates DeleteServerDefault with default headers values
func NewDeleteServerDefault(code int) *DeleteServerDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteServerDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete server default response
func (o *DeleteServerDefault) WithStatusCode(code int) *DeleteServerDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete server default response
func (o *DeleteServerDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete server default response
func (o *DeleteServerDefault) WithPayload(payload *models.Error) *DeleteServerDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete server default response
func (o *DeleteServerDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServerDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
