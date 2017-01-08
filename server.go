package basketapi

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
)

var ctx context.Context

func init() {
	http.HandleFunc("/", home)
	http.HandleFunc("/register", register)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func setContext(r *http.Request) {
	ctx = appengine.NewContext(r)
}

func getContext() context.Context {
	return ctx
}
