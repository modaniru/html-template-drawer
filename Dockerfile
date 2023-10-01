FROM golang:latest

WORKDIR /app
COPY . .
VOLUME [ "/resources" ]
RUN go build cmd/main.go

ENTRYPOINT [ "/app/main" ]