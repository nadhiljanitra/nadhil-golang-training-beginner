FROM golang:1.16-alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -tags musl -a -o engine app/main.go

FROM alpine

WORKDIR /app

EXPOSE 3000

COPY --from=builder /app/engine /app
CMD /app/engine rest