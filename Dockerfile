
# Stage 1: Build
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-jobs-tracker ./cmd/scraper

# Stage 2: Run - super leve
FROM alpine:3.19
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=builder /go-jobs-tracker .
EXPOSE 2112
HEALTHCHECK --interval=30s --timeout=3s CMD wget --no-verbose --tries=1 --spider http://localhost:2112/healthz || exit 1
CMD ["./go-jobs-tracker"]
