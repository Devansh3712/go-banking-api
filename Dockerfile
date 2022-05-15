FROM golang:alpine

WORKDIR /go/src/github.com/Devansh3712/go-banking-api
COPY . .

RUN apk update && apk add --no-cache git
RUN go get ./...
EXPOSE 8000

CMD ["go", "run", "main.go"]
