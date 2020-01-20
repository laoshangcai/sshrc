package cmd

import (
	"time"
	"io/ioutil"
	"os/exec"
	"fmt"
	"testing"
)

func Net_ping1(host string) int {
	cmds := "ping -c2 " + host + "| grep " + host + "| wc -l"
	cmd := exec.Command("/bin/bash", "-c", cmds)
	stdout, _ := cmd.StdoutPipe()

	//执行检测命令
	err := cmd.Start()
	if err != nil {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Error：命令执行出错,", err)
	}

	//输出检测结果
	out, _ := ioutil.ReadAll(stdout)
	str := string(out)

	if str == "4" {
		return 1
	}

	stdout.Close()
	//阻塞直到该命令执行完成
	cmd.Wait()
	return -1
}


func TestPing1(t *testing.T) {
	s := Net_ping1("172.16.20.12")
	fmt.Println(s)
}
