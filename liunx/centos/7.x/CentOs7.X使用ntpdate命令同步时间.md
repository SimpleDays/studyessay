# CentOs7.X使用ntpdate命令同步时间

### 1、遇到问题
在平常使用linux系统时不管是虚拟机还是物理机都可能遇到时间不正确的情况，遇到这类问题影响我们通过时间逻辑判断业务逻辑开发等，我们可以使用下面方式简单解决时间不一致或者不正确的问题。

### 2、如何安装
**使用ntpdate命令**

 1. 首先需要安装ntpdate
     >   $ yum update
     >   $ yum install -y ntpdate
     >   $ ntpdate ntp1.aliyun.com //通过阿里云服务器同步上海CST时间

ntp.api.bz ntp (上海服务器)  也是可以通过这个同步

 2. 可以做一个crontab定时任务
     >  $ crontab -e
     >  $ 0 5 * * *  /usr/sbin/ntpdate  -u ntp1.aliyun.com    #5点时间同步