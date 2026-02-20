# ---------- Stage 1: Build ----------
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY app/ .

# Компилируем статический бинарник
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o credit-risk-app main.go

# ---------- Stage 2: Runtime ----------
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/credit-risk-app .

CMD ["./credit-risk-app"]