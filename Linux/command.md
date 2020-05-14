### 常用命令
```shell
Git Commends 3 ❌✅✅✅✅✅✅✅
git status
git diff —staged
git diff —check
git mv a.md b.md
git rm —staged
git rm a.md
git log -p -2 比较最近两次提交
git log —stat 查看提交的变更文件
git log —pretty=oneline
git log —pretty=format:”%h - %an,%ar:%s”
git log —graph
git log —since=2.weeks
git log —author=Mr.Trans
git commit —amend -m “sss”
git fetch origin/[remote-name]
git remote show origin 查询跟踪状态
git remote rm [remote-name]
git tag
git tag -a v1.0 -m “first object” git show v1.0 展示标签附加提交信息
git tag -a v1.1 tag202002212202 // 部分校验和
git push origin v1.0
git tag -d v1.0
git tag -d :refs/tag/v1.0
git checkout -b master v2.0 // 前向移动，会和主分支不同，需要合并
git config —global alias.co checkout
git config —global alias.last ‘log -1 HEAD’
git log —oneline —decorate —graph /-all
git branch -v
git branch —merged/—no-merged
git rebase/merge 将提交到分支的修改都移到另一分支上，像重新播放一样
git add —patch 暂存
git reflog 查看引用日志
git ^refa refb 不被refa包含的内容 但在refb包含的内容
```
```shell
Linux Commends 3 ❌✅✅✅
cat -n f1 > f2
chattr -a /var/log/message
chgrp -v group1 flies
chgrp --reference=file1 file2
chmod r=4 w=2 x=1
chmod ug=rwx,o=x file1
chown user[:group] file1
awk -F, ‘{print  $1,$4}’ filename
awk -va=1 ‘{print $1+a}’ log.txt
awk -f cal.awk score.txt
begin {} {} end{}

# seq 9 | sed 'H;g' | awk -v RS='' '{for(i=1;i<=NF;i++)printf("%dx%d=%d%s", i, NR, i*NR, i==NR?”\n”:”\t”)}’每一行执行
cmp file1 file2 判断两个文件是否有差异
diff file1 file2 比较两个文件的不一样
file file1 查看文件是什么类型
find *.o 在指定目录下查找文件
find / -size +500M -print0|xargs -0 du -m|sort -nr
find指令为找出500M以上的文件，print0和xargs -0配合使用，用来解决文件名中有空格或特殊字符问题。du -m是查看这些文件的大小，并以m为单位显示。最后sort -nr是按照数字反向排序（大的文件在前）
In 为某一个文件在另一个位置新建一个同步链接 ln -s file1 file2
locate file1 在自己的数据库找，待更新
mv 移动 将文件移动到文件夹 将文件更名 将文件夹移动到目标文件夹 不存在则新建
paste file1 file2 合并输出
rcp 复制远程文件 rcp root@111.1.1.1:/sss sss
split -6 file 每6行分割一个小文件
touch file 修改文件属性，若不存在则新建
which 会在环境变量查找文件$PATH
cp -r dir1/ dir2 复制
whereis 在特定目录下查找二进制，原始代码或者帮助文件
rhmask file1 file2 加密文件
csplit 用于分割文件
fmt 重新编排文件 可自定义
fold -w 30 file1 将文件每一行折叠成30个字符
grep -r hhh folder/ 递归查找文件含有hhh字符，并输出行数
grep -v test test* 查找文件中不包含test的行
find . -name “*.log” | xargs grep “poll”
join file1 file2 将两个文件相同字符串链接起来
look A fille 查找A字母开头的行
nl file 列出行号
nl file | sed -n ‘/xxx/{s/aaa/bbb;p;q}’ 查找xxx关键字行，替换aaa为bbb q退出
expr length “hhhh”.  expr 14%5     expr index h “hhh” 表达式求值
uniq file 删除重复的行
wc file1 file2 我擦命令 打印文件行数 单词数 字节数
let a=5+2 多个表达式求值
df -h 展示磁盘使用率
dir -l dir/ 展示文件记录
du -h file1    或者document/ 展示文件大小
rmdir dir/ 删除dir
rmdir -p dir/test 若删除test后，dir为空目录，则一并删除 
stat testfile 展示文件信息
ls -ltr s* s开头的文件展示 按时间顺序排
ls -lR /bin 展示bin以下所有的目录和文件详细资料
ls -AF 列出当前工作目录的所有文件以及目录
lsof -c 进程名 查看某个进程的文件使用情况
lsof -p 1110 查看进程号1110打开了哪些文件
lsof -I -sTCP:LISTEN  查看tcp监听端口
lsof -i UDP@[url]www.akadia.com:123 //显示那些进程打开了到www.akadia.com的UDP的123(ntp)端口的链接
lsof -t -c file 返回pid
badblocks -s -v /dev/sda2. 检查坏区
nc -v -z -w2 1.01.1.1 80 tcp端口扫描
wall message 发送信息给每个上线使用者
netstat -a  显示网络信息
ping -c 3 ip.  icmp 检测是否与主机相连 发送3次包
ping -I 3 -s 1024 -5 255 Ip//-i  3 发送周期3秒 -s 设置发送包大小 -t 设置ttl值为255
traceroute 显示数据包到主机的路径
tcpdump -c 10 展示十个网络数据
useradd -g root trans 添加用户trans 并制定用户所在组为root
write trans // root与trans谈话
groupdel a // 删除一个群组
last -n 5 -a -i // 显示最近登陆用户信息 -a -i 显示ip地址
lastb // 查询登陆失败用户
ps -u trans // 显示指定用户进程信息
top -n 10 // 显示更新10次后自动退出
pstree -apnh // 显示树状进程关系
uname -a // 显示所有关于计算机的信息
w -s // 显示登陆信息
free -h // 合适展示内存使用状态
crontab -l/-e/-r 定时任务
0 */2 * * * /sbin/service httpd restart  意思是每两个小时重启一次apache 
declare -I ab // 声明变量 -I 整形 -a 数组 -r 只读
export -p // 列出所有环境变量
passwd -S user //.列出用户信息
passwd user // 修改用户密码
time -v date. // 测量显示特定指令执行时所需要的 时间
ulimit -n // 用于控制shell程序的资源
gzip * 压缩文件
tar -czvf   a.tar.gz a.c // 压缩文件a.c
tar -tzvf a.tar.gz // 列出压缩文件内容
tar -xzvf a.tar.gz
xargs 是一个强有力的命令，能够捕获一个命令的输出，然后传递给另外一个命令
xargs 一般是和管道一起使用
somecommand | xargs -item command
cat url-list.txt | xargs wget -c

strace -p pid // 查看某进程调用了那些系统调用
结束了，以后边查边用吧！

go bible 1 chapter ❌✅✅✅✅
(sort)✅(类型switch)✅channel✅⭕️ 结束了 看go夜读

Morning Bird Plan 0206
```

```shell
Docker command 结⭕️ 查着用吧
Docker command
docker ps -q | xargs docker rm -f // 删除所有容器
docker ps -n
docker attach 直接进入容器 退出会终止
docker exec -it etcd1 /bin/bash // 操作容器 退出不会终止
docker top etcd1 // 显示容器运行的进程
docker tag 镜像id respository:dev // tag
docker network create -d bridge trans-net // 建一个桥接网络
```
