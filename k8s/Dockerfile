
FROM golang:1.23.2-alpine3.20 AS builder

RUN apk --no-cache add build-base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app

###
FROM alpine:latest

# Create a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the working directory
WORKDIR /app
COPY --from=builder /app/app .
RUN chmod +x ./app

# Change ownership of app files to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser
CMD ["./app"]