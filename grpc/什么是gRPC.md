---
title: 什么是gRPC
date: 2020-06-17 09:05:00
tags: grpc
author: 小笼包
categories: grpc
---

## 什么是gRPC

> 作者: 小笼包  
> 2020-06-17 有雨

gRPC是Google开发的高性能、通用的开源RPC框架，其由Google主要面向移动应用开发并基于HTTP/2协议标准而设计，基于Protobuf(Protocol Buffers)序列化协议开发，且支持众多开发语言。

在gRPC中客户端可以像它是本地对象一样直接调用其他计算机上的服务器应用程序，与许多RPC系统一样，gRPC围绕定义服务的思想，指定可通过其参数和返回类型远程调用的方法。在服务器端，服务器实现此接口并运行gRPC服务器以处理客户端调用。在客户端，客户端具有一个存根（在某些语言中仅称为客户端），提供与服务器相同的方法。

gRPC客户端和服务器可以在各种环境中运行并相互通信，并且可以使用gRPC支持的任何语言编写。因此，例如，您可以使用Go，Python或Ruby的客户端轻松地用Java创建gRPC服务。此外，最新的Google API的接口将具有gRPC版本，可让您轻松地在应用程序中内置Google功能。

<!-- more -->

## 使用Protobuf协议

默认情况下，gRPC使用Protobuf协议，它提供了一种灵活、高效、自动序列化结构数据的机制，可以联想 XML，但是比 XML 更小、更快、更简单。仅需要自定义一次你所需的数据格式，然后用户就可以使用 Protobuf 编译器自动生成各种语言的源码，方便的读写用户自定义的格式化的数据。与语言无关，与平台无关，还可以在不破坏原数据格式的基础上，依据老的数据格式，更新现有的数据格式。

Protobuf 的特点简单总结如下几点：

* 作用与 XML、json 类似，但它是二进制格式，性能好、效率高  

* 代码生成机制，易于使用

* 解析速度快

* 支持多种语言

* 向后兼容、向前兼容

* 缺点：可读性差

使用协议缓冲区的第一步是为要在原始文件中序列化的数据定义结构：这是带有.proto扩展名的普通文本文件。协议缓冲区数据被构造为 消息，其中每个消息都是信息的小逻辑记录，其中包含一系列称为字段的名称/值对。这是一个简单的例子：

``` shell
message Person {
      string name = 1;
      int32 id = 2;
      bool has_ponycopter = 3;
}
```

然后，一旦您指定了数据结构，就可以使用ProtoBuf的编译器protoc从原型定义中以首选语言生成数据访问类。这些为每个字段（例如name()和）提供了简单的访问器set\_name()，以及将整个结构序列化为原始字节或从原始字节中解析出整个结构的方法。因此，例如，如果您选择的语言是C ++，则在上面的示例中运行编译器将生成一个名为Person的类。然后，您可以在应用程序中使用此类来填充，序列化和检索Person协议缓冲区消息。

您可以在普通的原始文件中定义gRPC服务，并使用RPC方法参数和返回类型指定为ProtoBuf的消息：

``` shell
// The greeter service definition.
service Greeter {
    // Sends a greeting
    rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.message HelloRequest {  string name = 1;}
```  

gRPC protoc与特殊的gRPC插件一起使用，可从您的原型文件生成代码：您将生成生成的gRPC客户端和服务器代码，以及用于填充，序列化和检索消息类型的常规协议缓冲区代码。

ProtoBuf开源用户已经使用了一段时间，大多数示例都使用协议缓冲区版本3（proto3），该协议的语法略有简化，提供了一些有用的新功能，并且支持更多语言。Proto3当前可用于Java，C ++，Dart，Python，Objective-C，C＃，精简版运行时（Android Java），Ruby和JavaScript的 协议缓冲区GitHub存储库中，以及golang / protobuf GitHub存储库中的Go语言生成器 ，正在开发更多语言。虽然您可以使用proto2（当前的默认协议缓冲区版本），但我们建议您将proto3与gRPC一起使用，因为它可以使用所有gRPC支持的语言，并且可以避免与proto2客户端通信时出现的兼容性问题。

文档参考： [Introduction to gRPC](https://grpc.io/docs/what-is-grpc/introduction/)