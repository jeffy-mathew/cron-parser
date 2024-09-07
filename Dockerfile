FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=builder /app/main /app/main

ENTRYPOINT ["./main"]