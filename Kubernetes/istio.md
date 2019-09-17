# istio

[TOC]

## service mesh(服务网格)

微服务从出现至今总共经历了三个阶段：微服务初期、Sidecar时期和Service Mesh时期

## istio
一个专注于处理服务间通信的基础设备层

由google，IBM和lyft公司共同研发。是service mesh的一个实现，可以看作是一个"微服务管理框架"，一般配好k8s使用

istio使用envoy作为sidecar，实现了Service Mesh对于微服务之间传输的诸多设想，并且添加了许多很棒的额外功能

## istio核心功能
- **流量管理**
- **安全**
- **可观察性**
- **平台独立**
- **集成和定制**
