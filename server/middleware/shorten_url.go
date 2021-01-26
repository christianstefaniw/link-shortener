package middleware

import (
	"../models"
	"../storage"
	"context"
	"encoding/json"
	"github.com/dchest/uniuri"
	"net/http"
	"strings"
	"sync"
)



var mutex = &sync.Mutex{}


func ShortenURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var wg sync.WaitGroup

	strippedWhiteSpace := strings.ReplaceAll(r.FormValue("url"), " ", "")

	// Create slice of urls from form
	urls := strings.Split(strippedWhiteSpace, ",")

	shortenedURLs := models.URLS{}

	for _, url := range urls{
		if url != ""{
			// Add a worker to the wait group
			wg.Add(1)
			go shortenURL(url, &shortenedURLs, &wg)
		}
	}

	// Wait for workers in the wait group to complete
	wg.Wait()

	formattedURLs, err := json.MarshalIndent(shortenedURLs.Urls, "", "  ")
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(formattedURLs)
}

func shortenURL(fullURL string, urls *models.URLS, wg *sync.WaitGroup){
	// Run after outside function returns
	defer wg.Done()
	defer mutex.Unlock()

	shortURL := uniuri.NewLen(7)
	newURL := models.ShortenedURL{
		FullURL: fullURL,
		ShortURL: shortURL,
	}

	/*
	Prevents races by making sure only one goroutine can access the code block between mutex.Lock() and mutex.Unlock()
	 */
	mutex.Lock()
	if insert(&newURL) != nil{
		return
	}
	urls.Urls = append(urls.Urls, &newURL)
}

func insert(url *models.ShortenedURL) error{
	_, err := storage.Collection.InsertOne(context.Background(), url)
	return err
}