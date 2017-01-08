package basketapi

import (
	"net/http"
	"strings"
)

func errorResponse(w http.ResponseWriter, status int, args ...string) {
	if len(args) == 0 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	errorMsgs := "Bad Request"
	var concatenator string
	for i, errorMsg := range args {
		if i == 0 {
			concatenator = ": "
		} else {
			concatenator = ", "
		}
		errorMsgs = strings.Join([]string{errorMsgs, errorMsg}, concatenator)
	}
	http.Error(w, errorMsgs, status)

	return
}
