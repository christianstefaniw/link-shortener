package router

import (
	"../middleware"
	"github.com/ChristianStefaniw/cgr"
	"net/http"
)

func Router() *cgr.Router{
	router := cgr.NewRouter()
	initRoutes(router)
	return router
}

func initRoutes(router *cgr.Router){
	router.Route("/lnk/:url").Handler(middleware.RetrieveURL).Method("GET").Insert(router)
	router.Route("/").Handler(middleware.ShortenURL).Method("POST").Insert(router)
	router.Route("/routes").Method("GET").Handler(
		func(writer http.ResponseWriter, request *http.Request) {
			for _, route := range router.ViewRouteTree() {
				writer.Write([]byte(route))
			}
		},
	).Insert(router)
}
