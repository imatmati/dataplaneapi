// Code generated by go-swagger; DO NOT EDIT.

package backend_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// ReplaceBackendSwitchingRuleOKCode is the HTTP code returned for type ReplaceBackendSwitchingRuleOK
const ReplaceBackendSwitchingRuleOKCode int = 200

/*ReplaceBackendSwitchingRuleOK Backend Switching Rule replaced

swagger:response replaceBackendSwitchingRuleOK
*/
type ReplaceBackendSwitchingRuleOK struct {

	/*
	  In: Body
	*/
	Payload *models.BackendSwitchingRule `json:"body,omitempty"`
}

// NewReplaceBackendSwitchingRuleOK creates ReplaceBackendSwitchingRuleOK with default headers values
func NewReplaceBackendSwitchingRuleOK() *ReplaceBackendSwitchingRuleOK {

	return &ReplaceBackendSwitchingRuleOK{}
}

// WithPayload adds the payload to the replace backend switching rule o k response
func (o *ReplaceBackendSwitchingRuleOK) WithPayload(payload *models.BackendSwitchingRule) *ReplaceBackendSwitchingRuleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace backend switching rule o k response
func (o *ReplaceBackendSwitchingRuleOK) SetPayload(payload *models.BackendSwitchingRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBackendSwitchingRuleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceBackendSwitchingRuleAcceptedCode is the HTTP code returned for type ReplaceBackendSwitchingRuleAccepted
const ReplaceBackendSwitchingRuleAcceptedCode int = 202

/*ReplaceBackendSwitchingRuleAccepted Configuration change accepted and reload requested

swagger:response replaceBackendSwitchingRuleAccepted
*/
type ReplaceBackendSwitchingRuleAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.BackendSwitchingRule `json:"body,omitempty"`
}

// NewReplaceBackendSwitchingRuleAccepted creates ReplaceBackendSwitchingRuleAccepted with default headers values
func NewReplaceBackendSwitchingRuleAccepted() *ReplaceBackendSwitchingRuleAccepted {

	return &ReplaceBackendSwitchingRuleAccepted{}
}

// WithReloadID adds the reloadId to the replace backend switching rule accepted response
func (o *ReplaceBackendSwitchingRuleAccepted) WithReloadID(reloadID string) *ReplaceBackendSwitchingRuleAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace backend switching rule accepted response
func (o *ReplaceBackendSwitchingRuleAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace backend switching rule accepted response
func (o *ReplaceBackendSwitchingRuleAccepted) WithPayload(payload *models.BackendSwitchingRule) *ReplaceBackendSwitchingRuleAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace backend switching rule accepted response
func (o *ReplaceBackendSwitchingRuleAccepted) SetPayload(payload *models.BackendSwitchingRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBackendSwitchingRuleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceBackendSwitchingRuleBadRequestCode is the HTTP code returned for type ReplaceBackendSwitchingRuleBadRequest
const ReplaceBackendSwitchingRuleBadRequestCode int = 400

/*ReplaceBackendSwitchingRuleBadRequest Bad request

swagger:response replaceBackendSwitchingRuleBadRequest
*/
type ReplaceBackendSwitchingRuleBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceBackendSwitchingRuleBadRequest creates ReplaceBackendSwitchingRuleBadRequest with default headers values
func NewReplaceBackendSwitchingRuleBadRequest() *ReplaceBackendSwitchingRuleBadRequest {

	return &ReplaceBackendSwitchingRuleBadRequest{}
}

// WithPayload adds the payload to the replace backend switching rule bad request response
func (o *ReplaceBackendSwitchingRuleBadRequest) WithPayload(payload *models.Error) *ReplaceBackendSwitchingRuleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace backend switching rule bad request response
func (o *ReplaceBackendSwitchingRuleBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBackendSwitchingRuleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceBackendSwitchingRuleNotFoundCode is the HTTP code returned for type ReplaceBackendSwitchingRuleNotFound
const ReplaceBackendSwitchingRuleNotFoundCode int = 404

/*ReplaceBackendSwitchingRuleNotFound The specified resource was not found

swagger:response replaceBackendSwitchingRuleNotFound
*/
type ReplaceBackendSwitchingRuleNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceBackendSwitchingRuleNotFound creates ReplaceBackendSwitchingRuleNotFound with default headers values
func NewReplaceBackendSwitchingRuleNotFound() *ReplaceBackendSwitchingRuleNotFound {

	return &ReplaceBackendSwitchingRuleNotFound{}
}

// WithPayload adds the payload to the replace backend switching rule not found response
func (o *ReplaceBackendSwitchingRuleNotFound) WithPayload(payload *models.Error) *ReplaceBackendSwitchingRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace backend switching rule not found response
func (o *ReplaceBackendSwitchingRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBackendSwitchingRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ReplaceBackendSwitchingRuleDefault General Error

swagger:response replaceBackendSwitchingRuleDefault
*/
type ReplaceBackendSwitchingRuleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceBackendSwitchingRuleDefault creates ReplaceBackendSwitchingRuleDefault with default headers values
func NewReplaceBackendSwitchingRuleDefault(code int) *ReplaceBackendSwitchingRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceBackendSwitchingRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace backend switching rule default response
func (o *ReplaceBackendSwitchingRuleDefault) WithStatusCode(code int) *ReplaceBackendSwitchingRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace backend switching rule default response
func (o *ReplaceBackendSwitchingRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the replace backend switching rule default response
func (o *ReplaceBackendSwitchingRuleDefault) WithPayload(payload *models.Error) *ReplaceBackendSwitchingRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace backend switching rule default response
func (o *ReplaceBackendSwitchingRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBackendSwitchingRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
