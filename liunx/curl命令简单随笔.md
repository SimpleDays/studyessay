# curl命令学习
在此之前只会简单使用下curl，最近因公司项目需求对curl做了一些更多的了解。
参考了阮一峰大佬2篇文章：[curl网站开发指南](http://www.ruanyifeng.com/blog/2011/09/curl.html) 和 [curl 的用法指南](http://www.ruanyifeng.com/blog/2019/09/curl-reference.html)
因此重新调整下。

## 入门

### 查看网页源码
```
curl https://www.baidu.com
```
通过 **-o [/文件目录/文件名] url** 下载网页源码到指定目录下生成文件
```
curl -o /home/user/downloads/baidu.html https://www.baidu.com
```

### -L 自动跳转
```
curl -L https://www.baidu.com
```

### 展示头部信息
```
curl -i https://www.baidu.com

curl -I https://www.baidu.com
```
**-i** : 展示头部信息和网页源码信息<br/>
**-I** : 仅仅展示头部信息

### **-v** 显示通讯的全过程(包括端口连接和http的request头信息)
```
curl -v https://www.baidu.com
```
如果还需要更详细相关信息通过 **--trace [/路径/文件名]** 命令来查看
```
curl --trace /home/user/downloads/baidutrace.txt https://www.baidu.com

curl --trace-ascii /home/user/downloads/baidutraceascii.txt https://www.baidu.com
```

### 基本请求操作（默认使用Get方式请求）
```
curl http://localhost:8080?mock=1
```
多个参数建议使用 **-G** 来构造url查询字符串 **-d** 来设置url相关参数
```
curl -G -d 'mock1=1' -d 'mock2=2' http://localhost:8080
```
上述请求相当于 **http://localhost:8080?mock1=1&mock2=2** <br/>

#### 1、curl 测试网关代理某些服务比如 POST
> $ curl -v http://[ip:port] -X POST -H "Content-Type:application/json" -d '{"key": "value"}'

> -w %{time_total} 

> -w：按照后面的格式输出

> time_namelookup：DNS 解析域名[www.taobao.com]的时间

> time_commect：client和server端建立TCP 连接的时间

> time_starttransfer：从client发出请求；到web的server 响应第一个字节的时间

> time_total：client发出请求；到web的server发送会所有的相应数据的时间

> speed_download：下周速度 单位 byte/s

#### 2、curl 测试网关代理某些服务比如GET加上 特有 querystring 比如 json={xxx}
我们公司在请求上还有(http://xxx/xx/xx?json={"key":"value"} ) 这样的请求我们如何通过 curl 请求。
**需要对querystring进行urlencode**

>$ curl -G -v "http://server:8080/test/demo?" --data-urlencode 'json={"ShippingTimeRange":{"Left":"2019-09-01T00:00:00","Right":"2019-09-11T00:00:00","LeftExclusive":false,"RightExclusive":false},"ReturnValue":null}'