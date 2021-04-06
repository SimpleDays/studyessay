# 常用的一些docker指令

``` docker images | grep none | awk '{print $3}' ``` : 此命令可以刷选出镜像下带有“none”关键词的镜像id，比如tag为none的。  

``` docker rmi $(docker images | grep "none" | awk '{print $3}') ``` : 把带有none的docker镜像，统一清理掉。  

``` docker system df ``` : 命令，类似于Linux上的df命令，用于查看Docker的磁盘使用情况。

``` docker system prune ``` : 命令可以用于清理磁盘，删除关闭的容器、无用的数据卷和网络，以及dangling镜像(即无tag的镜像)。

``` docker system prune -a ``` : 命令清理得更加彻底，可以将没有容器使用Docker镜像都删掉。注意，这两个命令会把你暂时关闭的容器，以及暂时没有用到的Docker镜像都删掉了…所以使用之前一定要想清楚.。我没用过，因为会清理 没有开启的  Docker 镜像。