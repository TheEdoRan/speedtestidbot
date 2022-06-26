# Build stage
FROM golang:1.18 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /speedtestidbot

# Run
FROM gcr.io/distroless/base-debian11 as runner
WORKDIR /app
COPY --from=builder /speedtestidbot /speedtestidbot
USER nonroot:nonroot
ENTRYPOINT ["/speedtestidbot"]