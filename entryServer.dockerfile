FROM golang:1.19
WORKDIR /app

COPY server/ ./server
COPY go.mod go.sum .env ./

RUN go mod download
RUN go build -o filesServer ./server/cmd/main.go

CMD ["./filesServer"]