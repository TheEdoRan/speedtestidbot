# Build stage
FROM golang:1.17 AS builder
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
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/speedtestidbot"]