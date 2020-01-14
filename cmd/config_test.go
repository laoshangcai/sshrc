package cmd

import (
	"testing"
	"fmt"
	goini "lsc.com/sshrc/utils"
)

func TestConfig(t *testing.T) {
	var conf = goini.InitConfig("D:/awesomeProject/src/lsc.com/sshrc/hosts.ini")
	//var conf = goini.InitConfig("./app.ini")

	value := conf.GetValue("hosts","host")

	fmt.Println(value)

}