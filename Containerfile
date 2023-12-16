FROM golang:1.21-alpine3.18 as build

WORKDIR /app

COPY . .

RUN go mod download && \
    go build -o joke-cli main.go

FROM alpine:edge
WORKDIR /app
COPY --from=build /app/joke-cli .
ENTRYPOINT ["./joke-cli"]