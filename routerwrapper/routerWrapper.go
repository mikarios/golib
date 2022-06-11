// Package routerwrapper is used to create gorilla/mux APIs easier if there are multiple optional query parameters.
// Instead of registering the same api 2 or more times you can create them like this
/*
	routerwrapper.New(myRouter, nil).
		HandleFunc("/api/v1/something", myHandlerFunc).
		Methods(http.MethodGet).
		Query("page", `\d+`, true).
		Query("size", `\d+`, true).
		Query("search", ``, true).
		Create()
*/
package routerwrapper

import (
	"net/http"

	"github.com/gorilla/mux"
	combinations "github.com/mxschmitt/golang-combinations"
)

type logger interface {
	Printf(format string, v ...interface{})
}

type routerWrapper struct {
	router     *mux.Router
	path       string
	handleFunc func(http.ResponseWriter, *http.Request)
	methods    []string
	queries    map[string]stringBool
	logger     logger
}

type stringBool struct {
	s string
	b bool
}

// New creates a new routerWrapper logger is optional if you wish to log the endpoints created.
func New(router *mux.Router, log logger) *routerWrapper {
	return &routerWrapper{router: router, queries: make(map[string]stringBool), logger: log}
}

// HandleFunc registers a new route with a matcher for the URL path.
func (wrapper *routerWrapper) HandleFunc(
	path string,
	handleFunc func(w http.ResponseWriter, r *http.Request),
) *routerWrapper {
	wrapper.path = path
	wrapper.handleFunc = handleFunc

	return wrapper
}

// Methods registers a new route with a matcher for HTTP methods.
func (wrapper *routerWrapper) Methods(methods ...string) *routerWrapper {
	wrapper.methods = methods

	return wrapper
}

// Query adds a matcher for URL query values.
func (wrapper *routerWrapper) Query(name, pattern string, optional bool) *routerWrapper {
	fullPattern := "{" + name + "}"

	if pattern != "" {
		fullPattern = "{" + name + ":" + pattern + "}"
	}

	wrapper.queries[name] = stringBool{fullPattern, optional}

	return wrapper
}

// Create actually constructs the router based on the given values.
func (wrapper *routerWrapper) Create() {
	mandatory := make([]string, 0)
	optionalQueryKeys := make([]string, 0)

	for k, v := range wrapper.queries {
		if v.b {
			optionalQueryKeys = append(optionalQueryKeys, k)
		} else {
			mandatory = append(mandatory, k, v.s)
		}
	}

	for i := len(optionalQueryKeys); i > 0; i-- {
		for _, optional := range combinations.Combinations(optionalQueryKeys, i) {
			all := mandatory
			for _, v := range optional {
				all = append(all, v, wrapper.queries[v].s)
			}

			wrapper.router.HandleFunc(wrapper.path, wrapper.handleFunc).Methods(wrapper.methods...).Queries(all...)

			if wrapper.logger != nil {
				wrapper.logger.Printf(
					"Created endpoint %v with methods: %v and query parameters: %v",
					wrapper.path,
					wrapper.methods,
					all,
				)
			}
		}
	}

	wrapper.router.HandleFunc(wrapper.path, wrapper.handleFunc).Methods(wrapper.methods...).Queries(mandatory...)

	if wrapper.logger != nil {
		wrapper.logger.Printf(
			"Created endpoint %v with methods: %v and query parameters: %v",
			wrapper.path,
			wrapper.methods,
			mandatory,
		)
	}
}
