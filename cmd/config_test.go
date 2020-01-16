package cmd

import (
	"testing"
	"fmt"
	"strings"
    "lsc.com/sshrc/utils"
)

type HostConfig struct {
	 Host string
	 User string
	 Passwd string
}

func Getvalue(st string) string {
	it := strings.Split(st," ")
	for _, v := range it {
		result := strings.Index(v,",")
		if result == -1 {
			fmt.Println(v)
		} else {
			i := strings.Split(v,",")
			for _, s := range i {
				fmt.Println(s)
			}
		}
	}
	return ""
}


func TestConfig(t *testing.T) {
	var conf= utils.InitConfig("D:/awesomeProject/src/lsc.com/sshrc/hosts.ini")
	value := conf.GetValue("server1", "host")
	Getvalue(value)
	
}