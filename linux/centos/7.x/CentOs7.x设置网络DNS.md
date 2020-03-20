# CentOs7.x设置网络DNS

### 如何操作
* 1、先通过 **ip addr** 确认下 当前网络配置 一般默认  **eth0** <br/>
`ip addr`

* 2、打开网络脚本配置位置 <br/>
`cd /etc/sysconfig/network-scripts`

* 3、修改 **network-script** 内的对应网络配置文件 <br/>
`sudo su vim /etc/sysconfig/network-scripts/ifcfg-eth0`

* 4、增加如下代码块，如有多个DNS <br/>
```
DNS1=XXX.XXX.XXX.XXX
DNS2=XXX.XXX.XXX.XXX
DNS3=XXX.XXX.XXX.XXX
```

* 5、保存 - 重启 **network** 网络配置
`service network restart`