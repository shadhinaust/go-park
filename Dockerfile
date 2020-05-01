FROM golang:alpine
WORKDIR /go/src/app
COPY . .
CMD ["go", "run", "main.go"]