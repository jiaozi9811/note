# Persistent Volume(持久化卷)

[TOC]

**persistentVolume(pv)** 由管理员设置的存储,是集群的一部分

PV是volume之类的卷插件，但具有独立于使用PV的pod的生命周期。此 API 对象包含存储实现的细节，即 NFS、iSCSI 或特定于云供应商的存储系统

**persistentVolumeClaim(PVC)** 是用户存储的请求

