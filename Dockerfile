FROM golang:1.14

WORKDIR /usr/src/milpost

COPY . .

RUN make build
CMD ["./server"]
