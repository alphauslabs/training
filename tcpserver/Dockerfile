FROM golang:latest

COPY . /app/

WORKDIR /app

RUN go mod init tcpserver

RUN go build -v

CMD ["./tcpserver"]
