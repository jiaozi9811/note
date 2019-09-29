# istio

[TOC]

https://archive.istio.io/v1.2/zh/docs/concepts/what-is-istio/

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
- - envoy是c++开发的高性能代理，用于调解服务网格中服务的入站和出站流量。istio中envoy用于sidecar，和对应的服务部署在同一个pod中。envoy调解所有出入服务的流量。所有经过envoy的流量都会调用mixer，为mixer提供一组描述请求和请求周围环境的attribute。根据envoy的配置和attribute，mixer会调用各种后台的基础设施资源。而这些attribute又可以在mixer中用于决策使用何种策略，并发送给监控系统，以提供整个网格行为的信息
- Pilot
- - pilot为sidecar提供**服务发现**功能，并管理高级路由(如A/B测试和金丝雀部署)和故障处理(超时、重试、熔断器等)的流量。pilot将这些"高级"的流量行为转换为详尽的sidecar配置项，并在运行时将它们配置到sidecar中
- - Pilot将服务发现机制提炼为供数据面使用的API，即任何Sidecar都可以使用的标准格式。这种松耦合的设计模式使Istio能在多种环境(Kubernetes、Consul 和 Nomad)下运行，同时保持用于流量管理操作的相同
- Mixer
- - mixer是独立于平台的组件。通过从sidecar和一些其他服务处收集数据，进而在整个service mesh上控制访问和执行策略。sidecar请求级别的attribute被发送到mixer进行评估
- - mixer中还包括一个灵活的插件，使其能接入各种主机环境和基础设施的后段，并得到sidecar代理和istio所管理的服务
- Citadel
- - citadel通过内置身份和凭证管理提供"服务间"和"最终用户"身份验证，cidadel可用于升级服务网格中未加密的流量，并能为运维人员提供基于服务标识而不是网络层的强制执行策略

![Istio service mesh 架构图](https://jimmysong.io/istio-handbook/images/006tNc79ly1fz73sprcdlj31580u046j.jpg)
















