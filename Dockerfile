FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o app ./cmd/server

EXPOSE 8080

CMD ["./app"]
