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

local-minikube-build:
	docker build -t fileshare-api:latest .
	docker tag fileshare-api:latest iamkimchi/blitzshare.fileshare.api:local-latest
	minikube image load fileshare-api:latest
	

local-docker-build:
	# docker build -f ./Dockerfile.local -t fileshare-api:latest .
	docker build -t fileshare-api:latest .
	docker tag fileshare-api:latest iamkimchi/blitzshare.fileshare.api:local-latest
	docker push iamkimchi/blitzshare.fileshare.api:local-latest
	make k8s-apply
	

k8s-apply:
	# kubectl delete namespace file-share-ns
	kubectl apply -f k8s/config/namespace.yaml 
	kubectl apply -f k8s/config/deployment.yaml
	kubectl apply -f k8s/config/service.yaml

k8s-destory:
	kubectl delete namespace file-share-ns
k8s-pf:
	kubectl port-forward $(kubectl get pods  | tail -n1 | awk '{print $1}') 8000:80

docker-img-cleanup:
	sudo docker rmi -f $(sudo docker images -a -q)
	sudo docker rm -vf $(sudo docker ps -a -q)
