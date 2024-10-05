FROM golang:1.23-alpine3.20 AS builder
##Нужна исключительно для сборки нашего приложения

COPY . /github.com/Chigvero/source/
WORKDIR /github.com/Chigvero/source/

RUN go mod tidy
RUN go mod download
##Будет скачивать все зависимости из go.mod

RUN go build -o ./bin/crud_server cmd/grpc_server/main.go

FROM alpine:latest
##Нужна для работы нашего приложения

WORKDIR /root/

COPY --from=builder /github.com/Chigvero/source/bin/crud_server .

CMD ["./crud_server"]

