FROM golang:1.17-alpine as builder
LABEL maintainer=jainpu@vmware.com
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /docker-gs-ping

FROM alpine:3.14
COPY --from=builder /docker-gs-ping /docker-gs-ping
EXPOSE 8080
CMD ["/docker-gs-ping"]
