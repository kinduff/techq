ARG IMAGE=scratch
ARG OS=linux
ARG ARCH=amd64

FROM golang:1.18.1-alpine3.15 as builder

WORKDIR /go/src/github.com/kinduff/techq

RUN apk --no-cache --virtual .build-deps add git alpine-sdk sqlite

COPY . .

RUN go mod download
RUN GOOS=$OS GOARCH=$ARCH go build -a -ldflags '-linkmode external -extldflags "-static"' -o binary .

FROM $IMAGE

LABEL name="techq"

WORKDIR /root/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/kinduff/techq/binary techq

EXPOSE 3000

CMD ["./techq"]
