FROM golang:1.17

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
COPY .env.docker .env

RUN go build -o main .

EXPOSE 8085

CMD ["./main"]
