# helm

[TOC]

helm是kubernetes的包管理工具,用于管理charts(包管理资源)

helm chart是用来封装kubernetes原生应用程序的yaml文件。可以在部署应用时自定义应用程序的一些metadata，便于应用程序的分发

**helm和tiller的版本需相同**

## 安装

### 安装客户端

下载
https://github.com/helm/helm/releases

将helm二进制文件复制到/usr/local/bin
```
helm version
```

创建tiller的serviceaccount和clusterrolebinding
```
kubectl create serviceaccount --namespace kube-system tiller
kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
```

### 安装服务端tiller

tiller以deployment的方式部署在集群中

**在所有节点安装socat**

```
yum install socat -y
helm init --upgrade -i registry.cn-hangzhou.aliyuncs.com/google_containers/tiller:v2.12.2 --stable-repo-url https://kubernetes.oss-cn-hangzhou.aliyuncs.com/charts
kubectl patch deploy --namespace kube-system tiller-deploy -p '{"spec":{"template":{"spec":{"serviceAccount":"tiller"}}}}'

```

## 删除tiller
```
kubectl get -n kube-system secrets,sa,clusterrolebinding -o name|grep tiller|xargs kubectl -n kube-system delete
kubectl get all -n kube-system -l app=helm -o name|xargs kubectl delete -n kube-system
```

```helm reset```

# 使用


## 检查chart是否有效

```
helm install --dry-run --debug <chart_dir>`
```

## 常用命令

```
helm search mysql # 查询chart
helm create mysql # 创建chart
helm lint <chart_dir> # 检查chart的错误
helm package <chart_dir> -d <dir> # 打包chart
helm status mysql # 查看状态
helm list -a #

# 下载
helm install <chart_dir> # 部署chart
helm install stable/mysql #指定chart上的包
helm install chart.tgz # 指定chart包
helm install https://example.com/charts/chart.tgz # 指定url
```

下载chart
```
helm fetch stable/mysql --version 0.2.8 --untar # 获取版本，并解压
```
