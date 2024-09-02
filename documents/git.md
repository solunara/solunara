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
查看所有分支名称及最近一次的信息：
```
$ git branch -av
```
根据commit hash值查看对应的文件类型（cmmit/tag/tree/blob）:
```
$ git cat-file -t 8a18528
commit
```
根据commit hash值查看对应的文件内容:
```
$ git cat-file -p 8a18528
tree fc2867590e0e68230e7b4779a7de1ff5bf0b69f7
parent f75d6b6b1905b856d2590c495cf27bccd75c5a2d
author frodelee <frodeli@163.com> 1725288921 +0800
committer frodelee <frodeli@163.com> 1725288921 +0800

add Document/git.md
```
## commit,tree,blob三个对象的关系
commit: 代表一次提交，是此次提交中所有文件夹及文件的快照；

tree：代表文件夹，里面可能有子文件夹或文件；

blob：代表文件，git里面不会有重复的文件，只要内容一样就只会有一份；

git通过这种树形结构保存信息，然后根据hash一层一层寻找文件。

![commit,tree,blob](../pictures/git/git_commit_tree_blob.png)


## .git文件夹
全部文件列表：
```
Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2023/4/17     22:57                hooks
d-----         2023/4/17     22:57                info
d-----         2023/4/17     22:57                logs
d-----         2023/4/17     23:42                objects
d-----         2023/4/17     22:57                refs
-a----          2024/9/2     22:55             20 COMMIT_EDITMSG
-a----          2024/9/2     22:30            419 config
-a----         2023/4/17     22:57             73 description
-a----          2024/9/2     22:59            225 FETCH_HEAD
-a----          2024/9/2     22:41            342 gitk.cache
-a----          2024/9/2     22:54             24 HEAD
-a----          2024/9/2     22:55            435 index
-a----          2024/9/2     22:53             41 ORIG_HEAD
-a----         2023/4/17     22:57            185 packed-refs
```
HEAD： 它描述了我们当前工作的分支，如果切换了分支，HEAD文件也会随之改变。

config： 它存放了跟本地仓库相关的配置信息。

refs： 描述了全部分支索引和tag信息。

objects: 存储了全部文件及hash值。