# Linux的生成.ssh公私钥
`$: ssh-keygen -t rsa -C your_email@example.com`

经过三次回车，生成 .ssh 文件夹
进入 .ssh 文件夹

`$: cd .ssh`
可以看到 两个文件 （id_rsa ， id_rsa.pub）
> id_rsa 生成的是 私钥
> 
> id_rsa.pub 生成的是公钥

我们只需要暴露 公钥即可。

比如  github 等 通过ssh 拉取代码等都可以用这类方式。
