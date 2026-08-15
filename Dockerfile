FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

FROM alpine:3.20
WORKDIR /app

COPY --from=builder /app/app /app/app
COPY --from=builder /app/templates ./templates

EXPOSE 8080

ENV APP_PORT=8080 \
    DB_HOST=host.docker.internal \
    DB_PORT=5432 \
    DB_USER=postgres \
    DB_PASSWORD=norimorienair4614 \
    DB_NAME=gestion_documentos \
    DB_SSLMODE=disable

CMD ["/app/app"]
