FROM golang:1.16 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/server/main.go

FROM scratch

COPY --from=builder /app/server .
ENTRYPOINT ["./server"]