FROM golang:alpine as builder

RUN mkdir -p /app
WORKDIR /app

# Required for aws-sdk
RUN apk add --no-cache \
  git mongodb-tools

COPY ./main.go /app/main.go

ENV GOPATH=/app
ENV GOBIN=/app/bin

RUN go get

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## RUNTIME #######
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /usr/bin/mongodump /usr/bin/mongodump
COPY --from=builder /app/main .

CMD ["./main"]
