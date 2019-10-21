# CentOS7.X 防火墙暴露端口

### 例子
#### 暴露 **80** 端口执行命令
`sudo  firewall-cmd --zone=public --add-port=80/tcp --permanent`

#### 命令含义：
字段|含义
:-:|:-:
**--zone**|作用域|
**--add-port=80/tcp** |添加端口，格式为：端口/通讯协议|
**permanent**|永久生效，没有此参数重启后失效|

#### 重启防火墙 命令 
`sudo firewall-cmd --reload`