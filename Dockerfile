# Build project
FROM golang as builder
WORKDIR /app
COPY . /app
RUN go get github.com/tools/godep && \
    godep restore && \
    go build

# Create sample image
FROM busybox
WORKDIR /app
COPY --from=builder /app/app /app/out/app
ENTRYPOINT ["/app/out/app"]
