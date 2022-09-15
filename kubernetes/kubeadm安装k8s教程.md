---
title: kubeadm安装k8s教程
date: 2022-09-14 01:09:00
tags: kubernetes
author: 小笼包
categories: Kubernetes
---

> 作者: 小笼包  
> 2022-09-14 大雨  

1、初始化centos虚拟化服务器，介绍下使用服务器是什么内核版本什么系统版本的，并且说下网段和硬件配置

2、通过kubeadm安装配置过程
参考地址 <https://gitee.com/moxi159753/LearningNotes/tree/master/K8S/3_%E4%BD%BF%E7%94%A8kubeadm%E6%96%B9%E5%BC%8F%E6%90%AD%E5%BB%BAK8S%E9%9B%86%E7%BE%A4>

3、安装dashboard说明
参考地址 <https://github.com/sskcal/kubernetes/tree/main/dashboard>

4、生成yaml模版，通过create命令
kubectl create deployment web --image-nginx -o yaml --dry-run

5、更具已有的deployment的信息快速生成yaml文件并到处成yaml类型的文件
kubectl get deploy nginx -o=yaml --export > nginx-demo.yaml
