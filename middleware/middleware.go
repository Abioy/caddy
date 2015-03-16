// Package middleware provides some types and functions common among middleware.
package middleware

import "net/http"

type (
	// Generator represents the outer layer of a middleware that
	// parses tokens to configure the middleware instance.
	Generator func(Controller) (Middleware, error)

	// Middleware is the middle layer which represents the traditional
	// idea of middleware: it is passed the next HandlerFunc in the chain
	// and returns the inner layer, which is the actual HandlerFunc.
	Middleware func(http.HandlerFunc) http.HandlerFunc

	// Controller is the type which middleware generators use to access
	// tokens and the server and any other information they need to
	// configure themselves.
	Controller interface {
		Next() bool
		NextArg() bool
		NextLine() bool
		NextBlock() bool
		Val() string
		Args(...*string) bool
		RemainingArgs() []string
		ArgErr() error
		Err(string) error
		Startup(func() error)
		Root() string
		Host() string
		Port() string
		Context() Path
	}
)