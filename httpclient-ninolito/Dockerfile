FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod init httpclient

RUN go build -v

EXPOSE 8080

CMD ["./httpclient"]
