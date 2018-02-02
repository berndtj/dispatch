///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/vmware/dispatch/pkg/application-manager/gen/restapi/operations/application"
)

// NewApplicationManagerAPI creates a new ApplicationManager instance
func NewApplicationManagerAPI(spec *loads.Document) *ApplicationManagerAPI {
	return &ApplicationManagerAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		ApplicationAddAppHandler: application.AddAppHandlerFunc(func(params application.AddAppParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ApplicationAddApp has not yet been implemented")
		}),
		ApplicationDeleteAppHandler: application.DeleteAppHandlerFunc(func(params application.DeleteAppParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ApplicationDeleteApp has not yet been implemented")
		}),
		ApplicationGetAppHandler: application.GetAppHandlerFunc(func(params application.GetAppParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ApplicationGetApp has not yet been implemented")
		}),
		ApplicationGetAppsHandler: application.GetAppsHandlerFunc(func(params application.GetAppsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ApplicationGetApps has not yet been implemented")
		}),
		ApplicationUpdateAppHandler: application.UpdateAppHandlerFunc(func(params application.UpdateAppParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ApplicationUpdateApp has not yet been implemented")
		}),

		// Applies when the "Cookie" header is set
		CookieAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (cookie) Cookie from header param [Cookie] has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*ApplicationManagerAPI VMware Dispatch - Application Manager APIs
 */
type ApplicationManagerAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// CookieAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Cookie provided in the header
	CookieAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// ApplicationAddAppHandler sets the operation handler for the add app operation
	ApplicationAddAppHandler application.AddAppHandler
	// ApplicationDeleteAppHandler sets the operation handler for the delete app operation
	ApplicationDeleteAppHandler application.DeleteAppHandler
	// ApplicationGetAppHandler sets the operation handler for the get app operation
	ApplicationGetAppHandler application.GetAppHandler
	// ApplicationGetAppsHandler sets the operation handler for the get apps operation
	ApplicationGetAppsHandler application.GetAppsHandler
	// ApplicationUpdateAppHandler sets the operation handler for the update app operation
	ApplicationUpdateAppHandler application.UpdateAppHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *ApplicationManagerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ApplicationManagerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ApplicationManagerAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ApplicationManagerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ApplicationManagerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ApplicationManagerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ApplicationManagerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ApplicationManagerAPI
func (o *ApplicationManagerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.CookieAuth == nil {
		unregistered = append(unregistered, "CookieAuth")
	}

	if o.ApplicationAddAppHandler == nil {
		unregistered = append(unregistered, "application.AddAppHandler")
	}

	if o.ApplicationDeleteAppHandler == nil {
		unregistered = append(unregistered, "application.DeleteAppHandler")
	}

	if o.ApplicationGetAppHandler == nil {
		unregistered = append(unregistered, "application.GetAppHandler")
	}

	if o.ApplicationGetAppsHandler == nil {
		unregistered = append(unregistered, "application.GetAppsHandler")
	}

	if o.ApplicationUpdateAppHandler == nil {
		unregistered = append(unregistered, "application.UpdateAppHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ApplicationManagerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ApplicationManagerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "cookie":

			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.CookieAuth)

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *ApplicationManagerAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types
func (o *ApplicationManagerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *ApplicationManagerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ApplicationManagerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the application manager API
func (o *ApplicationManagerAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ApplicationManagerAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"][""] = application.NewAddApp(o.context, o.ApplicationAddAppHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/{application}"] = application.NewDeleteApp(o.context, o.ApplicationDeleteAppHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/{application}"] = application.NewGetApp(o.context, o.ApplicationGetAppHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"][""] = application.NewGetApps(o.context, o.ApplicationGetAppsHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/{application}"] = application.NewUpdateApp(o.context, o.ApplicationUpdateAppHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ApplicationManagerAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *ApplicationManagerAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
