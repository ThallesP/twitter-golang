FROM golang:1.20-alpine
WORKDIR /app
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
RUN go mod download
COPY . /app
ENV CGO_ENABLED 0
RUN go build . -o bin

FROM alpine:3.14
RUN apk add ca-certificates
WORKDIR /app
COPY --from=0 /app/bin /app/bin
CMD /app/bin