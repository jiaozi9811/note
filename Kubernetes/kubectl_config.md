#kubectl config

[TOC]

kubectl config操作的是$HOME/.kube/config文件

## 查看config kubectl config view

```yaml
# kubectl config view
apiVersion: v1
clusters: []
contexts: []
current-context: ""
kind: Config
preferences: {}
users: []

# cat /root/.kub/config
```

## config主要配置项

* clusters ：配置要访问的kubernetes集群
* contexts ：配置访问kubernetes集群的具体上下文环境
* current-context: 配置当前使用的上下文环境
* users： 配置访问的用户信息，用户名以及证书信息
