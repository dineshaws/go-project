package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"gopkg.in/mgo.v2/bson"

	"github.com/dineshaws/go-project/common"
)

// Controller type declared
type Controller struct {
	DAO
}

// GetUser method to get user profile
func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet")
}

// AddUser method to register user in database
func (c *Controller) AddUser(w http.ResponseWriter, r *http.Request) {
	code := http.StatusOK
	message := "An unexpected error has occoured"
	var user User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		code = http.StatusBadRequest
		message = "Error while reading request"
		common.RespondWithError(w, code, message, err)
		return
	}
	if err := json.Unmarshal(body, &user); err != nil {
		code = http.StatusBadRequest
		message = "Error while decoding request"
		common.RespondWithError(w, code, message, err)
		return
	}
	// check username already exist or not.
	count, err := c.DAO.FindByUsername(user.Username)
	if err != nil {
		code = http.StatusBadRequest
		message = "Error while quering"
		common.RespondWithError(w, code, message, err)
		return
	}
	if count > 0 {
		code = http.StatusBadRequest
		message = "Username already exist"
		common.RespondWithError(w, code, message, errors.New("Existance error"))
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Username), 5)
	if err != nil {
		message = "Error while generating password request"
		common.RespondWithError(w, code, message, err)
	}
	user.Password = string(hash)
	user.ID = bson.NewObjectId()
	if err := c.DAO.Add(user); err != nil {
		common.RespondWithError(w, code, message, err)
		return
	}
	common.RespondWithJSON(w, http.StatusCreated, user)
}

// LoginUser method to Login user profile
func (c *Controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet")
}
