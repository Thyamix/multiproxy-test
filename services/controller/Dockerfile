FROM golang:1.26.2 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/main.go

FROM gcr.io/distroless/static-debian12

COPY --from=builder /main /main

EXPOSE 8080

ENTRYPOINT ["/main"]
