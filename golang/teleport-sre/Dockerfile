FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go build -oteleport-sre ./cmd/server

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/teleport-sre /teleport-sre
ENTRYPOINT ["/teleport-sre"]