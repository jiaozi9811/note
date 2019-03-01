# git使用

[TOC]

## Pro Git

<http://iissnan.com/progit/>

## git

### 全局设置

>git config --global user.name " "
>git config --global user.email " "

### 本地仓库初始化

>git init

### 添加远程仓库

>git remote add origin <https://github.com/google/go>

### 同步操作

> git add .
> git commit -m "commit commend"
> git pull --rebase origin master
> git push -u origin master

### 查看状态

>git status

### 查看commit记录

>git log

### 查看之前的commit

git checkout [commit] [file]
git checkout [commit]
git checkout [branch]

### 撤销公共修改

git revert [commit]

### 撤销本地修改

git reset
git clean

### 重写Git历史记录

git commit --amend
git rebase
git reflog

### 分支

git branch
git checkout
git merge

### 仓库同步

git remote
git fetch
git pull
git push




## ERROE

### error: src refspec master does not match any

- [x] 原因：

>本地仓库为空

- [x] 解决方法：
  
>git add .
>git commit -m "commit change"