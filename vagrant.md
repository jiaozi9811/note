# vagrant

[TOC]

Vagrant是为了方便的实现虚拟化环境而设计的，使用Ruby开发，基于VirtualBox等虚拟机管理软件的接口，提供了一个可配置、轻量级的便携式虚拟开发环境。使用Vagrant可以很方便的就建立起来一个虚拟环境，而且可以模拟多台虚拟机，这样我们平时还可以在开发机模拟分布式系统。

Vagrant还会创建一些共享文件夹，用来给你在主机和虚拟机之间共享代码用。这样就使得我们可以在主机上写程序，然后在虚拟机中运行。如此一来团队之间就可以共享相同的开发环境，就不会再出现类似“只有你的环境才会出现的bug”这样的事情。

团队新员工加入，常常会遇到花一天甚至更多时间来从头搭建完整的开发环境，而有了Vagrant，只需要直接将已经打包好的package（里面包括开发工具，代码库，配置好的服务器等）拿过来就可以工作了，这对于提升工作效率非常有帮助。

Vagrant不仅可以用来作为个人的虚拟开发环境工具，而且特别适合团队使用，它使得我们虚拟化环境变得如此的简单，只要一个简单的命令就可以开启虚拟之路。

https://www.wanglibing.com/vagrant/

## VAGRANT 中文文档

##安装配置

Vagrant只是一个让你可以方便设置你想要的虚拟机的便携式工具，它底层支持VirtualBox、VMware甚至AWS作为虚拟机系统

安装VirtualBox
安装Vagrant
http://www.vagrantup.com/downloads.html

box
http://www.vagrantbox.es/

https://app.vagrantup.com/boxes/search

https://cloud.centos.org

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

在Vagrant中添加box时，加载目录默认在 ~/.vagrant.d/，具体的目录结构是C:\Users\Your Username\.vagrant.d</em>

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

 网络

### 端口映射(Forwarded port)

`config.vm.forwarded_port 80, 8080`

### 私有网络(Private network)

`config.vm.network "private_network", ip: "192.168.21.4"`
`config.vm.network "private_network", type: "dhcp",auto_config: false`

Disable Auto-Configuration
 want to manually configure the network interface yourself, you can disable Vagrant's auto-configure feature by specifying auto_config
    auto_config: false

### 公有网络(Public network)

`config.vm.network "public_network", ip: "10.0.0.10"`


config.vm.box = "centos/7" //配置box
config.vm.hostname = “node1"配置hostname

config.vm.provider "virtualbox" do |vb|
  vb.name = "centos7" //配置虚拟机名称
  vb.cpus = 2 配置CPU个数
  vb.memory = "1024"配置内存单位：MB
end


## 端口转发

onfig.vm.network "forwarded_port", guest: 80, host: 8080
config.vm.network "forwarded_port", guest: 81, host: 8081

挂载/oracle/virtulbox/VBoxGuestAdditions.iso
mount挂载上
bash VBoxLinuxAdditions.run

验证”virtualbox的文件系统模块“是否加载了,lsmod=list modules
lsmod | grep vboxsf
 
如果查找不到，使用加载模块，接着再验证，每次reboot都要做这个操作
modprobe -a vboxsf

在安装增强功能的时候出现了，kernel headers not found for target kernel的错误。特记下我的解决方案。

1.update kernel

yum update kernel -y
2.Install the kernel-headers, kernel-devel and other required packages

yum install kernel-headers kernel-devel gcc make -y

## 导出box

$ vagrant package --base <虚拟机名称> --output BoxName.box

Vagrant.configure("2") do |config|
    config.vm.define :web do |web|
        web.vm.provider "virtualbox" do |v|
            v.customize ["modifyvm", :id, "--name", "web", "--memory", "1024"]
        end
        web.vm.box = "centos/7"
        web.vm.hostname = "node1"
        web.vm.network :private_network, ip : "11.11.11.1"
    end

    config.vm.define :db do |db|
        db.vm.provider "virtualbox" do |v|
            v.customize ["modifyvm", :id, "--name", "db", "--memory", "1024"]
        end
        db.vm.box = "centos/7"
        db.vm.hostname = "node2"
        db.vm.network :private_network, ip : "11.11.11.2"
    end
end


Vagrant.configure("2") do |config|
    config.vm.define :master do |master|
        web.vm.provider "virtualbox" do |v|
            v.customize ["modifyvm", :id, "--name", "master", "--memory", "512"]
        end
        web.vm.box = "centos"
        web.vm.hostname = "master"
        web.vm.network :private_network, ip : "10.0.0.10"
    end

    config.vm.define :node1 do |node1|
        db.vm.provider "virtualbox" do |v|
            v.customize ["modifyvm", :id, "--name", "node1", "--memory", "512"]
        end
        db.vm.box = "centos"
        db.vm.hostname = "node1"
        db.vm.network :private_network, ip : "10.0.0.11"
    end
end