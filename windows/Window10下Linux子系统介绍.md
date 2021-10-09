---
title: Window10下Linux子系统介绍
date: 2021-10-09 17:40:00
tags: system
author: 小笼包
categories: windows
---

## Window10下Linux子系统介绍

> 作者: 小笼包  
> 2020-10-09 晴

### 一、介绍
**wsl**（Windows Subsystem for Linux的简称）是 Windows10下的Linux子系统，在wsl环境下我们可以运行一些Linux程序。wsl提供了一个微软开发的Linux兼容内核接口（不包含Linux代码），来自Ubuntu的用户模式二进制文件在其上运行。该子系统不能运行所有Linux软件，例如那些图形用户界面，以及那些需要未实现的Linux内核服务的软件。不过，这可以用在外部VcXsrv服务工具上实现图形界面操作， 比如：https://blog.csdn.net/qq_20464153/article/details/79682274 。此子系统起源于命运多舛的Astoria项目，其目的是允许Android应用运行在Windows 10 Mobile上。此功能组件从**Windows 10 Insider Preview build 14316**开始可用。

### 二、安装与启动
安装前注意，wsl是在Windows 10在一周年更新（1607，内部版本14393）的时候加入的beta，到了Windows 10（1709，内部版本16299），wsl才正式脱离beta，逐渐趋于稳定，所以使用前确认自己的Windows10当前的版本，可以通过cmd命令，输入 “ systeminfo | findstr Build “ 命令，查看自己当前版本，如果版本太低，升级自己的Windows 10系统。

<!-- more -->

![](https://github.com/SimpleDays/studyessay/blob/master/windows/images/cmd.png)

* 1、安装之前需要打开设置，进入应用，选择程序和功能，点击启用或关闭Windows功能，勾选适用于Linux的Windows子系统（如下图），接下来重启。

![](https://github.com/SimpleDays/studyessay/blob/master/windows/images/windowsgn.png)

* 2、打开Windows 10下的 Microsoft Store应用商店里面搜索Ubuntu，选择安装Ubuntu或者Ubuntu18.04 LTS应用（都是免费的）。
* 3、装完打开应用会进行初始化加载，和 创建用户名密码（就是Ubuntu登录账户密码，可以 通过sudo su来提权到root）。到此Windows子系统安装完成，可以使用了。
* 4、疑问：既然是子系统，那么Windows下的磁盘（C,D,E,F）等，在子系统下如何访问呢，其实windows的目录全部挂载在/mnt这个目录下，如果是C盘下的就在"/mnt/c/" 下，如果是E盘的就是"/mnt/e/"下（如下图）。这样子很方便用于编程开发和学习中。

![](https://github.com/SimpleDays/studyessay/blob/master/windows/images/cmd2.png)