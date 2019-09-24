# StorageClass

[TOC]

storageclass提供了描述存储"class(类)"的方法。不同的class可能会映射到不同的服务质量等级或备份策略

## storageclass资源

storageclass包括provisioner，parameters和reclaimPolicy字段。当class需要动态分配Pv时使用

- provisioner(存储分配器)
- parameters
- reclaimPolicy(回收策略)
- - Delete(default)
- - Retain

```
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: standard
provisioner: kubernetes.io/aws-ebs
parameters:
  type: gp2
reclaimPolicy: Retain
mountOptions:
  - debug
```
