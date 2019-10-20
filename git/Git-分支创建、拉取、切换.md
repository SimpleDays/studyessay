# Git-分支创建、拉取、切换
***
#### 查看本地分支
`git branch`

#### 查看本地与远程分支
`git branch -a`

### 1. 创建本地分支
> git branch [分支名称]

`git branch v1.0.0`

### 2. 创建远程分支/推送远程分支
> git push --set-upstream origin [分支名]

`git push --set-upstream origin v1.0.0`

### 3. 删除本地分支
> git branch -d [本地分支名称]

`git branch -d v1.0.0`

删除远程分支
> git push origin -d [分支名称]

`git push origin -d v1.0.0`

### 4. 拉取远程分支
> git checkout -b [本地分支名称] [远程分支名称]

`git checkout -b v1.0.0 origin/v1.0.0`

