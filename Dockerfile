FROM golang:latest
WORKDIR /app
COPY . /app
RUN go build main.go
CMD ["./main"]
