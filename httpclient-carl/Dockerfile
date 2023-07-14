FROM golang:latest

WORKDIR /app

COPY main.go .
RUN go mod init main.go

RUN go build -o httpclient-carl
EXPOSE 8080

ENTRYPOINT ["./httpclient-carl"]
