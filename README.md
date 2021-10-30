# blitzshare.fileshare.api

[Dockerhub](https://hub.docker.com/repository/docker/iamkimchi/blitzshare.fileshare.api)

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

## init kubemq

https://account.kubemq.io/home/get-kubemq/kubernetes


## teleprecence
sudo curl -fL https://app.getambassador.io/download/tel2/linux/amd64/latest/telepresence -o /usr/local/bin/telepresence
