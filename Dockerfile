FROM golang:latest

WORKDIR /app-src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app main.go

ENTRYPOINT ["/app"]
