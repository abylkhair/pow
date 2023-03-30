FROM golang:1.17-alpine

WORKDIR /app
COPY . /app

RUN go build -o main .

CMD ["./main"]