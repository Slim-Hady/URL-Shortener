# What is URL-Shortener
URL Shortener allows users to input a long URL and receive a shortened version of it.

# How it is work 

it is so simple 

![Design](image.png)

and when the user try to click the new short link it is open the long link

# Project structure 
```
        url-shortener/
            ├── handlers/         # API logic for handling requests
            │   └── handlers.go
            ├── models/           # Data models
            │   └── url.go
            ├── router/           # Routing setup
            │   └── router.go
            ├── storage/          # Redis interactions
            │   └── redis-store.go
            ├── main.go           # Application entry point
            ├── Dockerfile        # Docker configuration
            ├── docker-compose.yml
            └── go.mod            # Go module file
```

# Learning resources 

1- 