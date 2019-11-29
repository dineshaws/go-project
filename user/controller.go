package user

import (
	"fmt"
	"net/http"
)

type Controller struct {
	DAO
}

// GetUser method to get user profile
func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet")
}

// AddUser method to get user profile
func (c *Controller) AddUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet")
}

// LoginUser method to get user profile
func (c *Controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet")
}
