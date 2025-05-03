package middleware

import (
	"bookhub/internal/types"
	"log"
	"net/http"
	"time"
)

// Logging logs all requests with its path and the time it took to process
func Logging() types.Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json") // set to return JSON data

			// Do middleware things
			start := time.Now()                                           // Get date
			defer func() { log.Println(r.URL.Path, time.Since(start)) }() // print path and date of access

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Function wraps an HTTP handler with a sequence of middleware, so each middleware can process the
// request and response before reaching the main handler.
func Chain(f http.HandlerFunc, middlewares ...types.Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
