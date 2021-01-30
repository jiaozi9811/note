# minikube

[TOC]

## install

```
sudo usermod -aG docker $USER && newgrp docker

minikube start  --cpus=2 --disk-size='10g' --image-mirror-country='cn' --image-repository='registry.cn-hangzhou.aliyuncs.com/google_containers' 
```
