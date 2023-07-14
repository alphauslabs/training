FROM golang:latest

WORKDIR /app

COPY main.go .

RUN go mod init main.go


RUN go build -o httpclient

EXPOSE 8080

ENTRYPOINT ["./httpclient"]