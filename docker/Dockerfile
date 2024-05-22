FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/ ./cmd/...

FROM scratch

WORKDIR /

COPY --from=builder /app/bin .

ENV XDG_DATA_HOME /data
ENV XDG_CONFIG_HOME /config

CMD ["/caddy", "run", "--config", "/Caddyfile", "--adapter", "caddyfile"]
