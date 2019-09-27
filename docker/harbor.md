# harbor

[TOC]

https://github.com/goharbor/harbor
http://www.unixfbi.com/219.html

Harbor的所有服务组件都是在Docker中部署的，所以官方安装使用Docker-compose快速部署，所以我们需要安装Docker、Docker-compose。由于Harbor是基于Docker Registry V2版本，所以就要求Docker版本不小于1.10.0，Docker-compose版本不小于1.6.0

harbor由7个容器组成
- nginx: nginx 负责流量转发和安全验证，对外提供的流量都是从 nginx 中转，所以开放 https 的 443 端口，它将流量分发到后端的 ui 和正在 docker 镜像存储的 docker registry。
- harbor-jobservice：harbor-jobservice 是 harbor 的 job 管理模块，job 在 harbor 里面主要是为了镜像仓库之前同步使用的;
- harbor-ui：harbor-ui 是 web 管理页面，主要是前端的页面和后端 CURD 的接口;
- registry：registry 就是 docker 原生的仓库，负责保存镜像。
- harbor-adminserver：harbor-adminserver 是 harbor 系统管理接口，可以修改系统配置以及获取系统信息。
这几个容器通过 Docker link 的形式连接在一起，在容器之间通过容器名字互相访问。对终端用户而言，只需要暴露 proxy （即 Nginx）的服务端口。
- harbor-db：harbor-db 是 harbor 的数据库，这里保存了系统的 job 以及项目、人员权限管理。由于本 harbor 的认证也是通过数据，在生产环节大多对接到企业的 ldap 中；
- harbor-log：harbor-log 是 harbor 的日志服务，统一管理 harbor 的日志。通过 inspect 可以看出容器统一将日志输出的 syslog。

## docker-compose

harbor是通docker-compose安装的

## 安装harbor

将harbor的tar包解压,配置文件harbor.cfg最多修改hostname就可以使用

```sh install.sh```

harbor的日志在/var/log/harbor

harbor的数据在/data/下

## harbor启停
```
docker-compose ps
docker-compose start
docker-compose stop
docker-compose restart
docker-compose
docker-compose
```

## 镜像操作

```
docker login -u admin -p Harbor12345 111.230.25.113:80
vim /etc/docker/daemon.json
"insecure-registries": ["111.230.25.113:80"]
```

## harbor配置TLS
