package handlers

import (
	"URL-SHORTENER/models"
	"URL-SHORTENER/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var store = storage.NewRedisStore()

/*
this function creating the short URL by taking the Long URL from user request
by applying SaveURL function hash the URL then save to response
*/
func ShortenURL(w http.ResponseWriter, r *http.Request) {

	var req models.ShortenerRequest
	err := json.NewDecoder(r.Body).Decode(&req) // this take request.body to the struct

	if err != nil {
		http.Error(w, "Invalid JSON ", http.StatusBadRequest)
		return
	}

	hashedURL, err := store.SaveURL(req.URL)

	if err != nil {
		fmt.Println("Error had occuer ", err)
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}
	// store the response on the hashedURL
	res := models.ShortenerResponse{
		ShortURL: "http://localhost:8080/" + hashedURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

/*
this function job is redirect to the website when someone use the shorturl we created on the ShortURL function
*/
func RedirectURL(w http.ResponseWriter, r *http.Request) {

	shortURLHash := strings.TrimPrefix(r.URL.Path, "/")

	if shortURLHash == "" {
		http.Error(w, "Invaild shortURL", http.StatusBadRequest)
		return
	}
	// on GetOriginURL function on redis-store.go i made the mapping on the redis bet the shortURL --> longURL and longURL --> shortURL
	originalURL, err := store.GetOriginURL(shortURLHash)

	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
