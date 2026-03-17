# Build Stage
FROM golang:1.26.1-alpine3.23 AS builder
WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 go build -o /app/cli .

# Run Stage
FROM scratch
LABEL maintainer=Mwawaka version=1.0.0
# Copies from named stage, not host machine
COPY --from=builder /app/cli /app/cli 
ENTRYPOINT ["/app/cli"]
