# Stage 1: Build go server
FROM golang:1.25 AS server-builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./main ./server/main.go

# Stage 2: Build node client
# FROM node:20-alpine AS client-builder
# WORKDIR /app
# RUN npm install -g pnpm
# COPY client/package.json client/pnpm-lock.yaml ./
# RUN pnpm install --frozen-lockfile
# COPY client/ .
# RUN pnpm run build

# Stage 3: live
FROM alpine:latest AS production
WORKDIR /app
RUN apk add --no-cache make
COPY --from=server-builder /app/main /app/main
# COPY --from=client-builder /app/dist /app/client/dist
ENTRYPOINT ["/app/main"]