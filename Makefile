SHELL := /bin/bash
CWD := $(shell cd -P -- '$(shell dirname -- "$0")' && pwd -P)

export GO111MODULE := on
export GOBIN := $(CWD)/.bin

DB_CONNECTION="postgresql://postgres:postgres@0.0.0.0:5438/identity?sslmode=disable"

install:
	go install $(shell go list -f '{{join .Imports " "}}' tools.go)
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
	kubectl apply -f k8s/deployment.yaml
	kubectl apply -f k8s/service.yaml
	kubectl apply -f k8s/hpa.yaml
	kubectl apply -f k8s/keystore-db.yaml
	kubectl rollout restart deployment blitzshare-api-dpl --namespace blitzshare-ns
k8s-destroy:
	kubectl delete deployment blitzshare-api-dpl
build-mocks:
	.bin/mockery --all --dir "./app/"
swag-gen:
	.bin/swag init --dir app --generalInfo routes/init.go -o docs
migration-create:
	[ -z "$(name)" ] && echo "Please provide migration name." && exit 1 || :
	.bin/migrate create -seq -ext sql -dir app/services/db/migration "$(name)"
migrate-up:
	.bin/migrate -path app/services/db/migration -database "$(DB_CONNECTION)" -verbose up
sqlc:
	.bin/sqlc generate
create-db:
	psql "$DB_CONNECTION_NODB" -c "CREATE DATABASE identity;"