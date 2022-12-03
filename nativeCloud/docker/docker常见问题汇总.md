# Docker常见问题汇总

记录日常使用docker出现问题记录，时常更新

## 1、Docker 错误 “port is already allocated”

我在启动容器时候发生了如下错误:

``` shell
docker: Error response from daemon: driver failed programming external connectivity on endpoint nginx-proxy (58a751022f039d37fd02c3096f31851e3bb4258244866c744d8cd5595c24ac75): Bind for 0.0.0.0:9376 failed: port is already allocated.
```

查看进程，发现相关的容器并没有在运行，而 docker-proxy 却依然绑定着端口, 如何查看命令如下:

``` shell
$ docker ps
$ ps -aux | grep -v grep | grep docker-proxy
```

停止 doker 进程，删除所有容器，然后删除 local-kv.db 这个文件，再启动 docker 然后 重启这个容器就能正常运行了，命令如下:

``` shell
$ sudo service docker stop
$ sudo rm /var/lib/docker/network/files/local-kv.db
$ sudo service docker start
```