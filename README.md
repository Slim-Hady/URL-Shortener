# What is URL-Shortener
URL Shortener allows users to input a long URL and receive a shortened version of it.

# Set-up 

clone the repo 

install go dep 

```bash
go mod tidy
```

Start Redis:

```bash
docker run -d -p 6379:6379 --name redis redis
```
Start Redis:
```bash
docker run -d -p 6379:6379 --name redis redis
```
run the server 

```bash
go run main.go
```

```bash
Server starts on http://localhost:8080
```

### API endpoints 

1- POST /shorten : Create the shortURL 

on the body json write 

```json
{
    "url": "" // here write the longURL bet ""
}
```

Response 

```json
{
  "short_url": "http://localhost:8080/{SomeNumbers}"
}
```

2- GET /{hash} : redirect to original URL


### Project Structure
```
url-shortener/
├── handlers/
│   └── handlers.go      # HTTP handlers (ShortenURL, RedirectURL)
├── models/
│   └── url.go           # Request/Response structs
├── router/
│   └── router.go        # Route registration + server start
├── storage/
│   └── redis-store.go   # Redis connection, SaveURL, GetOriginalURL
├── go.mod
└── main.go
```

## AI Usage :

Frontend is completely vibe coded using opencode zen free models MiMo v2.5 high 

it also modified the router/router.go to connect the front-end 

this part is vibecoded
```GO

	http.HandleFunc("/app/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/index.html")
	})

	http.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/app/", http.StatusTemporaryRedirect)
	})

```

# Photos

## This is not a longURL i know but it work on all URL btw :) 

![alt text](Images/image.png)

## The result 

![alt text](Images/result.png)