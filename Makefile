install:
	go install golang.org/x/tools/cmd/goimports@latest
	go mod vendor

test:
	ENV=test && go test -v ./... -v -count=1

fix-format:
	gofmt -w -s app/ pkg/ cmd/ mocks/ testhelpers
	goimports -w app/ pkg/ cmd/ mocks/ testhelpers

start:
	go run cmd/main.go

build:
	GIN_MODE=release go build -o entrypoint cmd/main.go

local-docker-build:
	docker build -t fileshare-api .
	docker tag fileshare-api:latest iamkimchi/blitzshare.fileshare.api:local-latest
	docker push iamkimchi/blitzshare.fileshare.api:local-latest

k8s-apply:
	kubectl apply -f k8s/config/namespace.yaml 
	kubectl apply -f k8s/config/deployment.yaml
	kubectl apply -f k8s/config/service.yaml

k8s-destory:
	kubectl delete namespace file-share-ns