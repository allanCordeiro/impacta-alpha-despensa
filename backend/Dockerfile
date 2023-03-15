FROM golang:1.20.2-alpine AS builder

WORKDIR /usr/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /despensa-be cmd/api/main.go

FROM scratch
COPY --from=builder /despensa-be /despensa-be

EXPOSE 8000
ENTRYPOINT [ "/despensa-be" ]