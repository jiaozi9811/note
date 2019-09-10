# dashboard

[TOC]

https://github.com/kubernetes/dashboard

https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml

```
docker pull registry.cn-hangzhou.aliyuncs.com/tdycloud/kubernetes-dashboard-amd64:v1.10.1
docker tag registry.cn-hangzhou.aliyuncs.com/tdycloud/kubernetes-dashboard-amd64:v1.10.1 k8s.gcr.io/kubernetes-dashboard-amd64:v1.10.1
docker rmi registry.cn-hangzhou.aliyuncs.com/tdycloud/kubernetes-dashboard-amd64:v1.10.1
wget https://raw.githubusercontent.com/kubernetes/dashboard/master/src/deploy/recommended/kubernetes-dashboard.yaml -O dashboard.yaml
kubectl create -f dashboard.yaml
kubectl proxy --address='100.101.148.180' --port='8001' --accept-hosts='^*$'
```

## 访问证书
openssl pkcs12 -export -in admin.pem  -out k8s1.10.p12 -inkey admin-key.pem
将生成的k8s1.10.p12导出到windows,双击安装证书

### 生成token
kubectl -n kube-system describe secret $(kubectl get secret -n kube-system | grep kubernetes-dashboard-token | awk '{print $1}') | grep token:

## 打开代理，否则只能从本地访问
kubectl proxy --address=local_ip --disable-filter=true &

http://10.0.0.10:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/login


