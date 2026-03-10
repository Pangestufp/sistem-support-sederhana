FROM golang:1.24-alpine AS builder

WORKDIR /app

# untuk module yg pakai git private
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

# copy semua source code
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# run step
FROM alpine:latest

WORKDIR /app

# copy binary dari builder
COPY --from=builder /app/main .

# expose port
EXPOSE 8080

# jalankan
CMD ["./main"]
