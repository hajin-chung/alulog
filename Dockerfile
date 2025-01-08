FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /server ./cmd/server

FROM alpine:3.18

WORKDIR /app
COPY --from=builder /server /app/server
EXPOSE 3000

CMD ["./server"]
