ARG GO_VERSION=1.18

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /gin-ticker
WORKDIR /gin-ticker

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./app ./main.go

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /gin-ticker
WORKDIR /gin-ticker
COPY --from=builder /gin-ticker/app .
#COPY --from=builder /gin-ticker/test.db .

EXPOSE 5000

ENTRYPOINT ["./app"]