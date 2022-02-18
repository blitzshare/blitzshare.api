[![CircleCI](https://circleci.com/gh/blitzshare/blitzshare.api/tree/main.svg?style=svg&circle-token=364d84161031d4804629b88aa00dab075d3825fe)](https://circleci.com/gh/blitzshare/blitzshare.api/tree/main)


![logo](./assets/logo.png)

# blitzshare.api
Main public api responsible for Blitzshare business logic.


## Getting started

```bash
# install dependencies
$ make install
# start local server
$ make start
```

## Tests
```bash
# unit tests
$ make test
# acceptance tests
$ make acceptance-tests
# re/build mocks
$ make build-mocks
# generate test coverage report
$ make coverage-report-html

```

## Api doc generation
```bash
$ make swag-gen
# observe docs directory with generated docs
```

## K8s resources
```bash
# apply k8s resources
$ make k8s-apply
# destroy k8s resources
$ make k8s-destroy
```
## Db migrations
```bash
# create new migration
$ make migration-create name=api_key_init
# create database
$ make create-db
# delete database
$ make drop-db
# run migrations
$ make migrate-up
# crate api key record
$ make migrate-up
```

## Tools
[kubemqctl](https://docs.kubemq.io/getting-started/quick-start)

[kubectl](https://kubernetes.io/docs/reference/kubectl/overview/)