package basketapi

import (
	"encoding/json"
	"log"
	"net/http"
)

func register(w http.ResponseWriter, r *http.Request) {
	setContext(r)
	if r.Method == http.MethodPost {
		if r.Body == nil {
			errorEmptyBody(w)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var user User
		err := decoder.Decode(&user)

		if err != nil {
			errorBadRequest(w)
			log.Printf("Error with decoding user JSON: %v", err)
			return
		}

		valid, err := user.Validate()
		if err != nil {
			errorBadRequest(w, err.Error())
			log.Printf("Error with validating user: %v", err)
			log.Print(user)
			return
		}

		userCheck, err := findUserByEmail(user.Email)
		if err != nil {
			errorBadRequest(w, err.Error())
			log.Printf("Error checking if user is unique: %v", err)
			log.Print(user)
			return
		}

		log.Println("User check")
		log.Println(userCheck)

		if valid && len(userCheck.Email) == 0 {
			key, err := user.Save()
			if err != nil {
				errorBadRequest(w, err.Error())
				log.Printf("Error saving user: %v", err)
				log.Println(user)
				return
			}

			log.Printf("New user saved: %v", key)
		}
	} else {
		errorNotAllowed(w)
	}

	return
}
