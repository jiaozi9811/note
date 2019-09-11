# statefulset

[TOC]

statefulset是为了解决有状态服务的问题，应用场景包括：
- 稳定的持久化存储，即pod重新调度后还能访问到相同的持久化数据，基于pvc实现
- 稳定的网络标志，即pod重新调度后其podname和hostname不变，基于headless service实现
- 有序部署/有序扩展，即pod是有顺序的，在部署或扩展的时候要依据定义的顺序依次依序进行，基于init containers实现
- 有序收缩，有序删除

组成
- 用于定义网络标志(DNS domain)的headless service
- 用于创建PersistentVolumes的volumeClaimTemplates
- 定义具体应用的statefulSet
    
## 更新statefulset
v1.7+支持statefulset的自动更新，通过spec.updateStrategy设置更新策略。目前支持两种策略
- onDelete:当.spec.template更新时，并不立即删除旧的pod，而是等待用户手动删除旧pod后，自动创建新pod。这是默认策略，兼容v1.6版本的行为
- RollingUpdate:当.spec.template更新时，自动删除旧pod并创建新pod替换。更新时，这些pod按逆序进行，依次删除，创建并等待pod变成ready状态才进行下一个pod的更新

## 组件
- 一个headless service用于控制网络域
- 一个statefulset,它的Spec中指定在有3个运行nginx容器的Pod
- volumeClaimTemplates 使用 PersistentVolume Provisioner 提供的 PersistentVolumes 作为稳定存储
```
apiVersion: v1
kind: Service
metadata:
  name: nginx
  labels:
    app: nginx
spec:
  ports:
  - port: 80
    name: web
  clusterIP: None
  selector:
    app: nginx
---
apiVersion: apps/v1beta1
kind: StatefulSet
metadata:
  name: web
spec:
  serviceName: "nginx"
  replicas: 3
  template:
    metadata:
      labels:
        app: nginx
    spec:
      terminationGracePeriodSeconds: 10
      containers:
      - name: nginx
        image: gcr.io/google_containers/nginx-slim:0.8
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
  volumeClaimTemplates:
  - metadata:
      name: www
      annotations:
        volume.beta.kubernetes.io/storage-class: anything
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 1Gi
```
