
test:
	ENV=test && go test -v ./... -v -count=1

install:
	go install golang.org/x/tools/cmd/goimports@latest
	go mod vendor

fix-format:
	gofmt -w -s app/ pkg/ cmd/ mocks/ testhelpers
	goimports -w app/ pkg/ cmd/ mocks/ testhelpers

start:
	go run cmd/web/main.go
