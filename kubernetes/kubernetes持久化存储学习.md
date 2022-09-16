---
title: kubernetes持久化存储学习
date: 2022-09-16 10:00:00
tags: kubernetes
author: 小笼包
categories: Kubernetes
---

> 作者: 小笼包  
> 2022-09-16 晴

## nfs数据持久化存储

首先我选择一台空服务器添加nfs服务，作为挂在远程存储服务器地址：192.168.20.132。

然后需要挂在的k8s节点也需要同样的装上nfs服务。

### 1、安装nfs

```shell
[root@localhost dongliang]# yum install -y nfs-utils
```

### 2、设置挂载路径

```shell
# 添加nfs读写权限
# /data/nfs *(rw,no_root_squash)
[root@localhost etc]# vim /etc/exports
[root@localhost etc]# cat exports
/data/nfs *(rw,no_root_squash)

# 创建nfs挂载路径 /data/nfs
[root@localhost etc]# cd /
[root@localhost /]# mkdir data
[root@localhost /]# cd data
[root@localhost data]# mkdir nfs
[root@localhost data]# ls
nfs
[root@localhost data]# cd nfs
[root@localhost nfs]# pwd
/data/nfs
```

### 3、启动nfs服务

```shell
[root@localhost ~]# systemctl start nfs
[root@localhost ~]# ps -ef | grep "nfs"
root      1602     2  0 15:49 ?        00:00:00 [nfsd]
root      1603     2  0 15:49 ?        00:00:00 [nfsd]
root      1604     2  0 15:49 ?        00:00:00 [nfsd]
root      1605     2  0 15:49 ?        00:00:00 [nfsd]
root      1606     2  0 15:49 ?        00:00:00 [nfsd]
root      1607     2  0 15:49 ?        00:00:00 [nfsd]
root      1608     2  0 15:49 ?        00:00:00 [nfsd]
root      1609     2  0 15:49 ?        00:00:00 [nfsd]
root      1622  1304  0 15:49 pts/0    00:00:00 grep --color=auto nfs
[root@localhost ~]# 
# 开机启动nfs服务
[root@localhost ~]# systemctl enable nfs
```

### 4、配置应用挂在nfs存储

创建一个 nfs-nginx.yaml文件来测试

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-nfs-demo1
spec:
  replicas: 1
  selector:
    matchLabels:
       app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
        volumeMounts:
        - name: wwwroot
          mountPath: /usr/share/nginx/html
        ports:
        - containerPort: 80
      volumes:
        - name: wwwroot
          nfs:
            server: 192.168.20.132
            path: /data/nfs
```

### 遇到的问题

在nfs服务器配置完毕之后我遇到了 **mount.nfs: No route to host** 无法访问nfs服务的问题，这主要是nfs服务器上的防火墙还开着导致的，所以在nfs配置前我们需要对nfs服务器做一些初始化设置，具体操作如下：

```shell
# 关闭防火墙
systemctl stop firewalld
systemctl disable firewalld

# 关闭selinux
# 永久关闭
sed -i 's/enforcing/disabled/' /etc/selinux/config  
# 临时关闭
setenforce 0  

# 关闭swap
# 临时
swapoff -a 
# 永久关闭
sed -ri 's/.*swap.*/#&/' /etc/fstab

# 设置服务器名为 nfs-server
hostnamectl set-hostname nfs-server

# 将桥接的IPv4流量传递到iptables的链
cat > /etc/sysctl.d/k8s.conf << EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
# 生效
sysctl --system  
```

然后重启下服务器，等nfs启动起来，重新部署服务就能正常挂载上去了。

## PV和PVC

### 1、什么是PV

PV是持久化存储， 对存储资源的抽象，对外提供可以调用的地方。（生产者）

### 2、什么是PVC

 PVC用于调用， 不需要关心内部如何实现。（消费者）

### PV与PVC的实现流程

![如图](./diagram/k8s-pv%26pvc.dio)

### 简单实现一个基本的PV和PVC存储

创建一个 pvc-nginx.yaml文件来测试

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-pvc-demo1
spec:
  replicas: 1
  selector:
    matchLabels:
       app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
        volumeMounts:
        - name: wwwroot
          mountPath: /usr/share/nginx/html
        ports:
        - containerPort: 80
      volumes:
        - name: wwwroot
          persistentVolumeClaim:
            claimName: my-pvc

---

apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-pvc
spec:
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 5Gi
```

定义一个pv.yaml

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: my-pv
spec:
  capacity:
    storage: 5Gi
  accessModes:
    - ReadWriteMany
  nfs:
    path: /data/nfs
    server: 192.168.20.132
```

通过对外暴露NodePort端口，修改nfs下html的内容来确定是否挂载成功

```shell
[root@localhost ~]# kubectl expose deployment nginx-pvc-demo1 --port=80 --target-port=80 --type=NodePort
service/nginx-pvc-demo1 exposed
[root@master01 dongliang]#
```
