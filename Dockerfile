# FROM golang:1.22 AS builder

# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .

# RUN go build -o server ./cmd/server/main.go

# FROM alpine:latest
# WORKDIR /root/
# COPY --from=builder /app/server .

# CMD ["./server"]

FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Production stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

# Copiar binário do builder
COPY --from=builder /app/main .

# Copiar .env se existir, ou criar vazio
COPY --from=builder /app/cmd/server/.env .

EXPOSE 8080
ENTRYPOINT ["./main"]