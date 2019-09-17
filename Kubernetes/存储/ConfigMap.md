# configmap

[TOC]

- ConfigMap是存储通用的配置变量的，类似于配置文件，使用户可以将分布式系统中用于不同模块的环境变量统一到一个对象中管理；
- 它与配置文件的区别在于它是存在集群的“环境”中的，并且支持K8S集群中所有通用的操作调用方式。


***pod使用ConfigMap，通常用于：设置环境变量的值、设置命令行参数、创建配置文件***


## 创建configmap
可使用kubectl create configmap从文件，目录或key-value字符串创建，也可通过kubectl create -f file创建

### 命令行
```kubectl create configmap lykops-config --from-file=db_config_file=database.conf --from-file=ver.conf --from-literal=username=test --from-literal=hostname=localhost```

- –from-file表示来自文件，直接把文件内容写入configmap中，可以为目录也可以为文件，如果是文件的话，可以使用db_config_file=database.conf来修改key值 
- –from-literal表示使用键值对配置
