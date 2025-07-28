FROM golang:1.24.5-alpine

WORKDIR /app

RUN apk add --no-cache git && go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN rm -r docs/

RUN swag init -g cmd/main.go

RUN go build -o bin/app ./cmd

EXPOSE 8080

CMD ["./bin/app"]
