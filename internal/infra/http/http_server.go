package httpserver

import (
	"mime/multipart"
	"net/http"
)

type HTTPServer interface {
	Run(addr ...string) error
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type HTTPRoutes interface {
	GET(string, ...HTTPHandlerFunc)
	POST(string, ...HTTPHandlerFunc)
	DELETE(string, ...HTTPHandlerFunc)
	PATCH(string, ...HTTPHandlerFunc)
	PUT(string, ...HTTPHandlerFunc)
	SetBasePath(basePath string)
	SetSwagger(string)
}

type HTTPContext interface {
	Header(key, value string)
	GetHeader(key string) string
	JSON(code int, obj any)
	BindJSON(obj any) error
	Param(key string) string
	DefaultQuery(key, defaultValue string) string
	MultipartForm() (*multipart.Form, error)
}

type HTTPHandlerFunc func(HTTPContext)

type Payload map[string]any
