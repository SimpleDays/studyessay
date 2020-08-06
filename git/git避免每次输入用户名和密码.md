# git避免每次输入用户名和密码

使用git pull或者git push每次都需要输入用户名和密码很不人性化，耽误时间

``` shell
git config --global credential.helper store

git pull /git push (这里需要输入用户名和密码，以后就不用啦)
```
