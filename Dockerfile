FROM golang:alpine as builder

RUN mkdir -p /app
WORKDIR /app

# Required for aws-sdk
RUN apk add --no-cache \
  git

COPY ./main.go /app/main.go

ENV GOPATH=/app
ENV GOBIN=/app/bin

RUN go get

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## RUNTIME #######
FROM alpine:latest

RUN apk --no-cache add mongodb-tools
WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
