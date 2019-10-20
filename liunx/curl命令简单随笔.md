# curl命令简单随笔
在linux中通过curl命令进行http请求一些记录

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