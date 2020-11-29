package webserver

import (
	"net/http"
)

type middleware func(next http.HandlerFunc) http.HandlerFunc

func dummyMW(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// add middleware methods here

		next.ServeHTTP(w, r)
	}
}

// nestedMiddleware provides syntactic sugar to create a new middleware
// which will be the result of chaining the ones received as parameters.
// https://www.codemio.com/2019/01/advanced-golang-tutorial-http-middleware.html
func nestedMiddleware(mw ...middleware) middleware {

	return func(final http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			last := final
			for i := len(mw) - 1; i >= 0; i-- {
				last = mw[i](last)
			}
			last(w, r)
		}
	}
}
