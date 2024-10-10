package handler

import "net/http"

// Route is an http.Handler that knows the mux pattern
// under which it will be registered.
type Route interface {
	http.Handler
	Pattern() string
}
