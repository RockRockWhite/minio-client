FROM golang:latest

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release \
    PORT=80

WORKDIR /app

COPY . .

RUN go build .

EXPOSE 8080

ENTRYPOINT ["./minio-client"]