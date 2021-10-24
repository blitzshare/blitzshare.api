FROM golang:1.17.2-alpine

WORKDIR /app

COPY . .

RUN go install golang.org/x/tools/cmd/goimports@latest
RUN go mod vendor
RUN go build -o entrypoint cmd/main.go # GIN_MODE=release 

EXPOSE 8000

CMD ls

# ENTRYPOINT [ "main"]
ENTRYPOINT [ "/app/entrypoint"]