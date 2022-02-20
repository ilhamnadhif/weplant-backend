FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o weplant-backend

EXPOSE 8080

CMD ["./weplant-backend"]