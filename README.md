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

