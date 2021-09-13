# Stage 1 (Build Golang Project)
FROM golang:alpine as builder
ENV GO111MODULE=on
RUN mkdir /alta-store
ADD . /alta-store
WORKDIR /alta-store
COPY go.mod ./
RUN go mod download
RUN go clean --modcache
COPY . .
RUN go build -o main

# Start 2 (Reduce Size Without Golang Image)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /alta-store/main .
EXPOSE 8080
CMD ["./main"]