FROM golang:alpine as builder
WORKDIR /go/src
COPY . .
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o server cmd/main.go

FROM alpine
WORKDIR /go/bin
COPY --from=builder /go/src .
EXPOSE 8080
CMD ["/go/bin/server"]