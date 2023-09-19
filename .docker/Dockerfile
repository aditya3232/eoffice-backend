FROM golang:1.20.1-alpine3.16 AS builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /go/bin/eoffice-v2-be main.go

FROM alpine:latest
COPY ./config/.env  /config/.env
COPY ./config/config.go  /config/config.go
COPY --from=builder /go/bin/eoffice-v2-be .
EXPOSE 8080
ENTRYPOINT [ "./eoffice-v2-be" ]
