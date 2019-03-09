# vagrant

[TOC]

Vagrant是为了方便的实现虚拟化环境而设计的，使用Ruby开发，基于VirtualBox等虚拟机管理软件的接口，提供了一个可配置、轻量级的便携式虚拟开发环境。使用Vagrant可以很方便的就建立起来一个虚拟环境，而且可以模拟多台虚拟机，这样我们平时还可以在开发机模拟分布式系统。

Vagrant还会创建一些共享文件夹，用来给你在主机和虚拟机之间共享代码用。这样就使得我们可以在主机上写程序，然后在虚拟机中运行。如此一来团队之间就可以共享相同的开发环境，就不会再出现类似“只有你的环境才会出现的bug”这样的事情。

团队新员工加入，常常会遇到花一天甚至更多时间来从头搭建完整的开发环境，而有了Vagrant，只需要直接将已经打包好的package（里面包括开发工具，代码库，配置好的服务器等）拿过来就可以工作了，这对于提升工作效率非常有帮助。

Vagrant不仅可以用来作为个人的虚拟开发环境工具，而且特别适合团队使用，它使得我们虚拟化环境变得如此的简单，只要一个简单的命令就可以开启虚拟之路。

## VAGRANT 中文文档

##安装配置

Vagrant只是一个让你可以方便设置你想要的虚拟机的便携式工具，它底层支持VirtualBox、VMware甚至AWS作为虚拟机系统

安装VirtualBox
安装Vagrant
http://www.vagrantup.com/downloads.html

box
http://www.vagrantbox.es/

https://app.vagrantup.com/boxes/search

## 常用命令

vagrant global-status

   vagrant box add [vagrant_name] [vagrant_url]  --name [name]

    vagrant init [vagrant_name]

    vagrant up          //启动vagrant

    vagrant ssh     127.0.0.1:2222 用户密码：vagrant@vagrant

    vagrant destroy     //销毁vagrant

     vagrant halt          关机





vagrant reload                    重新启动虚拟机，主要用于重新载入配置文件

vagrant suspend                 挂起当前的虚拟机

vagrant resume                   恢复前面被挂起的状态

vagrant status                    获取当前虚拟机的状态

vagrant box list                   显示当前已经添加的box列表

vagrant box remove



vagrant package                    打包命令，可以把当前的运行的虚拟机环境进行打包

vagrant plugin                    用于安装卸载插件

vagrant provision

vagrant ssh-config               输出用于ssh连接的一些信息

## box命令

add
list
outdated
prune
remove
repackage
update

## 修改 Vagrant Box 加载目录

默认加载目录在 ~/.vagrant.d/

环境变量是VAGRANT_HOME

### 永久设置环境变量

永久设置用户的环境变量
          setx VAGRANT_HOME "/your/path"
永久设置系统的环境变量
          setx VAGRANT_HOME "/your/path" /M
##临时设置环境变量 这样的方式只是临时性地将环境变量写入session,每次要打开新终端执行vagrant相关命令如重启关闭销毁等命令之前都需要敲设置一遍

          set VAGRANT_HOME="/your/path"
 
## root登录

将一下配置加到Vagrantfile文件

config.ssh.username = 'root'
config.ssh.password = 'vagrant'
config.ssh.insert_key = 'true'


## 共享目录

  config.vm.synced_folder "D:/python", "/code", create: true, owner: "root", group: "root", mount_options: ["dmode=755","fmode=644"]#, type: "rsync"  
  config.vm.synced_folder ".","/vagrant",disabled:true
