FROM golang:1.17.1-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main
EXPOSE 8080
CMD ["/app/main"]