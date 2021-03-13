package server

import (
	"../router"
	"github.com/ChristianStefaniw/cgr"
	"os"
)

func New(){
	r := router.Router()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	cgr.Run(port, r)
}
