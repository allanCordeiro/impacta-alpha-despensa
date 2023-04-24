FROM golang:1.20.3 AS builder

WORKDIR /usr/src/app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /despensa-be cmd/api/main.go

FROM scratch
COPY --from=builder /despensa-be /despensa-be
COPY --from=builder /usr/src/app/sql/migrations /sql/migrations
COPY --from=builder usr/src/app/.env.example /.env

EXPOSE 8000
ENTRYPOINT [ "/despensa-be" ]