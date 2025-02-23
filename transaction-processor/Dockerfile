FROM golang:1.22.3

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o main .

CMD ["./main"]
