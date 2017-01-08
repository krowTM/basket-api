package basketapi

import (
	"errors"
	"log"
	"net/mail"
	"time"

	"google.golang.org/appengine/datastore"
)

// User is the application user
type User struct {
	FirstName     string
	LastName      string
	Email         string
	EmailVerified bool
	Created       time.Time
}

// Validate checks if user data is valid
func (user User) Validate() (bool, error) {
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return false, errors.New("Invalid email address")
	} else if len(user.FirstName) == 0 {
		return false, errors.New("FirstName is required")
	} else if len(user.LastName) == 0 {
		return false, errors.New("LastName is required")
	}

	return true, nil
}

// Save saves a user to the DB
func (user User) Save() (datastore.Key, error) {
	user.Created = time.Now()
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
