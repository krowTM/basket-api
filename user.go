package basketapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"net/mail"
	"time"

	"golang.org/x/crypto/scrypt"

	"google.golang.org/appengine/datastore"
)

// PasswordSalt is used for random passwords
const PasswordSalt = "sADfROILbc"

// User is the application user
type User struct {
	FirstName     string
	LastName      string
	Email         string
	EmailVerified bool
	Password      []byte
	APIKey        string
	Created       time.Time
}

// ValidateRegister checks if user data is valid
func (user User) ValidateRegister() (bool, error) {
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return false, errors.New("Invalid email address")
	} else if len(user.FirstName) == 0 {
		return false, errors.New("FirstName is required")
	} else if len(user.LastName) == 0 {
		return false, errors.New("LastName is required")
	} else if len(user.Password) == 0 {
		return false, errors.New("Password is required")
	} else if len(user.Password) < 8 {
		return false, errors.New("Password is too short. Min 8")
	}

	return true, nil
}

// ValidateLogin checks if user login data is valid
func (user User) ValidateLogin() (bool, error) {
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return false, errors.New("Invalid email address")
	} else if len(user.Password) == 0 {
		return false, errors.New("Password is required")
	}

	return true, nil
}

// Save saves a user to the DB
func (user User) Save() (datastore.Key, error) {
	user.Created = time.Now()
	user.APIKey = RandStringBytesRmndr(64)
	password, err := scrypt.Key([]byte(user.Password), []byte(PasswordSalt), 16384, 8, 1, 32)
	user.Password = password
	if err != nil {
		return datastore.Key{}, err
	}
	key, err := datastore.Put(getContext(), datastore.NewIncompleteKey(getContext(), "User", nil), &user)
	if err != nil {
		return datastore.Key{}, err
	}

	log.Print(user)

	return *key, nil
}

// findUserByEmail searches for a user by email address
func findUserByEmail(email string) (User, error) {
	q := datastore.NewQuery("User").Filter("Email =", email).Limit(1)

	var users []User
	_, err := q.GetAll(getContext(), &users)
	if err != nil {
		return User{}, err
	}

	if len(users) == 0 {
		return User{}, nil
	}

	return users[0], nil
}

func register(w http.ResponseWriter, r *http.Request) {
	setContext(r)
	if r.Method == http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed, "Not Allowed")
		return
	}
	if r.Body == nil {
		errorResponse(w, http.StatusBadRequest, "Empty body")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)

	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Bad Request")
		log.Printf("Error with decoding user JSON: %v", err)
		return
	}

	valid, err := user.ValidateRegister()
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Bad Request", err.Error())
		log.Printf("Error with validating user: %v", err)
		log.Print(user)
		return
	}

	userCheck, err := findUserByEmail(user.Email)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Bad Request", err.Error())
		log.Printf("Error checking if user is unique: %v", err)
		log.Print(user)
		return
	}

	log.Println("User check")
	log.Println(userCheck)

	if valid && len(userCheck.Email) == 0 {
		key, err := user.Save()
		if err != nil {
			errorResponse(w, http.StatusBadRequest, "Bad Request", err.Error())
			log.Printf("Error saving user: %v", err)
			log.Println(user)
			return
		}

		log.Printf("New user saved: %v", key)
	}

	return
}

func login(w http.ResponseWriter, r *http.Request) {
	setContext(r)
	if r.Method == http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed, "Not Allowed")
		return
	}
	if r.Body == nil {
		errorResponse(w, http.StatusBadRequest, "Empty body")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)

	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Bad Request")
		log.Printf("Error with decoding user login JSON: %v", err)
		return
	}

	valid, err := user.ValidateLogin()
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Bad Request", err.Error())
		log.Printf("Error with validating user login: %v", err)
		log.Print(user)
		return
	}

	userCheck, err := findUserByEmail(user.Email)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Bad Request", err.Error())
		log.Printf("Error checking if user exists: %v", err)
		log.Print(user)
		return
	}

	if !valid || len(userCheck.Email) == 0 {
		errorResponse(w, http.StatusMethodNotAllowed, "Not Allowed")
	}

	log.Println("User check")
	log.Println(userCheck)

	password, err := scrypt.Key([]byte(user.Password), []byte(PasswordSalt), 16384, 8, 1, 32)
	if valid && bytes.Compare(password, userCheck.Password) == 0 {
		// login user

		log.Printf("User logged in: %v", userCheck)
	}

	return
}

// letterBytes is a list of characters used for password crypt
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandStringBytesRmndr generates a random string of n characters
func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
