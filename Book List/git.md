# Git
## 安装git
下载地址：https://git-scm.com/downloads

安装参考：https://git-scm.com/book/en/v2/Getting-Started-Installing-Git

查看git版本：
```
$ git --version
git version 2.46.0.windows.1
```

## 配置git
--local: 只对某个仓库有效，也是缺省值，优先级最高

--global: 对当前用户所有仓库有效，最常用的

--system: 对系统所有登录的用户有效
```
$ git config --global user.name 'your_name'
$ git config --global user.email 'your_email@domain.com'
```
查看配置，使用--list
```
$ git config --global --list
user.name=frodelee
user.email=frodeli@163.com
```
## 常用命令
以简洁的单行显示方式查看当前分支最近的5次提交log：
```
$ git log -n5 --oneline
```
以图形化的方式查看所有分支的提交日志：
```
$ git log --all --graph
```
以图形化，简洁的方式查看最近5次所有分支的提交日志：
```
$ git log --oneline --all -n5 --graph
```
打开图形界面：
```
$ gitk
```