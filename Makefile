install:
	go install golang.org/x/tools/cmd/goimports@latest
	go get -d github.com/vektra/mockery/v2/.../
	go mod vendor

test:
	go test -v ./app/... -v -count=1 -cover


acceptance-tests:
	go test -v ./test/... -v -count=1

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
	kubectl rollout restart deployment blitzshare-api-deployment --namespace blitzshare-ns

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

minikube-svc:
	minikube service blitzshare-api-svc -n blitzshare-ns

build-mocks:
	mockery --all --dir "./app/"