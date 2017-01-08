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
	http.HandleFunc("/user/register", register)
	http.HandleFunc("/user/login", login)
	http.HandleFunc("/cart/add", addToCart)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}

func setContext(r *http.Request) {
	ctx = appengine.NewContext(r)
}

func getContext() context.Context {
	return ctx
}
