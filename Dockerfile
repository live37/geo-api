FROM golang:1.24 as builder
WORKDIR /app

# ENV GOPROXY https://goproxy.cn

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build .

FROM alpine:3.10
WORKDIR /app

COPY --from=builder /app/geo-api ./run
COPY --from=builder /app/config.yaml ./config.yaml
COPY --from=builder /app/static ./static


EXPOSE 8080
CMD ["./run"]