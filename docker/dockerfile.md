# dockerfile

## Comment 注释

## FROM

用于为镜像文件构建过程指定基准镜像，后续的指令运行于此基准镜像所提供的运行环境
FROM　<repository>[:<tag>]
    repository  基础镜像的名称
    tag         镜像的标签

## MAINTAINER

用于让dockerfile制作者提供本人的详细信息
MAINTANIER <authtor's detail>

## LABEL 标签

为镜像添加元数据(etadata)

# COPY

从docker主机复杂文件到创建的镜像文件中
    COPY "<src>" "<dest>"

    如果src是目录，src中的文件和目录会递归复制，但src目录自身不会被复制
    如果有多个src，dest必须是一个目录，且必须以/结尾
    如果dest事先不存在，会被递归的创建

## ADD
    类似COPY，但支持使用tar文件和url路径
        ADD "<src>" "<dest>"

        如果src是url且dest不以/结尾，则src指定的文件将被下载，并直接被创建为dest
        如果src是一个本地系统上的tar文件，将被展开为一个目录，类似于tar -x

## WORKDIR

为dockerfile中的指令指定工作目录

## VOLUME

在image中创建一个挂载点，挂载docker host上的卷或其他容器中的卷
如果挂载点目录下此前在文件中存在，会在卷挂载后将此前的所有文件复制到新挂载点卷上
            VOLUME "mountpoint"


## EXPOSE

用于为容器打开指定要监听的端口以实现与外部通信
    EXPOSE <port>[/<protocol>] <port>[/<protocol>]       


## ENV

## RUN  

 RUN "<executable>" "<param1>" "<param2>"


## CMD

docker run命令传入的命令参数会覆盖CMD指令的内容

## ENTRYPOINT

类似CMD，用于为容器指定默认的运行程序，从而使容器像是一个单独的可执行程序
可存在多个ENTRYPOINT，但只有最后一个生效

## USER

指定运行RUN CMD ENTRYPOINT的用户

## HEALTHCHECK

## SHELL

指定运行的shell shell/csh/ksh

## STOPSIGNAL

## ARG

## ONBUILD

    定义一个触发器
