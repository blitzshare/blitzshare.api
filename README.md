# blitzshare.api

[Dockerhub](https://hub.docker.com/repository/docker/iamkimchi/blitzshare.api)

## Queue
https://docs.kubemq.io/getting-started/quick-start

kubectl apply -f https://deploy.kubemq.io/community
kubectl apply -f https://deploy.kubemq.io/key/0a5e3867-1149-40cf-b9f0-fe8321f52439

### kubemqctl
```
sudo curl -sL https://get.kubemq.io/install | sudo sh
```

### send message
```
kubemqctl queue send my-queue hello-world
```

### receive a message
```
kubemqctl queue send my-queue hello-world
```

### list queues 
```
kubemqctl queues list
```


## Local testing with minikube

minikube tunnel


resources
test: https://golang.hotexamples.com/examples/github.com.gin-gonic.gin/-/CreateTestContext/golang-createtestcontext-function-examples.html


IP=localhost
curl -i -XPOST http://localhost:80/p2p/registry -d '{"multiAddr": "ip/0.0.0.0/test", "oneTimePass":"secret-pass"}'
curl -i -XGET localhost/p2p/registry/secret-pass

curl -i -XGET localhost/p2p/registry/secret-pass
