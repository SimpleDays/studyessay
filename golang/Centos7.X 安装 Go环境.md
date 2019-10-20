# Centos7.X 安装 Go环境

### 如何操作
* 1、从 **https://studygolang.com/dl** Golang中国上下载 资源包 **go1.11.linux-amd64.tar.gz**。
* 2、安装包解压命令：`tar -C /usr/local -xzf go1.11.linux-amd64.tar.gz`  我解压到了 **/usr/local** 下统一管理。
* 3、添加环境变量到当前账户，进入当前账户根目录，举例 账户 **user**  进入  **/home/user** 目录下： `$： cd /home/user` 通过 `$: ls -a`找到 目录下的 **.bashrc** 
* 4、编辑 .bashrc 文件 `$: vim .bashrc`
* 5、添加如下代码：
> export GOPATH=/home/user/go<br/>
export GOROOT=/usr/local/go<br/>
export GOBIN=$GOROOT/bin/<br/>
export PATH=$PATH:$GOBIN<br/>

#### 注意
GOPATH 是自定义go开发项目路径将环境变量加载到内存中，执行`$: source /home/user/.bashrc` 。
运行 **go env**  和 **go version** 看下go环境是否生效。

