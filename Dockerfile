FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY ./cmd/app ./cmd/app
COPY ./db ./db
COPY ./internal ./internal

RUN go build -v -o /app/main /app/cmd/app/main.go

CMD ["/app/main"]