package middleware

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc
type GeneralHandle func(http.Handler) http.Handler

// Chain applies middlewares to a http.HandlerFunc
// @handler
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// GetVars - Return url vars
// @example /api/{key}/send
// @vars = {"key":data}
func GetVars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

// GetHeader - Return Header value stored on passed key
func GetHeader(r *http.Request, key string) string {
	return r.Header.Get(key)
}

// InjectHeader - Inject data on header request
func InjectHeader(r *http.Request, key, val string) {
	r.Header.Add(key, val)
}

// GetQueryes - Return queryes values
// @example /api?key=data
func GetQueryes(r *http.Request) url.Values {
	return r.URL.Query()
}

// GetBody - Return byte body data
func GetBody(r *http.Request) ([]byte, error) {
	return ioutil.ReadAll(r.Body)
}

// GetForm - Return parsed form data
func GetForm(r *http.Request) (form url.Values, err error) {
	err = r.ParseForm()
	form = r.Form
	return
}

// DecodeForm - Decoded parsed form data on interface
func DecodeForm(dst interface{}, src map[string][]string) error {
	decoder := schema.NewDecoder()
	return decoder.Decode(dst, src)
}
