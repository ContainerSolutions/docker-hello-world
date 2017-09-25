FROM golang:1.9-alpine as builder
WORKDIR /go/src/github.com/ContainerSolutions/docker-hello-world/
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.6
EXPOSE 80
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/ContainerSolutions/docker-hello-world/main .
CMD ["./main"]
