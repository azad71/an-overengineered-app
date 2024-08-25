# Stage 1: Build the application
FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk --no-cache add git bash

# Install go-migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.1


COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the application
RUN go build -o an-overengineered-app .

# Stage 2: Create the final image
FROM alpine:3.18

WORKDIR /app

RUN apk --no-cache add ca-certificates bash postgresql-client 

# Copy the built application and other files
COPY --from=builder /app/an-overengineered-app .
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY ./db/migrations /app/db/migrations
COPY init.sh /app/init.sh

COPY .env.docker /app/.env.docker

COPY certs /app/certs


# Set environment variables
ENV APP_ENV=docker

RUN chmod +x /app/init.sh

EXPOSE 8080

ENTRYPOINT [ "/app/init.sh" ]
