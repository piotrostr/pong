FROM golang:latest as builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -o app ./cmd/pong

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /build/app ./

EXPOSE 80
CMD [ "./app" ]
