FROM golang:alpine as builder

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN go get github.com/wqh66886/vue-gin-admin/server

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o ./run .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /go/src/app/run .
EXPOSE 8082
CMD [ "./run" ]