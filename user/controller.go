package user

import (
	"encoding/json"
	"errors"
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
	tokenString := r.Header.Get("access-token")
	username, err := common.VerifyToken(tokenString)
	if err != nil {
		code := http.StatusUnauthorized
		message := "Invalid token"
		common.RespondWithError(w, code, message, err)
		return
	}
	user, err := c.DAO.FindByUsername(username)
	if err != nil {
		code := http.StatusUnauthorized
		message := "Token is expired"
		common.RespondWithError(w, code, message, err)
		return
	}
	user.Password = ""
	common.RespondWithJSON(w, http.StatusOK, user)

}

// AddUser method to register user in database
func (c *Controller) AddUser(w http.ResponseWriter, r *http.Request) {
	code := http.StatusOK
	message := "An unexpected error has occoured"
	defer r.Body.Close()
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
	count, err := c.DAO.FindUsernameCount(user.Username)
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

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
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
	user.Password = ""
	common.RespondWithJSON(w, http.StatusCreated, user)
}

// LoginUser method to Login user profile
func (c *Controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	code := http.StatusInternalServerError
	message := "An unexpected error has occoured"

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		code = http.StatusBadRequest
		message = "Error while reading request"
		common.RespondWithError(w, code, message, err)
		return
	}
	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		code = http.StatusBadRequest
		message = "Error while decoding request"
		common.RespondWithError(w, code, message, err)
		return
	}

	dbUser, err := c.DAO.FindByUsername(user.Username)
	if err != nil {
		code = http.StatusBadRequest
		message = "Invalid username"
		common.RespondWithError(w, code, message, err)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		code = http.StatusBadRequest
		message = "Invalid password"
		common.RespondWithError(w, code, message, err)
		return
	}

	token, err := common.GenerateToken(dbUser.Username)
	if err != nil {
		message = "Error while generating token"
		common.RespondWithError(w, code, message, err)
		return
	}

	dbUser.Token = token
	dbUser.Password = ""
	common.RespondWithJSON(w, http.StatusCreated, dbUser)
}
