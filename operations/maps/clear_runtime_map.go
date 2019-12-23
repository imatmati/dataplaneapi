// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ClearRuntimeMapHandlerFunc turns a function with the right signature into a clear runtime map handler
type ClearRuntimeMapHandlerFunc func(ClearRuntimeMapParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ClearRuntimeMapHandlerFunc) Handle(params ClearRuntimeMapParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ClearRuntimeMapHandler interface for that can handle valid clear runtime map params
type ClearRuntimeMapHandler interface {
	Handle(ClearRuntimeMapParams, interface{}) middleware.Responder
}

// NewClearRuntimeMap creates a new http.Handler for the clear runtime map operation
func NewClearRuntimeMap(ctx *middleware.Context, handler ClearRuntimeMapHandler) *ClearRuntimeMap {
	return &ClearRuntimeMap{Context: ctx, Handler: handler}
}

/*ClearRuntimeMap swagger:route DELETE /services/haproxy/runtime/maps/{name} Maps clearRuntimeMap

Remove all map entries from the map file

Remove all map entries from the map file.

*/
type ClearRuntimeMap struct {
	Context *middleware.Context
	Handler ClearRuntimeMapHandler
}

func (o *ClearRuntimeMap) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewClearRuntimeMapParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
