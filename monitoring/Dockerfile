FROM golang:1.20-alpine
WORKDIR /app
COPY monitor.go .
RUN go build -o monitor monitor.go
CMD ["./monitor"]
