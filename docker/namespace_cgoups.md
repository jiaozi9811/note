namespace类型		系统调用参数
mount namespace		CLONE_NEWNS
uts namespace		CLONE_NEWUTS
ipc namespace		CLONE_NEWIPC
pid namespace		CLONE_NEWPID
network namespace	CLONE_NEWNET
user namespace		CLONE_NEWUSER

nts namespace
	主要用于隔离nodename和domainname两个系统标识。uts namespae中，每个namespace允许有自己的hostname

ipc namespace
	用于隔离system v ipc和posix message queues。每个ipc namespace都有自己的system v ipc和posix message queues

pid namespace
	用于隔离进程id

mount namespace
	用于隔离各个进程看到的挂载点视图		

user namespace
	主要隔离用户的用户组id

network namespace
	用于隔离网络设备，ip地址端口等网络栈namespace


linux Cgroups(control groups)
	提供了对一组进程及将来子进程的资源限制，控制和统计能力

cgoups组件
	cgroup 		对进程分组管理的一种机制。一个cgroup包含一组进程，并可以在这个cgroup上增加linux subsystem的各种参数配置，将一组进程和一组subsystem的系统参数关联起来
	subsystem 	一组资源控制的模块，包括
		blkio	设置对块设备输入输出的访问控制
		cpu 	设置cgroup中进程的cpu被调度的策略
		cpuacct 统计cgroup中进程cpu的占用
		cpuset 	在多核机器上设置cgroup中进程可以使用的cpu和内存
		devices 控制cgroup中进程对设备的访问
		freezer 用于挂起和恢复cgroup中的进程
		memory	控制cgroup中进程的内存占用
		net_cls 用于将cgroup中进程产生的网络包分类
		net_prio 设置cgoup中进程产生的网络流量的优先级
		ns 		使cgroup中的进程在新的namespace中fork新进程时，创建一个新cgroup
	hierarchy	把一组cgroup串成一个树状结构，从而做到继承
