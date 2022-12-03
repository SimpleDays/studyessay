# Centos7.XDocker开启远程API端口

### 前言
Docker开启远程端口的目的是可以通过Docker提供的 remoteApi文档 管理Docker并且可以操作Docker下容器，监控容器的各项指标，也可以通过remoteApi去实现自己监控Docker告警系统等。默认Docker并没有启动remoteApi，需要我们修改配置才能生效。

### 如何操作
* 1、默认Centos7.X下配置文件地址在 **/usr/lib/systemd/system/** 下面，修改 **/usr/lib/systemd/system/docker.service** 文件,命令：`sudo vim /usr/lib/systemd/system/docker.service`
* 2、在 **ExecStart=/usr/bin/dockerd** 配置文件后面加上 **-H tcp://0.0.0.0:2375 -H unix://var/run/docker.sock** 保存并退出。
* 3、注 : 端口 **2375** 就是docker remoteApi的 端口，确保此端口linux没有被**占用**。
* 4、执行 **重启 docker** 命令  docker重新读取配置文件，并重新启动docker服务 命令 :  `sudo systemctl daemon-reload && systemctl restart docker`