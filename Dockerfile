FROM golang:1.17.2-alpine AS builder
WORKDIR /go/src
ADD . /go/src
RUN go install golang.org/x/tools/cmd/goimports@latest
RUN go mod vendor
RUN GIN_MODE=release go build -o entrypoint cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/entrypoint /app/
EXPOSE 63785
ENV GIN_MODE=release
ENTRYPOINT ["/app/entrypoint"]