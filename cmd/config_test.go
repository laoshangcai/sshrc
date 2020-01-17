package cmd

import (
	"testing"
	"fmt"
	"strings"
    "lsc.com/sshrc/utils"
)

type Configinfo struct {
	 Host string
	 User string
	 Passwd string
}

var InfoList = make([]Configinfo, 0)
var InfoMap = make(map[string]Configinfo, 0)


func Getvalue(st string) string {
	it := strings.Split(st," ")
	for _, v := range it {
		result := strings.Index(v,",")
		if result == -1 {
			fmt.Println("格式不对")
		} else {
			i := strings.Split(v,",")

			info := Configinfo {
				Host: i[0],
				User: i[1],
				Passwd: i[2],
			}
			InfoList = append(InfoList,info)
			InfoMap[info.Host] = info
		}
	}
	return ""
}


func TestConfig(t *testing.T) {
	var conf= utils.InitConfig("D:/awesomeProject/src/lsc.com/sshrc/hosts.ini")
	value := conf.GetValue("server1", "=")
	fmt.Println(value)
	//Getvalue(value)
	//fmt.Println(InfoList)
	//fmt.Println(InfoMap)
}