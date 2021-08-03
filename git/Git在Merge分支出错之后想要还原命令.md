
# 当遇到使用git合并分支时候遇到大量merge的冲突之后无法下手处理，或则处理出错了想要放弃merge时候 可以考虑以下命令处理

``` shell
git fetch --all

git reset --hard origin/master # 还原当前指定版本不一定必须是master也可以是其他版本分支

git fetch
```
