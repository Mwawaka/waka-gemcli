#  === Build Stage ===
FROM golang:1.26.1-alpine3.23 AS builder
WORKDIR /app

# Cache modules for fast rebuilds
COPY go.mod go.sum /app/
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w -X main.version=1.0.0" -o /gemcli.

# === Run Stage ===
FROM scratch
LABEL maintainer="Mwawaka" \ 
      version="1.0.0" \
      description="Cyberpunk-themed Gemini CLI client"

# Copies from named stage, not host machine
COPY --from=builder /gemcli /gemcli 
ENTRYPOINT ["gemcli"]
