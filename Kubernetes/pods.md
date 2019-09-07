# pods

[TOC]

https://jimmysong.io/kubernetes-handbook/images/kubernetes-pod-cheatsheet.png

## pod的生命周期
* pending(挂起)
* running
* susseeded
* failed
* unknown

## pod状态
* PodScheduled
* Ready
* Initializwd
* Unschedulable
* ContainersReady

* CrashLoopBackOff： 容器退出，kubelet正在将它重启
* InvalidImageName： 无法解析镜像名称
* ImageInspectError： 无法校验镜像
* ErrImageNeverPull： 策略禁止拉取镜像
* ImagePullBackOff： 正在重试拉取
* RegistryUnavailable： 连接不到镜像中心
* ErrImagePull： 通用的拉取镜像出错
* CreateContainerConfigError： 不能创建kubelet使用的容器配置
* CreateContainerError： 创建容器失败
* m.internalLifecycle.PreStartContainer  执行hook报错
* RunContainerError： 启动容器失败
* PostStartHookError： 执行hook报错
* ContainersNotInitialized： 容器没有初始化完毕
* ContainersNotReady： 容器没有准备完毕
* ContainerCreating：容器创建中
* PodInitializing：pod 初始化中
* DockerDaemonNotReady：docker还没有完全启动
* NetworkPluginNotReady： 网络插件还没有完全启动


## 容器探针
### 探针
- **livenessProbe** 判断pod是否存活。
-- 如果探测到pod不健康，则kill pod，并执行重启策略
-- 如果没有liveness探针，则认为pod的livenessprobe的返回值永远为true
- **readinessProbe** 判断pod是否启动完成
-- 判断pod是否启动完成,即pod的Ready是否为true
-- 如果探测失败,控制器将pod的Endpoint从对应的service的Endpoint列表中移除,不再将请求调度到此pod,直到下次探测成功
### 探针的方法
* **EexcAction** 在容器内执行指定命令，如果命令退出时返回码为0,这判断成功
* **TCPSocketAction** 对指定端口上的容器的ip进行tcp检查，如果端口打开，则判断成功
* **HTTPGetAction** 对指定端口和路径上的ip执行http get请求，如果响应状态码大于等于200且小于400，则判断成功

### 探针的结果
- Success
- Failure
- Unknown 未能执行检查，因此不采取任何措施

### 重启策略 restartPolicy
- Always 默认
- OnFailure
- Never

### example

https://blog.51cto.com/newfly/2137136

- **periodSeconds** 检查的间隔时间，默认为10，单位：秒。
- **initialDelaySeconds** Pod启动后延迟多久才进行检查，单位：秒
- **timeoutSeconds** 探测的超时时间，默认为1，单位：秒。
- **successThreshold** 探测失败后认为成功的最小连接成功次数，默认为1，在Liveness探针中必须为1，最小值为1。
- **failureThreshold** 探测失败的重试次数，重试一定次数后将认为失败，在readiness探针中，Pod会被标记为未就绪，默认为3，最小值为1。

```
spec:
  containers:
  - name: liveness
    image: k8s.gcr.io/busybox
    args:
    - /bin/sh
    - -c
    - touch /tmp/healthy; sleep 30; rm -rf /tmp/healthy; sleep 600
    livenessProbe:
      exec:
        command:
        - cat
        - /tmp/healthy
      initialDelaySeconds: 5
      periodSeconds: 5
```


