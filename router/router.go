package router

import(
	"fmt"
	"URL-SHORTENER/handlers"
	"net/http"
)

func Start(){
	http.HandleFunc("/shorten", handlers.ShortenURL)

	http.HandleFunc("/app/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/index.html")
	})

	http.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/app/", http.StatusTemporaryRedirect)
	})

	http.HandleFunc("/", handlers.RedirectURL)

	fmt.Println("Server running on http://localhost:8080/app/")
	http.ListenAndServe(":8080", nil)
}
