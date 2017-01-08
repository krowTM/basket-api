package basketapi

import (
	"encoding/json"
	"log"
	"net/http"
)

// Cart is a shopping cart struct
type Cart struct {
	ProductID   string
	Quantity    int
	ProductData string
}

func addToCart(w http.ResponseWriter, r *http.Request) {
	setContext(r)
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed, "Not Allowed")
		return
	}

	// check header for APIKey

	if r.Body == nil {
		errorResponse(w, http.StatusBadRequest, "Empty body")
		return
	}

	// validate data
	decoder := json.NewDecoder(r.Body)
	var cartPost map[string]interface{}
	err := decoder.Decode(&cartPost)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Bad Request", err.Error())
		log.Printf("Error with validating cart rawpost: %v", err)
		log.Print(cartPost)
		return
	}

	// save to db

	log.Print(cartPost["Email"])
	log.Print(cartPost)

	return
}
