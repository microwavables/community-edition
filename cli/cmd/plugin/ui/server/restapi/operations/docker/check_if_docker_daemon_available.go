// Code generated by go-swagger; DO NOT EDIT.

package docker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CheckIfDockerDaemonAvailableHandlerFunc turns a function with the right signature into a check if docker daemon available handler
type CheckIfDockerDaemonAvailableHandlerFunc func(CheckIfDockerDaemonAvailableParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CheckIfDockerDaemonAvailableHandlerFunc) Handle(params CheckIfDockerDaemonAvailableParams) middleware.Responder {
	return fn(params)
}

// CheckIfDockerDaemonAvailableHandler interface for that can handle valid check if docker daemon available params
type CheckIfDockerDaemonAvailableHandler interface {
	Handle(CheckIfDockerDaemonAvailableParams) middleware.Responder
}

// NewCheckIfDockerDaemonAvailable creates a new http.Handler for the check if docker daemon available operation
func NewCheckIfDockerDaemonAvailable(ctx *middleware.Context, handler CheckIfDockerDaemonAvailableHandler) *CheckIfDockerDaemonAvailable {
	return &CheckIfDockerDaemonAvailable{Context: ctx, Handler: handler}
}

/* CheckIfDockerDaemonAvailable swagger:route GET /api/providers/docker/daemon docker checkIfDockerDaemonAvailable

Check if docker deamon is available

*/
type CheckIfDockerDaemonAvailable struct {
	Context *middleware.Context
	Handler CheckIfDockerDaemonAvailableHandler
}

func (o *CheckIfDockerDaemonAvailable) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCheckIfDockerDaemonAvailableParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
