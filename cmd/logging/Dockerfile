# STEP-1
# compile application

FROM golang:1.24.2-alpine3.21 AS builder

RUN apk add git

WORKDIR /build

COPY ./main.go ./

RUN go mod init logging

RUN go mod tidy

RUN go build -o logging .


# STEP-2
# build app container

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /build/logging ./

ENTRYPOINT [ "/app/logging" ]
