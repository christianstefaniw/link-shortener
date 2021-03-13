package router

import (
	"../handlers"
	"github.com/ChristianStefaniw/cgr"
	"net/http"
)

func Router() *cgr.Router{
	router := cgr.NewRouter()
	initRoutes(router)
	return router
}

func initRoutes(router *cgr.Router){
	router.Route("/:url").Handler(handlers.RetrieveURL).Method("GET").Insert()
	router.Route("/").Handler(handlers.ShortenURL).Method("POST").Insert()
	router.Route("/routes").Method("GET").Handler(
		func(writer http.ResponseWriter, request *http.Request) {
			for _, route := range router.ViewRouteTree() {
				writer.Write([]byte(route))
			}
		},
	).Insert()
}
