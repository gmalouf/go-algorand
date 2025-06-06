// Package data provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package data

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	. "github.com/algorand/go-algorand/daemon/algod/api/server/v2/generated/model"
	"github.com/algorand/go-algorand/data/basics"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Removes minimum sync round restriction from the ledger.
	// (DELETE /v2/ledger/sync)
	UnsetSyncRound(ctx echo.Context) error
	// Returns the minimum sync round the ledger is keeping in cache.
	// (GET /v2/ledger/sync)
	GetSyncRound(ctx echo.Context) error
	// Given a round, tells the ledger to keep that round in its cache.
	// (POST /v2/ledger/sync/{round})
	SetSyncRound(ctx echo.Context, round basics.Round) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// UnsetSyncRound converts echo context to params.
func (w *ServerInterfaceWrapper) UnsetSyncRound(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UnsetSyncRound(ctx)
	return err
}

// GetSyncRound converts echo context to params.
func (w *ServerInterfaceWrapper) GetSyncRound(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetSyncRound(ctx)
	return err
}

// SetSyncRound converts echo context to params.
func (w *ServerInterfaceWrapper) SetSyncRound(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "round" -------------
	var round basics.Round

	err = runtime.BindStyledParameterWithOptions("simple", "round", ctx.Param("round"), &round, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SetSyncRound(ctx, round)
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

	router.DELETE(baseURL+"/v2/ledger/sync", wrapper.UnsetSyncRound, m...)
	router.GET(baseURL+"/v2/ledger/sync", wrapper.GetSyncRound, m...)
	router.POST(baseURL+"/v2/ledger/sync/:round", wrapper.SetSyncRound, m...)

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
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
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
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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
