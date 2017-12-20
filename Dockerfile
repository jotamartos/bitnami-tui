# Build project
FROM golang as builder
WORKDIR /app
COPY . /app
ENV GOPATH=/app:/go
RUN go get github.com/mattn/go-runewidth github.com/gdamore/tcell && \
    go build

# Create sample image
FROM busybox
WORKDIR /app
COPY --from=builder /app/app /app/out/app
ENTRYPOINT ["/app/out/app"]
