# Build State
FROM golang:1.20-alpine3.18 AS BUILDER
WORKDIR /app
COPY . .
RUN go build -o app app.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Run State
FROM alpine:3.17
WORKDIR /app
COPY --from=BUILDER /app/app .
COPY --from=BUILDER /app/migrate .
COPY app.env .
COPY start.sh .
COPY src/db/migration ./src/db/migration

EXPOSE 8080
CMD ["/app/app"]
ENTRYPOINT ["/app/start.sh"]