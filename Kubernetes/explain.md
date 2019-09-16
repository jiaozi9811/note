# 资源配置清单

[TOC]

kubectl explain pods.spec.containers
kubectl explain pods|rs

资源的清单格式
```
    一级字段
        apiVersion(group/version)
        kind(类别)
        metadate(元数据)
            name
            namespace
            labels
            annotations
        spec
```
pod资源
**kubectl explain pods.spec.containers**
```
    spec.containers []object 

    -name [string]
     image [string]
     imagePullPolicy:[Always|Never|IfNotPresent]
     env
     ports
     volumeDevices
     volumeMounts
     workingDir

     command
     args
```
