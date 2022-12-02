---
title: 安装window11的注意事项
date: 2022-09-05 16:14:00
tags: system
author: 小笼包
categories: windows
---

> 作者: 小笼包  
> 2020-10-09 晴

## 通过U盘录制windows11的工具

首先我们去微软官方网站下载**创建 Windows11 安装**  

地址：<https://www.microsoft.com/zh-cn/software-download/windows11>  

下载完毕后启动软件跟着一步步录入至U盘下。

重启，bois设置u盘第一启动，然后选择安装，这时候会可能会出现，如下问题描述：

<!-- more -->

``` shell
这台电脑无法运行windows 11

这台电脑不符合安装此版本的 Windows 所需的最低系统要求。有关详细信息，请访问https://aka.ms/WindowsSysReq
```

这个时候我们需要在注册表中添加几个东西即可继续安装，操作方式如下：

- 快捷键 **Shift + F10** 打开cmd命令窗口， 然后输入 **regedit** 打开注册表
- 点开 **HKEY_LOCAL_MACHINE/SYSTEM/Setup**下添加 **项** ，并命名为labConfig
- 在 **labConfig** 项下新建 **DWCRD（32位）值（D）** 并命名为**BypassTPMCheck** 修改数据值为 **00000001**
- 在 **labConfig** 项下新建 **DWCRD（32位）值（D）** 并命名为**BypassSecureBootCheck** 修改数据值为 **00000001**

然后关闭注册表和命令行窗口，点击回上一步操作，继续下一步安装即可正常安装完成window11了。

至于怎么激活windows11这里就不描述了，自行百度去，建议不要用小马哥的里面流氓软件太多。

## 好用的windows基本软件推荐

- 火绒安全卫士，协助抵挡一些基本的病毒外，还能清理垃圾
- 7z解压软件，小巧轻型，无广告
- PotPlayer视频软件
- geek卸载助手
- GamePP游戏助手
- 硬件狗狗（一款可以跑分和监控硬件温度的工具）
- Mem Reduct <https://www.henrypp.org/product/memreduct> 清理内存的小工具

## N卡的GeForce下载地址

<https://www.nvidia.cn/geforce/geforce-experience/>
