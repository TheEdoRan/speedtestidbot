# Build stage
FROM golang:1.27 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /speedtestidbot

# Run
FROM gcr.io/distroless/static-debian12 AS runner
WORKDIR /app
COPY --from=builder /speedtestidbot /speedtestidbot
USER nonroot:nonroot
ENTRYPOINT ["/speedtestidbot"]
