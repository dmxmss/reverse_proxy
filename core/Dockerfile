FROM golang:1.24.3-alpine3.22 as builder
WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app .

FROM scratch
WORKDIR /app

COPY --from=builder /usr/src/app/app .
COPY --from=builder /usr/src/app/templates ./templates

EXPOSE 3000

CMD ["./app"]
