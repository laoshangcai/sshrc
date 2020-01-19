package cmd

import (
	"testing"
	"strings"
    "lsc.com/sshrc/utils"
	"fmt"
)

type Configin struct {
	 Host string
	 User string
	 Passwd string
}

var InfoLis = make([]Configin, 0)
var inf = Configin{}

func TestGetcf(t *testing.T) {
	var conf= utils.InitConfig("D:/awesomeProject/src/lsc.com/sshrc/hosts")
	value := conf.GetValue("server2")

	for i := 0; i < len(value); i++ {
		result := strings.Index(value[i], ",")
		if result == -1 {
			inf = Configin{Host: value[i]}
		} else {
			l := strings.Split(value[i], ",")
			inf = Configin{
				Host:   l[0],
				User:   l[1],
				Passwd: l[2],
			}
		}
		InfoLis = append(InfoLis, inf)
	}
	fmt.Println(InfoLis)
}