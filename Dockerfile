FROM golang:1.11.2 as builder

ENV GO111MODULE=on

# Make a directory for building
CMD mkdir -p /go/src/users.git
WORKDIR /go/src/users.git

# Copy all the files into the directory and build it
COPY . /go/src/users.git/
WORKDIR /go/src/users.git
RUN CGO_ENABLED=0 go build -o app main.go

FROM alpine:latest
WORKDIR /bin/
COPY --from=builder /go/src/users.git/app .
CMD ["./app"]