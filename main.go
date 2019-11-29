package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dineshaws/go-project/router"
)

// init function
func init() {
	fmt.Println("Main package init")
}

func main() {
	fmt.Println("GO Project setup with mongo db")

	r := router.InitRouter()
	log.Printf("Applicatioin is running on port %v ", 4000)
	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatal(err)
	}

}
