package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/haproxytech/client-native"
	"github.com/haproxytech/dataplaneapi/haproxy"
	"github.com/haproxytech/dataplaneapi/misc"
	"github.com/haproxytech/dataplaneapi/operations/backend"
	"github.com/haproxytech/models"
)

//CreateBackendHandlerImpl implementation of the CreateBackendHandler interface using client-native client
type CreateBackendHandlerImpl struct {
	Client      *client_native.HAProxyClient
	ReloadAgent *haproxy.ReloadAgent
}

//DeleteBackendHandlerImpl implementation of the DeleteBackendHandler interface using client-native client
type DeleteBackendHandlerImpl struct {
	Client      *client_native.HAProxyClient
	ReloadAgent *haproxy.ReloadAgent
}

//GetBackendHandlerImpl implementation of the GetBackendHandler interface using client-native client
type GetBackendHandlerImpl struct {
	Client *client_native.HAProxyClient
}

//GetBackendsHandlerImpl implementation of the GetBackendsHandler interface using client-native client
type GetBackendsHandlerImpl struct {
	Client *client_native.HAProxyClient
}

//ReplaceBackendHandlerImpl implementation of the ReplaceBackendHandler interface using client-native client
type ReplaceBackendHandlerImpl struct {
	Client      *client_native.HAProxyClient
	ReloadAgent *haproxy.ReloadAgent
}

//Handle executing the request and returning a response
func (h *CreateBackendHandlerImpl) Handle(params backend.CreateBackendParams, principal interface{}) middleware.Responder {
	t := ""
	v := int64(0)
	if params.TransactionID != nil {
		t = *params.TransactionID
	}
	if params.Version != nil {
		v = *params.Version
	}

	if t != "" && *params.ForceReload {
		msg := "Both force_reload and transaction specified, specify only one"
		c := misc.ErrHTTPBadRequest
		e := &models.Error{
			Message: &msg,
			Code:    &c,
		}
		return backend.NewCreateBackendDefault(int(*e.Code)).WithPayload(e)
	}

	err := h.Client.Configuration.CreateBackend(params.Data, t, v)
	if err != nil {
		e := misc.HandleError(err)
		return backend.NewCreateBackendDefault(int(*e.Code)).WithPayload(e)
	}

	if params.TransactionID == nil {
		if *params.ForceReload {
			err := h.ReloadAgent.ForceReload()
			if err != nil {
				e := misc.HandleError(err)
				return backend.NewCreateBackendDefault(int(*e.Code)).WithPayload(e)
			}
			return backend.NewCreateBackendCreated().WithPayload(params.Data)
		}
		rID := h.ReloadAgent.Reload()
		return backend.NewCreateBackendAccepted().WithReloadID(rID).WithPayload(params.Data)
	}
	return backend.NewCreateBackendAccepted().WithPayload(params.Data)
}

//Handle executing the request and returning a response
func (h *DeleteBackendHandlerImpl) Handle(params backend.DeleteBackendParams, principal interface{}) middleware.Responder {
	t := ""
	v := int64(0)
	if params.TransactionID != nil {
		t = *params.TransactionID
	}
	if params.Version != nil {
		v = *params.Version
	}

	if t != "" && *params.ForceReload {
		msg := "Both force_reload and transaction specified, specify only one"
		c := misc.ErrHTTPBadRequest
		e := &models.Error{
			Message: &msg,
			Code:    &c,
		}
		return backend.NewDeleteBackendDefault(int(*e.Code)).WithPayload(e)
	}

	err := h.Client.Configuration.DeleteBackend(params.Name, t, v)
	if err != nil {
		e := misc.HandleError(err)
		return backend.NewDeleteBackendDefault(int(*e.Code)).WithPayload(e)
	}
	if params.TransactionID == nil {
		if *params.ForceReload {
			err := h.ReloadAgent.ForceReload()
			if err != nil {
				e := misc.HandleError(err)
				return backend.NewDeleteBackendDefault(int(*e.Code)).WithPayload(e)
			}
			return backend.NewDeleteBackendNoContent()
		}
		rID := h.ReloadAgent.Reload()
		return backend.NewDeleteBackendAccepted().WithReloadID(rID)
	}
	return backend.NewDeleteBackendAccepted()
}

//Handle executing the request and returning a response
func (h *GetBackendHandlerImpl) Handle(params backend.GetBackendParams, principal interface{}) middleware.Responder {
	t := ""
	if params.TransactionID != nil {
		t = *params.TransactionID
	}

	bck, err := h.Client.Configuration.GetBackend(params.Name, t)
	if err != nil {
		e := misc.HandleError(err)
		return backend.NewGetBackendDefault(int(*e.Code)).WithPayload(e)
	}
	return backend.NewGetBackendOK().WithPayload(bck)
}

//Handle executing the request and returning a response
func (h *GetBackendsHandlerImpl) Handle(params backend.GetBackendsParams, principal interface{}) middleware.Responder {
	t := ""
	if params.TransactionID != nil {
		t = *params.TransactionID
	}

	bcks, err := h.Client.Configuration.GetBackends(t)
	if err != nil {
		e := misc.HandleError(err)
		return backend.NewGetBackendsDefault(int(*e.Code)).WithPayload(e)
	}
	return backend.NewGetBackendsOK().WithPayload(bcks)
}

//Handle executing the request and returning a response
func (h *ReplaceBackendHandlerImpl) Handle(params backend.ReplaceBackendParams, principal interface{}) middleware.Responder {
	t := ""
	v := int64(0)
	if params.TransactionID != nil {
		t = *params.TransactionID
	}
	if params.Version != nil {
		v = *params.Version
	}

	if t != "" && *params.ForceReload {
		msg := "Both force_reload and transaction specified, specify only one"
		c := misc.ErrHTTPBadRequest
		e := &models.Error{
			Message: &msg,
			Code:    &c,
		}
		return backend.NewReplaceBackendDefault(int(*e.Code)).WithPayload(e)
	}

	err := h.Client.Configuration.EditBackend(params.Name, params.Data, t, v)
	if err != nil {
		e := misc.HandleError(err)
		return backend.NewReplaceBackendDefault(int(*e.Code)).WithPayload(e)
	}

	if params.TransactionID == nil {
		if *params.ForceReload {
			err := h.ReloadAgent.ForceReload()
			if err != nil {
				e := misc.HandleError(err)
				return backend.NewReplaceBackendDefault(int(*e.Code)).WithPayload(e)
			}
			return backend.NewReplaceBackendOK().WithPayload(params.Data)
		}
		rID := h.ReloadAgent.Reload()
		return backend.NewReplaceBackendAccepted().WithReloadID(rID).WithPayload(params.Data)
	}
	return backend.NewReplaceBackendAccepted().WithPayload(params.Data)
}
