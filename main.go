package main

import(
	"fmt"
	"URL-SHORTENER/storage"
)	

func main(){
	store := storage.NewRedisStore()

	originalURL := "https://google.com"

	shortURL, err := store.SaveURL(originalURL)

	if err != nil {
		fmt.Println("Error occuer", err)
		return 
	}
	fmt.Println("ShortURL saved " , shortURL)
	originURL, err := store.GetOriginURL(shortURL)
	if err != nil { 
		fmt.Println("Error occuer", err)
		return
	}
	fmt.Println("Oringinal URL ",originURL)
	
}