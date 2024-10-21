package handler

import "net/http"

// Router is an http.Handler that knows the mux pattern
// under which it will be registered.
type Router interface {
	http.Handler
	Pattern() string
}
