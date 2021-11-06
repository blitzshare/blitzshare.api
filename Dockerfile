FROM golang:1.17.2-alpine AS builder
WORKDIR /go/src
ADD . /go/src
RUN go install golang.org/x/tools/cmd/goimports@latest
RUN go mod vendor
RUN go build -o /app/entrypoint cmd/main.go # GIN_MODE=release 

FROM alpine
WORKDIR /app
COPY --from=builder /app/entrypoint /app/
EXPOSE 8000
ENTRYPOINT [ "/app/entrypoint"]