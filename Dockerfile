FROM golang:latest

WORKDIR /app-src

COPY go.mod go.sum ./

RUN go mod download

COPY cli cli

RUN go build -o /app cli/main.go

COPY . .

ENTRYPOINT ["/app"]
CMD ["api"]
