FROM golang:latest
RUN mkdir -p /app/files/
ADD . /app/
WORKDIR /app/cmd
RUN go build -o main .
CMD ["/app/cmd/main"]