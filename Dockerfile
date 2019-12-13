FROM golang:1.12-alpine as builder

WORKDIR /go/app

ENV GOOS="linux"
ENV GOARCH="amd64"


RUN apk add --no-cache git

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . . 

RUN apk add --no-cache \
    && go build -o app

FROM alpine

WORKDIR /app

COPY --from=builder /go/app/app .

RUN addgroup go \
    && adduser -D -G go go \
    && chown -R go:go /app/app

CMD ["./app"]
