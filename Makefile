SHELL := /bin/bash
CWD := $(shell cd -P -- '$(shell dirname -- "$0")' && pwd -P)
export GO111MODULE := on
export GOBIN := $(CWD)/.bin

install:
	go install github.com/cucumber/godog/cmd/godog@v0.12.0
	go get -d github.com/vektra/mockery/v2/.../
	go mod vendor

test:
	go test -v ./app/... -v -count=1 -cover
	
acceptance-tests:
	cd "$(CWD)/test" && API_URL='http://0.0.0.0/api' ../.bin/godog ./**/*.feature

fix-format:
	gofmt -w -s app/  cmd/ mocks/ 
	goimports -w app/ cmd/ mocks/

start:
	go run cmd/main.go

build:
	GIN_MODE=release go build -o entrypoint cmd/main.go

k8s-apply:
	kubectl apply -f k8s/namespace.yaml
	kubectl apply -f k8s//deployment.yaml
	kubectl apply -f k8s/service.yaml
	kubectl apply -f k8s/hpa.yaml
	kubectl rollout restart deployment blitzshare-api-dpl --namespace blitzshare-ns

k8s-destroy:
	kubectl delete namespace blitzshare-ns


build-deploy:
	make dockerhub-build
	make k8s-apply

docker-build:
	docker build -t  blitzshare.api:latest .
	
	
dockerhub-build:
	make dockerhub-build
	docker tag blitzshare.api:latest iamkimchi/blitzshare.api:latest
	docker push iamkimchi/blitzshare.api:latest

build-mocks:
	.bin/mockery --all --dir "./app/"