---
title: 初识Envoy
date: 2020-05-20 21:46:00
tags: Envoy
author: 小笼包
categories: servermesh
---

## 初识Envoy

> 作者: 小笼包  
> 2020-05-20 晴

### Envoy是什么

Envoy是一个七层负载均衡的代理，专门为现代大型服务方向架构设计。  

简单来说Envoy是一种"服务网格"，为各种语言的应用程序提供通用的基础类库，例如：  

服务发现、负载均衡、服务限流、断路器、可观测性（统计）、日志、链路追踪 等功能。

<!-- more -->

### Envoy理念基于什么

Envoy设计初衷理念 ： The network should be transparent to applications. When network and application problems do occur, it should be easy to determine the source of the problem.  （网络对于应用应该是透明的，当网络和应用出现问题时，应该很容易找到问题的根源）。

### Envoy的一些故事分享

### Envoy现有基础功能详解

### Envoy的线程模型是如何的

### Envoy可以做些什么

### 上手试连，通过Envoy实现一个百度简单代理