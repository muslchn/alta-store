# FROM golang:1.17.1-alpine3.14
# RUN mkdir /alta-store
# ADD . /alta-store
# WORKDIR /alta-store
# RUN go clean --modcache
# RUN go build -o main
# EXPOSE 8080
# CMD ["/alta-store/main"]

# Stage 1 (Build Golang Project)
FROM golang:1.17.1-alpine3.14 as builder
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
FROM alpine
WORKDIR /root/
COPY --from=builder /alta-store/main .
EXPOSE 8080
CMD ["./main"]