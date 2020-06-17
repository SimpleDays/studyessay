---
title: Kubernetes删除状态为Terminating的命名空间
date: 2020-06-17 14:27:00
tags: kubernetes
author: 小笼包
categories: Kubernetes
---

## Kubernetes删除状态为Terminating的命名空间

> 作者: 小笼包  
> 2020-06-17 有雨

## 前言

今天公司开发环境Kubernetes集群使用起来很卡，于是我们对Kubernetes进行清理，主要删除无用的命名空间和重复多余的部署项目。当我通过程序删除无用的Kubernetes命名空间之后，发现命名空间的状态一直处于**Terminating**的状态，一直没有删除掉，虽然找了运维协助帮忙，但是这个问题不常遇到，在尝试了master节点逐个重启，和降低master机器负载等方式，该问题依旧存在，运维也没办法解决了，网上资料排查也没一个有效解决这类问题的标准方式，最后我只能自己尝试寻找解决方案，最终根据几个网站资料，找到了可以解决这问题的方法。  

<!-- more -->

## 我遇到的问题记录

我是通过程序调用K8s的API进行删除命名空间的，不过这里我用命令来进行删除如下：

``` shell
[root@sy-suz-srv91 ~]# kubectl delete ns f5
[root@sy-suz-srv91 ~]# kubectl delete ns lonely
[root@sy-suz-srv91 ~]# kubectl delete ns monitor2
```

然后我们查看当前k8s下的命名空间时候发现被我删除的命名空间呈现如下状态：  

``` shell
[root@sy-suz-srv91 ~]# kubectl get ns
NAME              STATUS        AGE
default           Active        246d
erp100            Active        245d
f5                Terminating   97d
kube-node-lease   Active        246d
kube-public       Active        246d
kube-system       Active        246d
lonely            Terminating   230d
mall              Active        242d
monitor           Active        245d
monitor1          Active        69d
monitor2          Terminating   79d
monitoring        Active        183d
```

一开始认为这些命名空间下的资源还没彻底删除完，需要等待一段时间，但是当我等了很久，然后通过K8s的dashboard确认了这些命名空间下的资源已经被释放殆尽之后，再次确认时候发现这几个被我删除的命名空间依旧如上面命令执行展示一样，然后这时候我再次执行删除命名空间时候，k8s被我报了如下的错误：  

``` shell
[root@sy-suz-srv91 ~]# kubectl delete ns f5
Error from server (Conflict): Operation cannot be fulfilled on namespaces "f5": The system is ensuring all content is removed from this namespace.  Upon completion, this namespace will automatically be purged by the system.
```

错误大概意思就是：

>服务器错误（冲突）：无法在名称空间“f5”上实现操作：系统正在确保从此名称空间中删除所有内容。 完成后，系统将自动清除此命名空间。

也就是让我等待系统等资源清理完毕之后才会自动删除，可是我已经确认已经删除完了，这时候我又查了资料尝试使用强制删除这个命名空间：  

``` shell
[root@sy-suz-srv91 ~]# kubectl delete ns f5 --force --grace-period=0
warning: Immediate deletion does not wait for confirmation that the running resource has been terminated. The resource may continue to run on the cluster indefinitely.
Error from server (Conflict): Operation cannot be fulfilled on namespaces "f5": The system is ensuring all content is removed from this namespace.  Upon completion, this namespace will automatically be purged by the system.
```

错误信息大致意思：  
> 警告： 立即删除不等待确认正在运行的资源已被终止。 资源可能继续无限期地运行在集群上。 
> 服务器错误（冲突）：无法在名称空间“f5”上实现操作：系统正在确保从此名称空间中删除所有内容。 完成后，系统将自动清除此命名空间。

## 解决方式记录

发现似乎正常的操作都无法删除这个命名空间了，咨询运维，运维除了重启master似乎也没有其他办法，但是重启也无济于事，根据我查询资料，可以通过如下操作删除这样僵死的k8s命名空间。  

### 1、获得需要删除的k8s命名空间的描述，并且以json的格式保存至当前路径下

``` shell
[root@sy-suz-srv91 ~]# kubectl get namespace f5 -o json > f5.json
```

### 2、打开**f5.json**文件，删除其中的**spec**下字段**finalizers**的内容

``` shell
[root@sy-suz-srv91 ~]# vim f5.json
{
    "apiVersion": "v1",
    "kind": "Namespace",
    "metadata": {
        "creationTimestamp": "2020-03-11T08:32:40Z",
        "deletionTimestamp": "2020-06-17T01:30:12Z",
        "name": "f5",
        "resourceVersion": "80204490",
        "selfLink": "/api/v1/namespaces/f5",
        "uid": "b965d164-05f1-4e99-90cb-be03ba4257c6"
    },
    "spec": {
        "finalizers": [
            "kubernetes"  #把这行删除，并且保存即可
        ]
    },
    "status": {
        "phase": "Terminating"
    }
}
```

### 3、一般K8s集群都是带有认证的，要先克隆一个新的代理会话，如下操作

``` shell
[root@sy-suz-srv91 ~]# kubectl proxy --proxy=8081
Starting to serve on 127.0.0.1:8081
```

### 4、上述操作会让当前会话挂起，此时重新打开一个窗口，然后通过这个临时代理端口去删除命名空间

``` shell
[root@sy-suz-srv91 ~]# curl -k -H "Content-Type: application/json" -X PUT --data-binary @f5.json http://127.0.0.1:8081/api/v1/namespaces/f5/finalize
{
    "apiVersion": "v1",
    "kind": "Namespace",
    "metadata": {
        "creationTimestamp": "2020-03-11T08:32:40Z",
        "deletionTimestamp": "2020-06-17T01:30:12Z",
        "name": "f5",
        "resourceVersion": "80204490",
        "selfLink": "/api/v1/namespaces/f5",
        "uid": "b965d164-05f1-4e99-90cb-be03ba4257c6"
    },
    "spec": {
        "finalizers": [

        ]
    },
    "status": {
        "phase": "Terminating"
    }
}
```

### 5、这时候再查看命名空间时候发现这个僵死的命名空间已经被删除了

``` shell
[root@sy-suz-srv91 ~]# kubectl get ns
NAME              STATUS   AGE
default           Active   246d
erp100            Active   245d
kube-node-lease   Active   246d
kube-public       Active   246d
kube-system       Active   246d
mall              Active   242d
monitor           Active   245d
monitor1          Active   70d
monitoring        Active   183d
```

通过以上的方式可以有效删除K8s下僵死的命名空间。

参考文档：  
1、<https://copyfuture.com/blogs-details/20191021202129836irvun8sxdavssr3>  
2、<https://www.cnblogs.com/wangxu01/articles/11792908.html>  
3、<https://blog.csdn.net/tongzidane/article/details/88988542>  
4、<https://blog.csdn.net/weixin_44774358/article/details/97134277>  
