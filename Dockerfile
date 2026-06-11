FROM golang:1.24

WORKDIR /URL-SHORTENER

COPY go*.mod ./

RUN go mod tidy

COPY . .

ENV PORT=8080

EXPOSE 8080

RUN go build -o main .
CMD ["./main"]