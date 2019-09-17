# Persistent Volume(持久化卷)

[TOC]

**persistentVolume(pv)** 由管理员设置的存储,是集群的一部分。PV是volume之类的卷插件，但具有独立于使用PV的pod的生命周期。此 API 对象包含存储实现的细节，即 NFS、iSCSI 或特定于云供应商的存储系统

**persistentVolumeClaim(PVC)** 是用户存储的请求。PVC与pod相似，pod消耗node资源，PVC消耗PV资源

**PV属于集群中的资源，PVC是对PV的请求，也作为对资源请求的检查**

## PV的生命周期

### PV的配置
有两种方式配置PV
- 静态
- 动态

### PVC绑定PV
PVC绑定PV是排他性的,是一对一映射的


### 使用


### PVC的保护
目的是确保pod使用的PVC不会从系统中移除。

当pod处于pending状态且pod已分配给节点，或pod为running时，PVC处于活跃状态

当启用PVC的保护功能时，如果删除一个正在使用的PVC，此PVC不会立即被删除，直到PVC不被是、任何pod使用时才会被删除

### 回收与保留






