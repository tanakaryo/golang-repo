FROM golang:alpine3.21

WORKDIR /app

COPY . .

RUN go build -o myapp

EXPOSE 8080

CMD ["./myapp"]