package basketapi

import (
	"net/http"
	"strings"
)

func errorNotAllowed(w http.ResponseWriter) {
	http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
	return
}

func errorEmptyBody(w http.ResponseWriter) {
	http.Error(w, "Empty body", http.StatusBadRequest)
	return
}

func errorBadRequest(w http.ResponseWriter, args ...string) {
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
	http.Error(w, errorMsgs, http.StatusBadRequest)

	return
}
