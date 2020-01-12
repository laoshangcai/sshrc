package cmd

import (
	"flag"
	)

var (
	Choosehost = flag.String("h", "192.168.211.130" ,"主机IP地址")
	Chooseuser = flag.String("u", "root" ,"用户名称")
	Choosepassword = flag.String("p","123456","用户密码")
	Choosecmds = flag.String("c", "number", "执行的命令")
)

func Execute() {
	flag.Parse()
	Cmd(*Choosehost, *Chooseuser, *Choosepassword,*Choosecmds)
}