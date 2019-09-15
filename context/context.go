package context

import (
	"net/http"
)

// Context struct
type Context struct {
	AppId   string
	AppKey  string
	Sign    string
	Writer  http.ResponseWriter
	Request *http.Request
}
