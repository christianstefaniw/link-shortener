package router

import (
	"../middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router{
	router := mux.NewRouter()
	initRoutes(router)
	return router
}

func initRoutes(router *mux.Router){

	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("../client/build/static")))
	router.PathPrefix("/static/").Handler(staticHandler)

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/{url}", middleware.RetrieveURL).Methods("GET")
	router.HandleFunc("/", middleware.ShortenURL).Methods("POST")
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/build/index.html")
}