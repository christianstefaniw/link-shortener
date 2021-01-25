package middleware

import (
	"../models"
	"../storage"
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)


func RetrieveURL(w http.ResponseWriter, r *http.Request) {
	var result models.ShortenedURL
	shortUrl := mux.Vars(r)["url"]

	filter := bson.M{"shorturl": shortUrl}
	if err := storage.Collection.FindOne(context.Background(), filter).Decode(&result); err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(shortUrl + " not found"))
		return
	}

	http.Redirect(w, r, result.FullURL, http.StatusMovedPermanently)
}
