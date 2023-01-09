// Package experimental provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package experimental

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	. "github.com/algorand/go-algorand/daemon/algod/api/server/v2/generated/model"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Simulates a raw transaction or transaction group as it would be evaluated on the network. WARNING: This endpoint is experimental and under active development. There are no guarantees in terms of functionality or future support.
	// (POST /v2/transactions/simulate)
	SimulateTransaction(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// SimulateTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) SimulateTransaction(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.SimulateTransaction(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface, m ...echo.MiddlewareFunc) {
	RegisterHandlersWithBaseURL(router, si, "", m...)
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/v2/transactions/simulate", wrapper.SimulateTransaction, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}