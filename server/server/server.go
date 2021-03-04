package server

import (
	"../router"
	"github.com/ChristianStefaniw/cgr"
	"os"
)

func New(){
	newRouter := router.Router()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	cgr.Run(port, newRouter)
}
