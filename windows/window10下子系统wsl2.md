---
title: Window10下子系统wsl2介绍
date: 2021-10-09 17:40:00
tags: system
author: 小笼包
categories: windows
---

## Window10下子系统wsl2介绍

> 作者: 小笼包  
> 2020-10-09 晴

## 一、介绍

在之前介绍过了windows10如何安装wsl，这次介绍基于wsl1升级新的更稳定、更兼容的子系统wsl2。

## 二、安装步骤

### 1、检查当前windows10版本，因wsl2安装对win10系统版本有所要求

- 对于 x64 系统：版本 1903 或更高版本，采用 内部版本 18362 或更高版本。
- 对于 ARM64 系统：版本 2004 或更高版本，采用 内部版本 19041 或更高版本。
- 低于 18362 的版本不支持 WSL 2, 我们需要通过系统更新来提高win10系统版本。

若要检查 Windows 版本及内部版本号，选择 Windows 徽标键 + R，然后键入“winver”，选择“确定”。 更新到“设置”菜单中的最新 Windows 版本。
<!-- more -->

### 2、在控制面板下启动相关windows功能

- 打开 ”控制面板“ -> ”程序“ - >"启用或关闭windows功能”
- 勾选 “Hypper-v" 和 勾选 ”适用于 Linux的 windows子系统、还有勾选 "虚拟机平台"
- 设置wsl2为默认， ```wsl --set-default-version 2```

### 3、下载Linux内核更新包

- [适用于 x64 计算机的 WSL2 Linux 内核更新包](https://wslstorestorage.blob.core.windows.net/wslblob/wsl_update_x64.msi)

- 在微软应用商店安装 ubunut版本可以选择 18和 20 两个版本，安装完毕之后根据命令查看linux的wsl的版本是否为2.0， ``` wsl -l -v ```

``` shell
PS E:\work\github\studyessay> wsl -l -v
  NAME            STATE           VERSION
* Ubuntu-20.04    Running         2
PS E:\work\github\studyessay>
```

- 如果显示并非wsl2版本，可以通过 ```wsl --set-version Ubuntu-20.04 2``` 来设置。

### 4、安装docker

``` shell
sudo apt-get update
sudo apt-get install docker-ce
service docker start 启动
service docker stop 关闭
service docker restart 重启
```

安装完毕就可以直接通过docker安装各种软件以供开发使用，非常方便，ubuntu即使有问题也可以销毁重新构建。

### wsl2内docker端口访问

根据我使用下来发现，docker内映射出来的端口，在宿主机上可以通过 "localhost" 或者 "127.0.0.1" 来进行访问，如果想通过宿主机暴露端口至局域网需要宿主机上做端口映射才能把wls2下docker容器端口暴露出去，具体做法如下：

``` shell
# 通过管理员方式打开powershell，然后通过以下命令做映射
~> netsh interface portproxy add v4tov4 listenport=【宿主机windows平台监听端口】 listenaddress=0.0.0.0 connectport=【wsl2平台监听端口】 connectaddress=【wsl2平台ip】protocol=tcp

# 可以通过以下命令来查看当前多少端口做了映射
~> netsh interface portproxy show all

# 删除映射端口命令
~> netsh interface portproxy delete v4tov4 listenport=【宿主机windows平台监听端口】 listenaddress=0.0.0.0
```

PS: 当然如果还是不行，排查防火墙是否打开，如果打开可以麻烦些设置入栈规则，或者索性关闭防火墙即可

### 相关安装wsl2资料

<https://docs.microsoft.com/zh-cn/windows/wsl/install-manual>  

<https://blog.csdn.net/qq_36872046/article/details/106437748>

<https://www.jianshu.com/p/c27255ede45f>  

<https://www.jianshu.com/p/0aa542003b93>

<https://blog.csdn.net/cf313995/article/details/108871531>
