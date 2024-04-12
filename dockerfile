FROM golang:1.21

WORKDIR /http-user-service

COPY go.mod ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o hus

EXPOSE 9091

CMD ["./hus"]
