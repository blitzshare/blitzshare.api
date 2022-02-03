SHELL := /bin/bash
CWD := $(shell cd -P -- '$(shell dirname -- "$0")' && pwd -P)
export GO111MODULE := on
export GOBIN := $(CWD)/.bin

install:
	go install github.com/cucumber/godog/cmd/godog@v0.12.0
	go get -d github.com/vektra/mockery/v2/.../
	go mod vendor
start:
	go run cmd/main.go
test:
	go test  --tags='test' -v ./app/... -v -count=1 -cover -coverprofile=coverage.out
coverage-report-html:
	go tool cover -html=coverage.out		
acceptance-tests:
	cd "$(CWD)/test" && API_URL='http://0.0.0.0/api' ../.bin/godog ./**/*.feature
fix-format:
	gofmt -w -s app/  cmd/ mocks/ 
	goimports -w app/ cmd/ mocks/
build:
	GIN_MODE=release go build -o entrypoint cmd/main.go
k8s-apply:
	kubectl apply -f k8s/namespace.yaml
	kubectl apply -f k8s//deployment.yaml
	kubectl apply -f k8s/service.yaml
	kubectl apply -f k8s/hpa.yaml
	kubectl rollout restart deployment blitzshare-api-dpl --namespace blitzshare-ns
k8s-destroy:
	kubectl delete deployment blitzshare-api-dpl
build-mocks:
	.bin/mockery --all --dir "./app/"