package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dineshaws/go-project/router"
)

func main() {
	fmt.Println("GO Project setup with mongo db")

	r := router.NewRouter()
	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatal(err)
	}

}
