---
title: Http压测工具-wrk
date: 2020-03-13 12:40:00
author: 小笼包
tags: tools
categories: studyessay
---

## Http压测工具

> 作者: 小笼包  
> 2020-03-13 晴

## wrk

地址：<https://github.com/wg/wrk>
> wrk 的一个很好的特性就是能用很少的线程压出很大的并发量. 原因是它使用了一些操作系统特定的高性能 io 机制, 比如 select, epoll, kqueue 等. 其实它是复用了 redis 的 ae 异步事件驱动框架. 确切的说 ae 事件驱动框架并不是 redis 发明的, 它来至于 Tcl的解释器 jim, 这个小巧高效的框架, 因为被 redis 采用而更多的被大家所熟知.
<!-- more -->

## 安装方式

## CentOs

安装wrk前需要的一些依赖组件  

  ``` shell
  # sudo yum install curl-devel expat-devel gettext-devel openssl-devel zlib-devel asciidoc  

  # sudo yum install gcc perl-ExtUtils-MakeMaker  
  ```

下载wrk源码包，可以考虑直接通过git拉取下来  

```  shell
git clone https://github.com/wg/wrk.git  
```

编译wrk文件

``` shell
# cd wrk

# make  
```

若出现错误：xmlto: command not found，可以尝试重新安装xmlto:  

``` shell
# yum -y install xmlto  
```

创建软链到指定目录

``` shell
# ln -s /home/user/tools/wrk/wrk /usr/local/bin
```

通过执行“wrk ”来确认是否安装成功

## wrk简单的使用

1、对于简答不带Querystring的url可以直接通过如下命令压测  

``` shell
# wrk -t12 -c2000 -d1m -T10s --latency https://www.baidu.com
```

2、如果GET带有QueryString或者body参数通过lua脚本  

``` shell
wrk -t12 -c2000 -d1m -T10s --latency -s /root/testlua/test.lua http://10.1.62.91

#test.lua
request = function ()
  return wrk.format("GET","/sz/Member/MaxMemberInfoRequest?accesstoken=6f8cbe0a4b2743d2&customerguid=60570d08-5dc9-4603-973a-a73440cd910a&sourcetype=9&MaxMemberVersion=1")
end  

```

## 命令解释

> -c, --connections: total number of HTTP connections to keep open with each thread handling N = connections/threads  
> 总的http并发数,要保持打开状态的HTTP连接总数每个线程处理N=连接/线程。  
>一般线程数不宜过多. 核数的2到4倍足够了. 多了反而因为线程切换过多造成效率降低. 因为 wrk 不是使用每个连接一个线程的模型, 而是通过异步网络 io 提升并发量. 所以网络通信不会阻塞线程执行. 这也是 wrk 可以用很少的线程模拟大量网路连接的原因. 而现在很多性能工具并没有采用这种方式, 而是采用提高线程数来实现高并发. 所以并发量一旦设的很高, 测试机自身压力就很大. 测试效果反而下降.  

> -d, --duration:    duration of the test, e.g. 2s, 2m, 2h  
> 持续压测时间, 比如: 2s, 2m, 2h  

> -t, --threads:     total number of threads to use  
> 使用的总线程数

> -s, --script:      LuaJIT script, see SCRIPTING  
> LuaJIT脚本  
> 一些说明可以参考地址 ： <https://github.com/wg/wrk/blob/master/SCRIPTING>  
> 参考示例：<https://github.com/wg/wrk/tree/master/scripts>  

> -H, --header:      HTTP header to add to request, e.g. "User-Agent: wrk"  
> 添加http header, 比如. "User-Agent: wrk"  
> --latency:     print detailed latency statistics  
> 在控制台打印出延迟统计情况  
> --timeout:     record a timeout if a response is not received within this amount of time.  
> http超时时间  

## 返回结果参数解释

``` shell
[root@sy-suz-srv51 ~]# wrk -t12 -c100 -d10s -T1s --latency  https://www.baidu.com
Running 10s test @ https://www.baidu.com
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   104.40ms  142.88ms 978.03ms   87.06%
    Req/Sec   122.00     49.45   320.00     68.16%
  Latency Distribution
     50%   32.02ms
     75%  152.79ms
     90%  269.98ms
     99%  673.96ms
  14577 requests in 10.03s, 218.66MB read
  Socket errors: connect 0, read 74, write 0, timeout 32
Requests/sec:   1453.37
Transfer/sec:     21.80MB
```

> 12 threads and 100 connections: 总共是12个线程,100个连接(不是一个线程对应一个连接)  

> latency和Req/Sec ：代表单个线程的统计数据,latency代表延迟时间,Req/Sec代表单个线程每秒完成的请求数，他们都具有平均值, 标准偏差, 最大值, 正负一个标准差占比。一般我们来说我们主要关注平均值和最大值. 标准差如果太大说明样本本身离散程度比较高. 有可能系统性能波动很大.  

> 14577 requests in 10.03s, 218.66MB read ：在10秒内总共请求了14577次，总共读取218.66MB的数据  

> Socket errors: connect 0, read 74, write 0, timeout 32 ：总共有74个读错误，32个超时。  

> Requests/sec和Transfer/sec ：所有线程平均每秒钟完成了1453.37个请求,每秒钟读取21.80MB数据量  

> Latency Distribution ：响应时间的分布，上述请求代表，50%请求在32.02ms内完成，90%在269.98ms内完成
