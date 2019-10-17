# prometheus

[TOC]

https://prometheus.io

## 组件
- Prometheus Server 作为服务端，用来存储时间序列数据。
- 客户端库用来检测应用程序代码。
- 用于支持临时任务的推送网关。
- Exporter 用来监控 HAProxy，StatsD，Graphite 等特殊的监控目标，并向 Prometheus 提供标准格式的监控样本数据。
- alartmanager 用来处理告警。
- 其他各种周边工具。

## 架构

![架构](https://hugo-picture.oss-cn-beijing.aliyuncs.com/images/9Qt5yi.jpg)

## prometheus.yml

