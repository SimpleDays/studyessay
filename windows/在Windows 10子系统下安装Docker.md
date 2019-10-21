# 在Windows 10子系统下安装Docker

## 一、目的
为了方便开发需要在子系统下安装Docker来部署推送一些镜像资源，可惜的是目前的wsl是不支持Docker的守护进程的，但是可以通过安装Docker利用Docker CLI链接到远程的Docker守护进程来实现。

## 二、安装和配置
```
sudo apt update
sudo apt install docker.io
export DOCKER_HOST=tcp://127.0.0.1:2375
```
通过 上面简单三个步骤 就可以使用Docker了。

前提是 你需要提供一台vm或者其他服务器上的可用Docker服务，并且这个Docker服务必须开启远程API端口（默认Docker服务不开启），如何开启可以参考：[Centos7.XDocker开启远程API端口](https://github.com/SimpleDays/studyessay/blob/master/docker/Centos7.XDocker%E5%BC%80%E5%90%AF%E8%BF%9C%E7%A8%8BAPI%E7%AB%AF%E5%8F%A3.md)

注意：直接通过export命令 只能当前会话环境有效，如果想一直生效，最好方式有2种

* 1、修改当前用户 **.bashrc** 文件 ，执行 `vim .bashrc` ，最后一行 增加 `export DOCKER_HOST=tcp://127.0.0.1:2375`，然后保存退出，执行 `source .bashrc` 。

* 2、修改 **/etc/profile** 文件 ， 执行 `sudo vim /etc/profile` ，最后一行 增加 `export DOCKER_HOST=tcp://127.0.0.1:2375`, 然后保存退出， 执行 `sudo source /etc/profile` 。