package cmd

import (
	"flag"
	)

var (
	Choosecfg = flag.String("i", "hosts", "host, user, passwd config")
	Choosecmds = flag.String("c", "cmd", "shell")
)

func Execute() {
	flag.Parse()
	Getcfgvalue(*Choosecfg)

	for i := 0; i < len(InfoList); i++ {
		Cmd(InfoList[0].Host, InfoList[0].User, InfoList[0].Passwd, *Choosecmds)
	}

}