package cmd

import (
	"strings"
	"lsc.com/sshrc/utils"
)

type Configinfo struct {
	Host string
	User string
	Passwd string
}

var InfoList = make([]Configinfo, 0)
var info = Configinfo{}

func Getcfgvalue(group string) {
	var conf= utils.InitConfig("hosts")
	//var conf= utils.InitConfig("D:/awesomeProject/src/lsc.com/sshrc/hosts")
	value := conf.GetValue(group)

	for i := 0; i < len(value); i++ {
		result := strings.Index(value[i], ",")
		if result == -1 {
			info = Configinfo {Host: value[i]}
		} else {
			l := strings.Split(value[i], ",")
			info = Configinfo {
				Host:   l[0],
				User:   l[1],
				Passwd: l[2],
			}
		}
		InfoList = append(InfoList, info)
	}
}
