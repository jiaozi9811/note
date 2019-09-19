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

## chart仓库

### 启动本地chart库

```
helm serve --address ip:port # 以http server的方式提供chart库
```

### 添加chart库

chart仓库用来存储和分享打包的chart，官方chart仓库由Kubernetes Charts维护， Helm允许创建私有chart仓库

chart仓库是一个可用来存储index.yml与打包的chart文件的HTTP server，当要分享chart时，需要上传chart文件到chart仓库。任何一个能能够提供YAML与tar文件的HTTP server都可以当做chart仓库，比如Google Cloud Storage (GCS) bucket、Amazon S3 bucket、Github Pages或创建你自己的web服务器

```
helm repo add <chart_name> <url>
helm repo add fabric8 https://fabric8.io/helm
```



# 使用

## 命令自动补全

```
yum install -y bash-completion
source /usr/share/bash-completion/completions/docker
sh /usr/share/bash-completion/bash_completion
source <(kubectl completion bash)
source <(helm completion bash)
```


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

## 版本回滚
 ```
 helm hist mysql # 查看版本
 helm rollback --debug mysql 1 #回滚
 ```
 
 ## 删除chart
 ```
 helm delete --purge mysql #彻底删除，release不会保留
 helm delete mysql #release会保留
 helm ls -a mysql # 确认删除
 ```
 
 
 
 
 
 
 
 
 
