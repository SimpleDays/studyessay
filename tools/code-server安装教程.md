---
title: code-server安装教程
date: 2022-09-01 17:14:00
tags: tools
author: 小笼包
categories: life
---

## code-server安装教程

> 作者: 小笼包  
> 2022-09-01 晴

### 一、介绍下code-server是做什么的？

官网的描述是:

> 在任何地方的任何机器上运行VS Code并在浏览器中访问它

- 在具有一致开发环境的任何设备上编写代码
- 使用云服务器加速测试、编译、下载等
- 在旅途中保持电池寿命；所有密集型任务都在您的服务器上运行

也就是说把我们常用的vscode工具通过bs的形式呈现给我们，让我们通过浏览器就可以编码测试和调试，这样子只要我们拥有浏览器和网络的地方即可随时随地的进行编码开发了，统一的云环境让我们可以快速的上手编码，vscode通用性的组件让我们爱不释手，看到code-server之后我就想在自己的云服务器上安装下。

<!-- more -->

### 二、建议安装的云服务器要求

根据官方的建议至少要有1 GB 内存和2个CPU核心。
具体的要求可以参考链接：
[云服务器的设备要求](https://github.com/coder/code-server/blob/main/docs/requirements.md)

### 三、通过Docker容器方式进行安装

官方提供了容器镜像安装的方式，我觉得很方便，这里就介绍下如何通过容器镜像进行安装code-server，以及配置相关开发环境和一些基础设施，同时把遇到的问题一同记录下来，并提供对应的解决方案，经供参考。

#### 1、首先拉去官方镜像

``` shell
docker pull lscr.io/linuxserver/code-server:latest
```

镜像大概有603MB大小，同时国内下载比较慢，如有镜像源可以考虑使用镜像源下载基础镜像。

[Dockerhub上的镜像源地址](https://hub.docker.com/r/linuxserver/code-server)

#### 2、提供下我的容器启动脚本

``` shell
#!/bin/bash

docker rm -f code-server

docker run -d \
  --name=code-server \
  --privileged=true \
  --security-opt seccomp=unconfined \
  -e PUID=1000 \
  -e PGID=1000 \
  -e TZ=Asia/Shanghai \
  -e PASSWORD='自定义密码' \
  -e SUDO_PASSWORD='自定义密码' \
  -p 8443:8443 \
  -v ${pwd}/config:/config \
  --restart unless-stopped \
  linuxserver/code-server
```

#### 3、对于容器启动脚本官方镜像的一些配置说明

|范围|功能|
|:-----:|:-----:|
|-p 8443|网页对外暴露的端口默认8443|
|-e PUID=1000|对于用户 ID - 请参阅下面的说明|
|-e PGID=1000|对于 GroupID - 请参阅下面的说明|
|-e TZ=Asia/Shanghai|指定时区以使用 EG Europe/London ZH Asia/Shanghai|
|-e PASSWORD='自定义密码' |可选的 web gui 密码，如果提供PASSWORD或HASHED_PASSWORD 不提供，将没有 auth。|
|-e SUDO_PASSWORD='自定义密码' |如果设置了此可选变量，用户将使用指定的密码在代码服务器终端中进行 sudo 访问。|
|-e SUDO_PASSWORD='自定义密码' |如果设置了此可选变量，用户将使用指定的密码在代码服务器终端中进行 sudo 访问。|
|-v ${pwd}/config:/config |包含所有相关的配置文件,以及你的相关配置目录都在下面可以挂载，以防丢失。|
|--privileged=true |container内的root拥有真正的root权限,可以看到很多host上的设备，并且可以执行mount, 不建议使用这个配置，应该也不需要，不太安全, 官方也没建议加|
|--security-opt seccomp=unconfined |运行时不使用默认的seccomp配置文件, 为了解决golang无法调试编译问题所加上的参数，如果你没有这个需求可以不加，非官方必要参数。|

[其余的配置可以参考官方文档](https://github.com/linuxserver/docker-code-server#parameters)

#### 4、对于登录密码生成

对于密码生成，我这里提供一个网址：<http://www.icosaedro.it/PasswordGenerator.htm>，提供一个复杂又随机的密码。

#### 5、对于用户组的设置说明

当在主机操作系统和容器之间使用卷（-v标志）权限问题时，我们通过允许您指定用户PUID和组来避免这个问题PGID。

确保主机上的任何卷目录都归您指定的同一用户所有，并且任何权限问题都会像魔术一样消失。

在这种情况下PUID=1000，PGID=1000找到你的用途id user如下：

``` shell
  $ id username
    uid=1000(dockeruser) gid=1000(dockergroup) groups=1000(dockergroup)
```

通常我们用户和用户组默认就都是1000所以不需要额外修改。

### 四、构建好容器后，对于基础环境的一些简单配置

容器启动后，确认容器已经完整启动之后，通过浏览器打开code-server，并通过密码登录。
通过快捷键 **ctrl+~** 打开终端窗口，可以操作和配置容器的基础环境，提供的镜像是Debain的操作系统，git是天然带上的。

``` shell
# 先更新下操作系统
$abc: sudo apt update

# 安装wget工具方便下载一些资源包等应用
$abc: sudo apt -y install wget

# 安装telnet工具方便连接一些tcp的端口等
$abc: sudo apt -y install telnet

# 安装vim方便编辑一些文件等
$abc: sudo apt -y install vim

# 通过ssh的命令生成公私钥,记住三次回车
# 如果git使用id_rsa.pub里面的内容绑定，下载git相关的项目等
$abc: ssh-keygen -t rsa -C "example@email.com"
```

### 五、构建golang开发环境

本人近些年一直学习和使用golang，所以优先构建golang环境和配置调试和开发的插件。

#### 1、下载golang的sdk并配置环境变量

在[golang的中文社区](https://studygolang.com/dl)网站中下载自己需要的sdk版本，
比如最新的1.9的下载地址：<https://studygolang.com/dl/golang/go1.19.linux-amd64.tar.gz>

``` shell
$abc: wget -O ${想要下载的位置，例如：/config/download/} https://studygolang.com/dl/golang/go1.19.linux-amd64.tar.gz
```

解压sdk并放到想要的位置

``` shell
$abc: tar -C ${想要解压的位置，例如：/config/tools/} -xzf go1.19.1.linux-amd64.tar.gz
```

在当前用户下配置golang的环境变量

``` shell
$abc: cd

$abc: vim .bashrc

#在最后一行加入如下:
export GOPATH=/config/go
export GOROOT=/config/tools/go
export GOBIN=$GOROOT/bin/
export PATH=$PATH:$GOBIN
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct

# 重新加载bashrc文件
$abc: source .bashrc

# 确认golang是否已经配置完毕
$abc: go env

$abc: go version
```

#### 2、安装golang插件

在扩展（Extension）下搜索Go的组件第一个直接安装。
**这里注意一个细节，如果你先装了插件，再安装了golang的sdk，可能需要重启下code-server容器，让golang的组件能读取到sdk的环境变量配置，如果是提前装的应该是不需要的。**

#### 3、创建一个golang项目并配置代码调试和单元测试功能

``` shell
$abc mkdir testgo

$abc cd testgo && go mod init testgo

$abc touch main.go
```

等待golang配置加载，并提示下载安装golang的几个组件， gopls，dlv, golint, gooutline, gopkgs等，等待安装成功。

#### 4、调试代码过程中遇到的问题

在当前运行项目下，在左侧边栏找到 “Run and Debug”图标（应该是倒数第三个），选择 “create a launch.json file"。
它会为你当前项目下创建一个**.vscode**的文件，并且提供一个launch.json的文件模板供你修改，一般我们改下参数”program“即可，代码如下：

``` json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}"
        }
    ]
}
```

然后直接快捷键**f5**来进行运行代码，这时候我遇到了一个棘手的问题，代码无法运行调试，报出了如下错误：

``` shell
# 这个可执行文件，无法被操作的错误，没有具体的细节，只有这么一处错误，让人感觉很莫名
# 原本本地的vscode软件我是可以调试运行的，可到了这里突然报这个错误就有点摸不着头脑了
 could not launch process: fork/exec config/workspace/testgo/__debug_bin: operation not permitted
```

根据我google的资料大量翻阅，并也修改了launch.json的代码，让其打印出来相关日志查看，最后还是把这个无法调试运行的问题解决了，
这里先不分析导致的原因是什么，先提供解决方案。
如下是我修改可以看更多日志信息的launch.json的配置信息：

``` shell
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "logOutput": "dap",
            "trace": "verbose",
            "showLog": true
        }
    ]
}
```

[参考的issue](https://github.com/go-delve/delve/issues/515)
这里面我看到这么一句话如下：

> derekparker commented on 27 Apr 2016
> Alright, so you're running within a Docker container. Docker has security settings preventing ptrace(2) operations by default with in the container. Pass --security-opt=seccomp:unconfined to docker run when
> starting.

根据上述的描述意思就是Docker具有安全设置，默认情况下防止容器中的ptrace（2）操作，那么linux下ptrace是什么？
ptrace() 是一个由 Linux 内核提供的系统调用，允许一个用户态进程检查、修改另一个进程的内存和寄存器，通常用在类似 gdb、strace 的调试器中，用来实现断点调试、系统调用的跟踪。
所以也就是默认情况下容器阻止它了，golang通过dlv进行调试跟踪时候用到ptrace的功能，导致没有操作权限，所以后续我就在容器启动时添加了**--security-opt=seccomp:unconfined**
然后就可以正常调试和运行golang了，同时验证单元测试和单元测试函数调试功能也是正常运行，就此golang的环境已经搭建完毕。

#### 5、了解下vscode下dlv是什么？

dlv就是Delve(<https://github.com/go-delve/delve>) 是 Go 编程语言的调试器。该项目的目标是为 Go 提供一个简单、功能齐全的调试工具。Delve 应该易于调用和使用。
Delve不仅可以用来调试源代码，还可以调试二进制文件。
Delve是GDB调试器的有效替代品。与GDB相比，它能更高的理解Go的运行时，数据结构以及表达式。Delve目前支持Linux，OSX以及Windows的amd64平台。
具体的操作文档可以参考如下地址：<https://github.com/derekparker/delve/tree/master/Documentation>

#### 6、容器上的security-opt又是什么？seccomp:unconfined有代表了什么意思呢？

安全计算模式 ( seccomp) 是 Linux 内核特性。
默认seccomp配置文件为使用 seccomp 运行容器提供了一个合理的默认设置，并在 300 多个系统调用中禁用了大约 44 个系统调用。它具有适度的保护性，同时提供广泛的应用兼容性。
实际上，配置文件是一个白名单，默认情况下拒绝访问系统调用，然后将特定的系统调用列入白名单。
Docker 的默认 seccomp 配置文件是一个允许列表，它指定允许的调用。下表列出了由于不在允许列表中而被有效阻止的重要（但不是全部）系统调用。该表包括每个系统调用被阻止而不是列入白名单的原因。
那这些信息可以参考文档：<https://docs.docker.com/engine/security/seccomp/>

如果设置了unconfined，就代表可以传递unconfined以运行没有默认 seccomp 配置文件的容器。
seccomp有助于以最低权限运行 Docker 容器。不建议更改默认seccomp配置文件。
这边修改主要为了可以方便调试golang。

#### 7、单元测试时候出现缺少gcc配置，如何解决

在linux就十分简单了，操作如下：

``` shell
$abc: sudo apt-get install build-essential

$abc: sudo apt-get  build-dep  gcc

# 确认gcc是否安装成功
$abc: gcc -v
```

#### 8、git拉取项目提交注意事项

如果希望团队多人操作同一个code进行开发的话，那么建议一个项目绑定自己的提交用户名，如果只是自己使用那么直接设置全局的git用户名和邮箱即可。

``` shell
# 如果是全局增加--global
$abc: git config user.name 你的目标用户名

$abc: git config user.email 你的目标邮箱名

# 确认用户名和邮箱配置是否成功了
$abc: git config --list
```

### 六、vscode常用插件推荐和vscode几处基本优化

#### 1、vscode的几处基础优化代码

优化的代码也是借鉴别人的这里记录下，如何操作，点击左侧菜单图标的设置选择**settings**, 然后选择右上角的第一个图标（Open settings json），通过json文件方式打开，输入以下配置代码：

``` json
    "files.autoSave": "afterDelay",
    "files.autoGuessEncoding": true,
    "editor.smoothScrolling": true,
    "editor.cursorSmoothCaretAnimation": true,
    "workbench.list.smoothScrolling": true,
    "editor.cursorBlinking": "smooth",
    "editor.mouseWheelZoom": true,
    "editor.formatOnPaste": true,
    "editor.formatOnSave": true,
    "editor.formatOnSaveMode": "file",
    "editor.formatOnType": true,
    "editor.wordWrap": "on",
    "editor.suggest.snippetsPreventQuickSuggestions": false,
    "editor.acceptSuggestionOnEnter": "smart",
    "editor.suggestSelection": "recentlyUsed",
    "debug.showBreakpointsInOverviewRuler": true,
```

这些配置的具体用处可以参考如下视频:<https://www.bilibili.com/video/BV1Hd4y1o7CN?spm_id_from=333.1007.tianma.1-3-3.click&vd_source=07aaab1dc7483c097c8990e528c71711>

#### 2、vscode常用插件推荐

- Go： golang的常用开发环境组件，golang开发人员使用vscode必装。
- Git Graph：git代码的提交图形展示，方便知道提交的整体过程。
- Error Lens： 错误信息可以在错误的那一行代码旁进行展示，时刻督促你对错误配置和代码进行整改。
- Bracket Pair Colorizer 2： 代码括号嵌套颜色区分，可以更好的让我们区分括号和花括号成对的展示。
- Draw.io Integration: 可以使用vscode画流程图，时序图，并改扩展名可以直接转换成图片等展示在文档中，随时画，随时记录，非常方便。
- Image preview： 可以在配置图片地址的代码处，直接通过鼠标点击，查看图片。
- koroFileHeader： 一个注释生成器，可以提供很多有意思的注释代码格式。
- Material Icon Theme： 一个提供文件icon图标的好用工具。
- Path Intellisense： 对导入路径提供智能的提示功能。
- Prettier-Code formatter： 通用的代码格式化工具。
- Markdown All in One： markdown书写的综合工具。
- Markdownlint： markdown书写时候规范检查，方便书写标准的markdown代码。

vscode插件离线安装地址：<https://marketplace.visualstudio.com/search?target=VSCode&category=All%20categories&sortBy=Installs>

后续未完......

后续如果我还需要安装nodejs或者前端开发的工具，我会继续完善这个文档。
