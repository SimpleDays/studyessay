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

1、**Matt Klein（马特-克莱因） ：**  现Envoy开源社区负责人，Lyft公司工程师。  

> [Matt Klein的推特](https://twitter.com/mattklein123)

2、Envoy在开源成立前，由Lyft公司研发与使用大概长达1.5年时间，主要想解决Lyft公司零星  
的网络和服务调用失败的持续问题，以至于大多数开发人员都不敢在关键路径上进行大量的服务  
调用。难以理解问题发生在何处，最终Lyft公司还通过Envoy替代了基础架构中对Amazon的ELB  
使用。  

> [参考Matt Klein写的文章(Announcing Envoy: C++ L7 proxy and communication bus)](https://eng.lyft.com/announcing-envoy-c-l7-proxy-and-communication-bus-92520b6c8191)  

3、当Envoy正式成立开源之后，大量Google的开发人员参与为Envoy开源推广添砖加瓦，不久Lyft和Google  
联合推出了istio，很快也受到很多云原生大厂的喜爱，陆续的其他开源社区大佬都加入到了Envoy开源社区，国内蚂蚁金服的云原生布道者Jimmy Song为Envoy文档提供了中文版文档，在这里感谢这些大佬，使得我们接下来了解和学习Envoy、ServerMesh有了一个良好的环境。    

4、Matt Klein 并没有与微服物布道者(克里斯-理查森)一样开设公司专职于Envoy推广，反而崇尚开源精神，完全交给开源社区，并与开源社区一起努力推广Envoy，本人也是非常倾佩这样行为（确实Envoy的成功可以为他增加很多用他的词汇就是更多的货币收入），相信开源社区。  

> [详见（Optimizing impact: why I will not start an Envoy platform company）](https://medium.com/@mattklein123/optimizing-impact-why-i-will-not-start-an-envoy-platform-company-8904286658cb)

### Envoy现有基础功能简单介绍

1、**进程外的架构模式 :** Envoy是一个自包含的流程，旨在与每个应用程序服务器一起运行。所有的Envoy都构成一个透明的通信网，每个应用程序在其中都与localhost之间发送和接收消息，并且不知道网络拓扑。与传统库方法进行服务到服务的通信相比，进程外架构具有两个实质性的好处：  

- Envoy可与任何应用程序语言一起使用。单个Envoy部署可以在Java，C ++，Go，PHP，Python等之间形成网格。面向服务的体系结构使用多种应用程序框架和语言正变得越来越普遍。特使透明地弥合了差距。  

- 任何使用大型面向服务的体系结构的人都知道，部署库升级可能非常痛苦。Envoy可以透明地在整个基础架构中快速部署和升级。  

2、**现代C ++ 11代码库 :** Envoy用C ++ 11编写。选择本机代码是因为我们认为，Envoy之类的体系结构组件应尽可能避免使用。由于共享云环境中的部署以及使用生产力很高但性能不是特别好的语言（例如PHP，Python，Ruby，Scala等），现代应用程序开发人员已经处理了难以解释的尾部延迟。本机代码通常可提供出色的延迟属性，不会给已经令人困惑的情况带来额外的混乱。与其他用C编写的本机代码代理解决方案不同，C ++ 11提供了出色的开发人员生产力和性能。  

3、**L3 / L4过滤器体系结构 :** Envoy是L3 / L4网络代理。可插入的 过滤器链机制允许编写过滤器以执行不同的TCP代理任务，并将其插入主服务器。已经编写了过滤器来支持各种任务，例如原始TCP代理， HTTP代理，TLS客户端证书认证等。  




### Envoy的线程模型是如何的

### Envoy可以做些什么

### 上手试连，通过Envoy实现一个百度简单代理