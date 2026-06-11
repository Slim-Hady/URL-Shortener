package router

import(
	"fmt"
	"URL-SHORTENER/handlers"
	"net/http"
)

func Start(){
	http.HandleFunc("/shorten", handlers.ShortenURL)
	http.HandleFunc("/", handlers.RedirectURL)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
