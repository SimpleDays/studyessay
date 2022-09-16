---
title: helm学习使用
date: 2022-09-16 10:00:00
tags: kubernetes
author: 小笼包
categories: Kubernetes
---

> 作者: 小笼包  
> 2022-09-16 晴

## 使用helm可以解决什么问题？

- 使用helm可以把这些yaml作为一个整体管理

- 实现yaml高效复用

- 使用helm应用级别的版本管理

## helm是什么？

helm是一个kubernetes的包管理工具，就像Linux下的包管理器一样，如 yum/apt等，可以很方便的将之前打包好的yaml文件部署到kubernetes上。

## 3个重要概念

- helm：一个命令行客户端工具，主要用于kubernetes应用chart的创建、打包、发布和管理。

- chart：应用的描述，一系列用于描述k8s资源相关文件的集合

- release：基于chart的部署实体，一个chart被helm运行后将会生成对应的一个release；将在k8s中创建出真实运行的资源对象。

## helm在2019年发布v3版本，和之前版本有哪些变化？

- v3版本删除了Tiller。

- release可以在不同的命名空间中重用，此前版本不支持。

- 将chart推送至docker仓库中去

### helm的架构区别

![如图](./diagram/k8s-helm.dio)

## helm的安装

### 1、打开helm官网下载二进制包

官网地址：<https://helm.sh/zh/docs/intro/install/>

可以选择自己下载二进制包安装，也可以通过命令脚本安装。具体参考官网提供的方式。

建议直接翻墙下载二进制包比较快，用脚本下载太折磨人。

二进制下载下来，解压 tar zxf xxxx

拷贝helm文件至 **/usr/bin/**下

```shell
cp helm /usr/bin/
```

### 2、配置helm仓库

helm repo add 仓库名称 仓库地址

例如：

```shell
helm repo add stable http://mirror.azure.cn/kubernetes/charts

helm repo add aliyun https://kubernetes.oss-cn-hangzhou.aliyuncs.com/charts

helm repo update

# 查看仓库配置
helm repo list

# 查询仓库名叫stable
helm search repo stable

# 移除仓库
helm repo remove aliyun
```

## 使用helm部署应用

安装helm界面weave

### 1、使用命令搜索应用

```shell
[root@master01 linux-amd64]# helm search repo weave
NAME                    CHART VERSION   APP VERSION     DESCRIPTION                                       
aliyun/weave-cloud      0.1.2                           Weave Cloud is a add-on to Kubernetes which pro...
aliyun/weave-scope      0.9.2           1.6.5           A Helm chart for the Weave Scope cluster visual...
stable/weave-cloud      0.3.9           1.4.0           DEPRECATED - Weave Cloud is a add-on to Kuberne...
stable/weave-scope      1.1.12          1.12.0          DEPRECATED - A Helm chart for the Weave Scope c...
[root@master01 linux-amd64]#
```

### 2、根据搜索的应用选择安装

```shell
[root@master01 linux-amd64]# helm install ui stable/weave-scope
WARNING: This chart is deprecated
NAME: ui
LAST DEPLOYED: Fri Sep 16 11:51:50 2022
NAMESPACE: default
STATUS: deployed
REVISION: 1
NOTES:
You should now be able to access the Scope frontend in your web browser, by
using kubectl port-forward:

kubectl -n default port-forward $(kubectl -n default get endpoints \
ui-weave-scope -o jsonpath='{.subsets[0].addresses[0].targetRef.name}') 8080:4040

then browsing to http://localhost:8080/.
For more details on using Weave Scope, see the Weave Scope documentation:

https://www.weave.works/docs/scope/latest/introducing/
[root@master01 linux-amd64]# helm list
NAME    NAMESPACE       REVISION        UPDATED                                 STATUS          CHART                   APP VERSION
ui      default         1               2022-09-16 11:51:50.100234491 +0800 CST deployed        weave-scope-1.1.12      1.12.0     
[root@master01 linux-amd64]# helm status
Error: "helm status" requires 1 argument

Usage:  helm status RELEASE_NAME [flags]
[root@master01 linux-amd64]# helm status
Error: "helm status" requires 1 argument

Usage:  helm status RELEASE_NAME [flags]
[root@master01 linux-amd64]# helm status ui
NAME: ui
LAST DEPLOYED: Fri Sep 16 11:51:50 2022
NAMESPACE: default
STATUS: deployed
REVISION: 1
NOTES:
You should now be able to access the Scope frontend in your web browser, by
using kubectl port-forward:

kubectl -n default port-forward $(kubectl -n default get endpoints \
ui-weave-scope -o jsonpath='{.subsets[0].addresses[0].targetRef.name}') 8080:4040

then browsing to http://localhost:8080/.
For more details on using Weave Scope, see the Weave Scope documentation:

https://www.weave.works/docs/scope/latest/introducing/
[root@master01 linux-amd64]# 
```
