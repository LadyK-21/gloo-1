FROM golang:alpine AS builder

RUN apk --no-cache add make
WORKDIR /app
COPY . .

FROM alpine:3.18.3
RUN ls
COPY ./server /server
CMD ["./server"]
