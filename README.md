```
                      _ _  _
     ___  ___  _   _ | / _\ \  _ __
    / __|/ __|| |_| || |_ /_/ / _ _'
    \__ \\__ \| |_| || | \ \ ' (_ _
    |___/|___/|_| |_||_|  \_\_\_ __'

```
sshrc依赖于cobra包构建命令行支持，该包是支持通用的命令行构建库。该工具主要适用于linux系统，开箱即用，无需部署环境变量
## 功能
* 远程执行命令，支持复杂命令。
* 拷贝本地主机文件到远程主机；拷贝远程的文件到本地
## 使用sshrc
### centos安装sshrc
* 下载二进制文件后，执行如下命令  
```
chmod +x main && mv main /usr/local/bin/
```  

[下载链接]:https://github.com/laoshangcai/sshrc/releases/download/v1.1.0/sshrc
* [下载链接]  

### 远程执行命令示例  
``` 
sshrc --host 192.168.1.10 --user root --passwd your-server-password --cmd "ls -a"
``` 
* 多台主机执行命令 
```
sshrc --host 192.168.1.10 \  
   --host 192.168.1.11 \  
   --host 192.168.1.12 \  
   --user root \  
   --passwd your-server-password \  
   --cmd "ls -a"

注：多台主机，密码需一致
```

>--host 主机IP  
>--user 主机用户；默认使用root用户，使用其他用户需要使用此参数指定  
>--passwd 主机用户密码  
>--cmd 需要执行的命令  


### 拷贝文件示例 
* 拷贝本地主机文件到远程主机  
```
sshrc copy --host 192.168.1.10 --passwd your-server-passwor --src /opt/kube.tar.gz --dest /opt
``` 
>copy    子命令，指定使用拷贝模块：拷贝本地主机文件到远程主机  
>--src   本地文件路径  
>--dest  远程主机文件存放路径  
* 拷贝远程主机文件到本地主机  
``` 
sshrc fetch --host 192.168.1.10 --passwd your-server-passwor --src /opt/kube.tar.gz --dest /opt
``` 
>fetch   子命令，指定使用拷贝模块：拷贝远程主机文件到本地主机  
>--src   本地文件路径  
>--dest  远程主机文件存放路径  

### 注：当前版本未支持使用配置文件配置主机信息