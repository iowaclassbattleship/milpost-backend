FROM golang:1.14

WORKDIR /usr/src/milpost

COPY . .
RUN go get ./...

RUN make build
CMD ["./server"]
