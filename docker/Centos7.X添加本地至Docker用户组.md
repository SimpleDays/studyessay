# Centos7.X添加本地至Docker用户组

### 前言
根据上篇文章[《Centos7.X通过rpm包安装Docker》](https://github.com/SimpleDays/studyessay/blob/master/docker/Centos7.X%E9%80%9A%E8%BF%87rpm%E5%8C%85%E5%AE%89%E8%A3%85Docker.md)安装好Docker之后，发现必须使用sudo提权的方式或者直接使用root方式才能操作docker，实际上我们平常登录都是自己的账户，这样操作实在不方便，我们可以通过把本地用户添加至Docker用户组来实现自己账户直接操作Docker。

### 如何操作
* 1、如果还没有 docker group 就添加一个（正常都会有）： `sudo groupadd docker`
* 2、将用户加入该 group 内。然后退出并重新登录就生效啦: `sudo gpasswd -a [本地用户名] docker`
* 3、重启 docker 服务： `sudo service docker restart`
* 4、切换当前会话到新 group 或者重启 X 会话 : `sudo newgrp - docker OR pkill X`

### 注意
> 最后一步是必须的，否则因为 groups 命令获取到的是缓存的组信息，刚添加的组信息未能生效，所以 docker images 执行时同样有错。<br/>
> 如果上面操作本地还不能访问，那么我们就注销下自己账户重新登录即可。