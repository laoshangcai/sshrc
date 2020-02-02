#sshrc\
sshrc是一个Go语言开发的程序，可批量远程执行命令。\
1、配置文件名hosts，文件名不可变，必须与程序在同一级目录 配置文件示例:\
   [server1] \
   192.168.211.128,root,123456

   #多台服务器信息配置： \
    [server] \
    192.168.211.128,root,123456 \
    192.168.211.129,root,123456 \
    192.168.211.129,root,123456 \
    192.168.211.129,root,123456

2、程序使用示例：\
     # ./sshrc -i server1 -c "ls -lh" \
    [下载地址](https://github.com/laoshangcai/sshrc/releases/download/v1.0.0/sshrc) \
     参数说明： 
     -i  hosts配置文件中组的关键字，例如 "server1" 。
     -c  执行的命令行。
