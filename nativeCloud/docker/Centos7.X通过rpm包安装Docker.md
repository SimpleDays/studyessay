# Centos7.X通过rpm包安装Docker

### 前言
Docker已经火了很多年，现在各大公司都会使用它。那么在我们日常开发中也经常使用，比如我就通过Docker方便快捷在本地安装很多基础服务（Redis、Nginx、Mongodb、RabbitMQ、K8s）等方便学习和使用。今天记录下如何通过rpm在centos7.x系统里面安装Docker。

### 1、Docker官网下载rpm包
首先我们去Docker官网下载(docker-ce-18.06.1.ce-3.el7.x86_64.rpm)rpm包，地址 ： https://download.docker.com/linux/centos/7/x86_64/stable/Packages/

### 2、通过liunx命令安装rpm包
进入到安装包所在路径，执行 <font color="#660000"><b>sudo yum install *.rpm</b></font> 命令进行安装。安装好之后设置docker开机自动启动和启动服务 命令 ：<font color="#660000"><b>sudo systemctl enable docker</b></font> 。

### 3、迁移镜像存储路径
这一步可以迁移也可以忽略，主要目的默认路径在 /var/lib 下 ，一般我们linux的 /home 目录容量会大很多，并且方便我们扩展，所以我都会迁移至/home目录下，docker镜像和容器存储容量还是蛮大的。(迁移时候最好提权至root方便操作sudo su)

* 停止 Docker: `systemctl stop docker`
* 为了安全做个备份 `tar -zcC /var/lib/docker > /home/mnt/var_lib_docker-backup-$(date + %s).tar.gz`
* 迁移 /var/lib/docker 目录到 /home/mnt 目录下: `mv /var/lib/docker /home/mnt/`
* 建个 symlink: `ln -s /home/mnt/docker /var/lib/docker`
* 启动 `systemctl start docker`
  
 好了至此我们的Docker算是安装完毕了，可以通过官网文档说的（https://docs.docker.com/install/linux/docker-ce/centos/）测试下
 `sudo docker run hello-world` 