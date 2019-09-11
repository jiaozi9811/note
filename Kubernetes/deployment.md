# deployment

[TOC]

![Deployment 结构示意图](https://jimmysong.io/kubernetes-handbook/images/deployment-cheatsheet.png)

## deployment
deployment为pod和replicaset提供了一个声明式定义(declarative)方法，用来替代以前的replicationcontroller来方便的管理应用
只需在deploymeng中描述想要的目标状态，deployment controller就会将pod和replica set的实际状态改变到目标状态。
可以定义一个全新的deployment，也可创建一个新的deployment，以替换旧的

应用场景
- 定义deployment来创建pod和replicaset
- 滚动升级和回滚应用
- 扩容和缩容
- 暂停和继续deployment

```
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.7.9
        ports:
        - containerPort: 80
```
        
## 扩容
```kubectl scale deployment nginx-deployment --replicas 10```

如果集群支持horizontal pod autoscaling，deployment可设置自动扩展
```kubectl autoscale deployment nginx-deployment --min=10 --max=15 --cpu-percent=80```

## 更新镜像
```kubectl set image deployment/nginx-deployment nginx=nginx:1.9.1```

## 回滚
```kubectl rollout undo deployment/nginx-deployment```

deployment当且仅当deployment的pod template中的label更新或镜像更改时被触发。其他更新不会触发rollout


## 暂停和恢复deployment
触发一次或多次更新前暂停一个deployment，然后再恢复它。在此期间进行一些恢复工作，并不会触发不必要的rollout

## 暂停deployment
```kubectl rollout pause deployment/nginx-deployment```
## 更新deployment
```kubectl set image deploy/nginx nginx=nginx:1.9.1```
## 恢复deployment
```kubectl rollout resume deploy nginx```
