FROM golang:1.14

WORKDIR /usr/src/milpost

COPY . .
RUN go get ./...

RUN go build
CMD ["./server"]
