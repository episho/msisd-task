FROM golang:1.12 as builder

ENV GO111MODULE=on

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/cmd ./cmd/server.go
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/cmd ./cmd/client.go


######## Start a new stage from scratch #######
FROM alpine:latest

WORKDIR /root/

EXPOSE 2323

#COPY --from=builder /go/bin/cmd ./client
COPY --from=builder /go/bin/cmd ./server

RUN apk --no-cache add ca-certificates
