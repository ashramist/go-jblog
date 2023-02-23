# 账号配置

# 当前项目配置

git config --local user.name '<用户名>'

git config --local user.email '<github 邮箱地址>'

# 分支操作

# 分支创建

# 1. 仓库未创建时

    git init -b <分支名>

# 2. 仓库已经创建

     # 2.1 仅创建分支
       git branch <分支名>
     # 2.2  创建并切换到分支
        git checkout -b <分支名>

# 查看分支

git branch -a
git branch -r

# 切换分支

git chekcout <分支名>

# 删除分支

git branch -d
git branch -D

# 合并分支

git merge <所要合并的分支>

# 重命名分支

git branch -m <原来名称> <新分支名称>

# 从版本库文件恢复

# 1. 已删除（或者修改），未 add

    git checkout <文件路径>

# 2. 已删除（或修改），已 add, 未 commit

    未提交的过程中，不能从缓存区去恢复文件，得从仓库中去获取原始文件，因此得找到上次的commitId

    commitID 通过 git log 命令查看

    git checkout <commitId> <文件路径>

# 3. 已删除（或修改），已 commit,push

  未提交的过程中，不能从缓存区去恢复文件，得从仓库中去获取原始文件，因此得找到上次的commitId

  commitID 通过 git log 命令查看

  git checkout <commitId> <文件路径>

# 关联远程库  
  
  1. 本地仓库的创建
      git init -b <分支名>
  2. 远程仓库的创建
  3. 本地与远程关联
     origin 为远程库名称 
     git remote add origin <远程库地址>
  4. 代码提交到远程库 
     git push -u origin <分支名>


  5. 关联其他远程仓库
   git remote add gitee <远程库地址>

   git push gitee <分支名>
  
# commit 提交修正  (前提还没有推送到远程仓库)

  问题 1. 代码提交有误或者还需要修改
  git commit --amend -m "<提交注释>"

# 日志 
  日志 显示

     条数限制 git log -10 
     当行显示 git log --oneline
     图表日志 git log  --graph 
     显示更改摘要 git log --stat 
     显示修改位置   git log -p 
  日志筛选
     按照日期 
        after 某个日期（不包含after之前的日期） 在 before 某个日期到 
         git log --after="<日期>" --before="<日期>"
         git log --before= "toady" 等同于 git log ,之前用yesterday 
         git log --before="30 day ago"
         git log --after="4 week ago"
         git log --after="1 month ago"
    按照作者
       git log --author="gary"
       git log --author="@qq.com"
    按提交信息 
       git log --grep="重要更新"
  日志引用
     git reflog 


# git 远程分支操作 
 推送 某个分支为 

    git push <远程库remote名称> <分支名>

 所有分支
   git push <远程库remote名称> --all 


   删除远程分支 

   git push <远程库remote名称> -d  <分支名>
