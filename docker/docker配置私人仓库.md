
## 修改配置文件  
修改配置文件/etc/default/docker  
DOCKER_OPTS="--insecure-registry ****"  
重启docker

## 将私人仓库添加到信任仓库
vim /etc/docker/daemon.json
增加需要免认证的仓库地址  
{"insecure-registries":["***","***"]}  
重启docker  
查看docker信息 docker info  
查看Insecure Registries: 中是否添加了私人仓库  
