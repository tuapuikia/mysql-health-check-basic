FROM golang:alpine AS builder

WORKDIR /opt/builder

COPY . /opt/builder

RUN go mod tidy

RUN CGO_ENABLED=0 go build

FROM alpine:latest

COPY --from=builder /opt/builder/mysql-go /usr/local/bin/mysql-go

ENTRYPOINT ["/usr/local/bin/mysql-go"]
