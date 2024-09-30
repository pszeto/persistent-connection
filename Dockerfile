FROM golang:1.23.1-bullseye AS builder
WORKDIR /
COPY go.* ./
RUN go mod download
COPY main.go ./
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM golang:1.23.1-alpine3.20

COPY --from=builder /main /main
EXPOSE 8000
ENTRYPOINT ["/main"]