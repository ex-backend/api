FROM golang AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLE=0 GOOS=linux go build -o api

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app .
EXPOSE 8080

ENTRYPOINT "/app/api"
