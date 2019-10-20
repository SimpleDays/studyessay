# Go在window环境下的安装

### 前言
window下安装我直接选择**msi**方式 安装部署。

### 如何操作
* 1、首先从 https://studygolang.com/dl  下载 golang sdk （go1.11.windows-amd64.msi）
* 2、手动安装 go1.9.2.windows-amd64.msi 安装路径自己选择 ，安装完毕之后我发现，相关的两个（GOROOT、PATH）环境变量已经帮我自动注册上去了。
* 3、我们需要手动指定GOPATH 在哪个目录下存放我们Go项目,在环境变量里面加上。
* 4、然后我们在cmd 中 执行 **go env** 和  **go version** 命令确认 go环境是否部署成功！
