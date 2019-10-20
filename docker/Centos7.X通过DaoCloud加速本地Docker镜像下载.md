# Centos7.X通过DaoCloud加速本地Docker镜像下载

### 前言
根据上两篇[《Centos7.X通过rpm包安装Docker》](https://github.com/SimpleDays/studyessay/blob/master/docker/Centos7.X%E9%80%9A%E8%BF%87rpm%E5%8C%85%E5%AE%89%E8%A3%85Docker.md)和 [《Centos7.X添加本地至Docker用户组》](https://github.com/SimpleDays/studyessay/blob/master/docker/Centos7.X%E6%B7%BB%E5%8A%A0%E6%9C%AC%E5%9C%B0%E8%87%B3Docker%E7%94%A8%E6%88%B7%E7%BB%84.md)我们可以使用Docker了，但是在使用过程中发现有些基础组件镜像很难下载，可以通过DaoCloud来加速我们下载镜像。

### 加速网址
DaoCloud网址 : https://www.daocloud.io/mirror#accelerator-doc 
在网站下面有一系列说明。

### 如何执行
* 1、执行 网址告知的命令 ： `curl -sSL https://get.daocloud.io/daotools/set_mirror.sh | sh -s http://f1361db2.m.daocloud.io`
* 2、然后重启Docker即可：  `sudo systemctl restart docker`


