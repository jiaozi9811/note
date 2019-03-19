# kubeadm

[TOC]

## 准备环境

```shell
关闭防火墙：
# systemctl stop firewalld
# systemctl disable firewalld

关闭selinux：
# sed -i 's/enforcing/disabled/' /etc/selinux/config 
# setenforce 0

关闭swap：
# swapoff -a  # 临时
# vim /etc/fstab  # 永久
# sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab

添加主机名与IP对应关系：
# cat /etc/hosts
192.168.0.11 k8s-master
192.168.0.12 k8s-node1
192.168.0.13 k8s-node2

同步时间：
# yum install ntpdate -y
# ntpdate  ntp.api.bz

wget https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

vim /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg


yum install -y docker-ce

vim /usr/lib/systemd/system/docker.service
Environment="HTTPS_PROXY=http://www.ik8s.io:10080"
Environment="NO_PROXY=127.0.0.0/8,127.0.0.0/16"

systemctl daemon-reload
systemctl enable docker && systemctl start docker
docker info

安装kubeadm，kubelet和kubectl
rpm --import https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg
rpm --import https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
# yum install -y kubelet kubeadm kubectl --disableexcludes=kubernetes
# systemctl enable kubelet && systemctl start kubelet

echo 1 > /proc/sys/net/bridge/bridge-nf-call-iptables

vim /etc/sysconfig/kubelet
KUBELET_EXTRA_ARGS="--fail-swap-on=false"

初始化Master
#kubeadm init --kubernetes-version=1.13.4 --pod-network-cidr=10.244.0.0/16 --service-cidr=10.96.0.0/12  --ignore-preflight-errors=Swap 

kubeadm init --image-repository=registry.aliyuncs.com/google_containers --kubernetes-version=1.13.4 --pod-network-cidr=10.244.0.0/16 --service-cidr=10.96.0.0/12 --ignore-preflight-errors=Swap --apiserver-advertise-address=10.0.0.10

```

https://kubethink.com/posts/create-an-on-premises-cluster-in-china-v1-12/

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

kubeadm join 10.0.4.10:6443 --token xcqp4o.70xcijxgogyoysrq --discovery-token-ca-cert-hash sha256:395ce644da3f730873df2adb4568172235b395f57006659bc1d2a3873cfe3193


部署flannel
docker pull registry.cn-hangzhou.aliyuncs.com/gaven_k8s/flannel:v0.11.0-amd64

docker tag registry.cn-hangzhou.aliyuncs.com/gaven_k8s/flannel:v0.11.0-amd64 quay.io/coreos/flannel:v0.11.0-amd64

docker rmi registry.cn-hangzhou.aliyuncs.com/gaven_k8s/flannel:v0.11.0-amd64

kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml



#!/bin/bash
images=(
kube-controller-manager:v1.13.4
kube-apiserver:v1.13.4
kube-scheduler:v1.13.4
kube-proxy:v1.13.4
pause:3.1
etcd:3.2.24
coredns:1.2.6
)

for imageName in ${images[@]} ; do
    docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName
    docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName k8s.gcr.io/$imageName
done



### 版本信息
K8S_VERSION=v1.11.0
ETCD_VERSION=3.1.12
DASHBOARD_VERSION=v1.8.3
FLANNEL_VERSION=v0.10.0-amd64
DNS_VERSION=1.14.8
PAUSE_VERSION=3.1
## 基本组件
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver-amd64:$K8S_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager-amd64:$K8S_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler-amd64:$K8S_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy-amd64:$K8S_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/etcd-amd64:$ETCD_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/pause-amd64:$PAUSE_VERSION
### 网络
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-sidecar-amd64:$DNS_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-kube-dns-amd64:$DNS_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-dnsmasq-nanny-amd64:$DNS_VERSION
docker pull quay.io/coreos/flannel:$FLANNEL_VERSION
### 前端
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kubernetes-dashboard-amd64:$DASHBOARD_VERSION

## 修改tag
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver-amd64:$K8S_VERSION k8s.gcr.io/kube-apiserver-amd64:$K8S_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager-amd64:$K8S_VERSION k8s.gcr.io/kube-controller-manager-amd64:$K8S_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler-amd64:$K8S_VERSION k8s.gcr.io/kube-scheduler-amd64:$K8S_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy-amd64:$K8S_VERSION k8s.gcr.io/kube-proxy-amd64:$K8S_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/etcd-amd64:$ETCD_VERSION k8s.gcr.io/etcd-amd64:$ETCD_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/pause-amd64:$PAUSE_VERSION k8s.gcr.io/pause-amd64:$PAUSE_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-sidecar-amd64:$DNS_VERSION k8s.gcr.io/k8s-dns-sidecar-amd64:$DNS_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-kube-dns-amd64:$DNS_VERSION k8s.gcr.io/k8s-dns-kube-dns-amd64:$DNS_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-dnsmasq-nanny-amd64:$DNS_VERSION k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64:$DNS_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kubernetes-dashboard-amd64:$DASHBOARD_VERSION k8s.gcr.io/kubernetes-dashboard-amd64:$DASHBOARD_VERSION

## 删除镜像
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver-amd64:$K8S_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager-amd64:$K8S_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler-amd64:$K8S_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy-amd64:$K8S_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/etcd-amd64:$ETCD_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/pause-amd64:$PAUSE_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-sidecar-amd64:$DNS_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-kube-dns-amd64:$DNS_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-dnsmasq-nanny-amd64:$DNS_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/kubernetes-dashboard-amd64:$DASHBOARD_VERSION



部署节点

docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.1
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.1 k8s.gcr.io/pause:3.1
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.1


docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy:v1.13.4
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy:v1.13.4 k8s.gcr.io/kube-proxy:v1.13.4
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy:v1.13.4


kubeadm join 192.168.92.56:6443 --token 67kq55.8hxoga556caxty7s --discovery-token-ca-cert-hash sha256:7d50e704bbfe69661e37c5f3ad13b1b88032b6b2b703ebd4899e259477b5be69

kubectl run nginx-deploy --image=naginx:1.14-alpine --port=80 --replicas-1 --dry-run=true