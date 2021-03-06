# volume

[TOC]

## kubernetes支持的volume类型
```
emptyDir
hostPath
gcePersistentDisk
awsElasticBlokStore
nfs
iscsi
flocker
glusterfs
rbd
cephfs
gitRepo
secret
persistentVolumeClaim
downwardAPI
azureFileVolume
azureDisk
vsphereVolume
Quobyte
PortworxVolume
ScaleIO
FlexVolume
StorageOS
local
```

## 本地数据卷(local volume)
local volume代表一个本地存储设备，比如磁盘，分区或目录等。主要的应用场景包括分布式存储和数据库等需要高性能和高可靠性的环境里
local volume同时支持块设备和文件系统，通过spec.local.path指定
local volume只能以静态创建的PV使用



## emptyDir
设置emptyDir类型volume，pod被分配到node上时，会创建emptyDir，只要pod运行在node上，emptydir就会存在(容器挂掉不会导致emptydir丢失数据)，但如果pod被删除或迁移，emptydir就会被删除，并永久丢失
```
apiVersion: v1
kind: Pod
metadata:
  name: test-pd
spec:
  containers:
  - image: gcr.io/google_containers/test-webserver
    name: test-container
    volumeMounts:
    - mountPath: /cache
      name: cache-volume
  volumes:
  - name: cache-volume
    emptyDir: {}
```    
    
## hostPath
hostPath允许挂载node上的文件系统到pod里，如果pod需要使用node上的文件，可以使用hostpath
```
apiVersion: v1
kind: Pod
metadata:
  name: test-pd
spec:
  containers:
  - image: gcr.io/google_containers/test-webserver
    name: test-container
    volumeMounts:
    - mountPath: /test-pd
      name: test-volume
  volumes:
  - name: test-volume
    hostPath:
      path: /data
 ```     

## NFS(network file system)
kubernetes通过简单的配置可以挂载nfs。nfs中的数据可以永久保存，同时nfs支持同时写操作
```
volumes:
- name: nfs
  nfs:
    server:10.254.234.223
    path: "/"
```    

## gcsPersistentDisk
gcsPersistentDisk可以挂载GCE上的永久磁盘到容器中，需要kubernetes运行在GCE的vm中
```
volumes:
  - name: test-volume
    gcsPersistentDisk:
      pdName: my-data-disk
      fsType: ext4
```      

## subPath
pod的多个容器使用同一个volume时，subpath非常有用
```
apiVersion: v1
kind: Pod
metadata:
  name:my-lamp-site
spec:
  containers:
  - name: mysql
    image: mysql
    volumeMounts:
    - mountPath: /var/lib/mysql
      name: site-data
      subPath: mysql
  - name: php
    image: php
    volumeMounts:
    - mountPath: /var/www/html
      name: site-data
      subPath: html
  volumes:
  - name: site-data
    persistentVolumeClaim:
      claimName: my-lamp-site-data
```
