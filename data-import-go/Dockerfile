FROM golang:1.21-alpine AS builder

RUN apk update && apk add xsltproc

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main main.go

CMD ["/main", "cfr.xsl", "example.xml"]