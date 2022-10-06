FROM golang:1.15-buster as builder

WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o /usr/local/bin/server ./example/server

EXPOSE 9998
CMD ["/usr/local/bin/server"]
