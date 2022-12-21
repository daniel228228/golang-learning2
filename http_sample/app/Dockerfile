FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /app/test-service

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -race -o ./bin/test-service ./cmd/main.go


FROM alpine:latest

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /usr/share/zoneinfo/Europe/Moscow /usr/share/zoneinfo/Europe/Moscow
ENV TZ Europe/Moscow

WORKDIR /app/bin

COPY --from=builder /app/test-service/bin .

CMD [ "./test-service" ]