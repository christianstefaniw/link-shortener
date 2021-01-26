package server

import (
	"../router"
	"fmt"
	"log"
	"os"

	"net/http"
)

func New(){
	newRouter := router.Router()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	fmt.Printf("Listening on port: %s\n", port)
	//log.Fatal(http.ListenAndServe("192.168.1.131:"+port, newRouter))
	log.Fatal(http.ListenAndServe(":"+port, newRouter))
}
