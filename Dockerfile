# Stage 1: base
FROM golang:1.23 AS base
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .

# Stage 2: development (APP_ENV=development)
FROM base AS development
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/air-verse/air@latest
ENTRYPOINT ["air", "-c", "air.toml"]

# Stage 3: build
FROM base AS build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install github.com/pressly/goose/v3/cmd/goose@latest
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build

# Stage 3.5: test
# FROM build AS test
# RUN make test

# Stage 4: production (APP_ENV=production)
FROM alpine:latest AS production
WORKDIR /app
RUN apk add --no-cache make
COPY --from=build /go/bin/goose /usr/local/bin/goose
COPY --from=build /app/goose /app/goose
COPY --from=build /app/Makefile /app/Makefile
COPY --from=build /app/tmp/api/main /app/api/main
ENTRYPOINT ["/app/api/main"]