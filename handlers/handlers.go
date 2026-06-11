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

func ShortenURL (w http.ResponseWriter, r *http.Request) {

	var req models.ShortenerRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w,"Invalid JSON ", http.StatusBadRequest)
		return
	}
	hashedURL, err := store.SaveURL(req.URL)
	if err != nil {
		fmt.Println("Error had occuer ", err)
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}

	res := models.ShortenerResponse{
		ShortURL: "http://RealSlimHady" + hashedURL,
	}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)

}

func RedirectURL(w http.ResponseWriter, r *http.Request){
	shortURLHash := strings.TrimPrefix(r.URL.Path, "/")

	if shortURLHash == ""{
		http.Error(w , "Invaild shortURL" , http.StatusBadRequest)
		return
	}
	originalURL, err := store.GetOriginURL(shortURLHash)

	if err != nil {
        http.Error(w, "URL not found", http.StatusNotFound)
        return
    }

	http.Redirect(w,r,originalURL,http.StatusFound)
}