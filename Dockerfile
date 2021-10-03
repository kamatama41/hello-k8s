FROM golang:1.17-alpine

COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o hello-server main.go

CMD ["/app/hello-server"]
