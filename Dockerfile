FROM golang:1.20-alpine

WORKDIR /app
COPY . .

RUN go build -o bioinfo-aligner .

EXPOSE 8080

CMD ["./bioinfo-aligner"]
