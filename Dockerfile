FROM golang:1.26 AS builder
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 


COPY ./cmd/ ./cmd
COPY ./internal/ ./internal


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tetris ./cmd/server/main.go



FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/tetris . 
COPY ./static/ ./static/


EXPOSE 9090

CMD ["./tetris"]
