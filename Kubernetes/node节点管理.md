# node节点管理

[TOC]

## node的状态

节点的状态包括如下信息

- address
- - HostName: 可以被kubelet中的--hostname-override参数替代
- - ExternalIP: 可以被集群外部路由到的ip地址
- - InternalIP: 集群内部使用的ip,集群外部无法访问
- condition
- - OutOfDisk: 磁盘空间不足是为true
- - Ready: node controller40秒内无法收到node的状态报告为unknown,健康为true，否则为false
- - MemoryPressure： 当node有内存压力是为true,否则为false
- - DiskPressure: 当node有磁盘压力时为true,否则为false
- capacity
- - cpu
- - 内存
- - 可运行的最大pod数
- info 节点的版本信息,如os,kubernetes,docker等

## node管理

禁止pod调度到该节点

```shell
kubectl cordon <node>
kubectl uncordon <node>
```

驱逐该节点上的所有pod
```kubectl drain <node>```

该命令会删除该节点上的所有Pod（DaemonSet除外），在其他node上重新启动它们，通常该节点需要维护时使用该命令。直接使用该命令会自动调用kubectl cordon <node>命令。当该节点维护完成，启动了kubelet后，再使用kubectl uncordon <node>即可将该节点添加到kubernetes集群中

