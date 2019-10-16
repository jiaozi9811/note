# kubernetes镜像拉取失败解决方法

[TOC]

https://youendless.com/post/kubernetes_image_pull/

Docker Hub以及利用开源harbor项目搭建的镜像仓库服务，对于Docker Client发起的docker login、docker push、docker pull等命令都会做基本的用户认证， 最简单常用的认证方式就是Basic Auth，即在发起的http请求头中添加一个Authorization，其值为base64(username:password)，当前Docker Client都是这么处理

在Kubernetes中，Secret资源对象用来存储和管理一些敏感信息，比如密码、Auth Token以及SSH keys，把这些敏感信息放入Secret对象中，相对来说更安全更灵活。 Kubernetes可以通过环境变量、文件挂载等方式将Secret信息推到每一个Pod中，通过文件挂载形式还能使Secret在pod中实时更新，Kubernetes统一管理










