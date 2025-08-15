# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build binary statis untuk Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Run
FROM alpine:latest

WORKDIR /app

# Copy binary
COPY --from=builder /app/main .

# Copy .env (opsional, tapi sebaiknya mount dari VPS)
# COPY .env . 

EXPOSE 8004

CMD ["./main"]