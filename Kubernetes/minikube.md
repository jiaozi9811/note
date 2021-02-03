# minikube

[TOC]

## install

```
sudo usermod -aG docker $USER && newgrp docker

minikube start --driver=docker --cpus=2 --disk-size='10g' --image-mirror-country='cn' --image-repository='registry.cn-hangzhou.aliyuncs.com/google_containers' 
```

## Dashboard对外暴露访问链接
```
minikube dashboard
minikube addons enable dashboard
kubectl proxy --port=8001 --address=‘10.211.55.6’ --accept-hosts=’^.*’ &
http://10.211.55.6:8001/api/v1/namespaces/kubernetes-dashboard/services/http:kubernetes-dashboard:/proxy/#/overview?namespace=default
```
