# istio

[TOC]

## service mesh(服务网格)

微服务从出现至今总共经历了三个阶段：微服务初期、Sidecar时期和Service Mesh时期

## istio
一个专注于处理服务间通信的基础设备层

由google，IBM和lyft公司共同研发。是service mesh的一个实现，可以看作是一个"微服务管理框架"，一般配好k8s使用

istio使用envoy作为sidecar，实现了Service Mesh对于微服务之间传输的诸多设想，并且添加了许多很棒的额外功能

## istio核心功能
- **流量管理** istio通过pilot提供的api动态的配置所有pod中sidecar的路由规则，进而控制服务间的流量和api调用
- **安全** istio提供给开发人员应用程序级别的安全性
- **可观察性** istio的mixer组件负责策略控制和遥测收集
- **平台独立** istio是独立于平台的。旨在运行在各种环境中，包括跨云，内部部署、kubernetes、mesos等
- **集成和定制** 策略执行组件可以扩展和定制，以便于和现有的acl，日志，监控，配额，审计等方案的集成


## istio框架
- **控制面** 负责管理和配置代理路由流量。通过mixer来实施策略和收集各个sidecar的遥测数据
- **数据面** 由一组sidecar构成。这些sidecar可以调节和控制微服务及mixer之间所有的网络通信

### 组件的功能
- sidecar(默认为Envoy)
- Pilot
- Mixer
