package utils

import (
	"testing"
	"os"
	"bufio"
	"fmt"
)

type Cfg struct {
	filepath string                         //your ini file path directory+file

}

type InfoC struct {
	Host string
	User string
	Passwd string
}

var InfoMap = make(map[string]InfoC, 0)

func TestReadfile(t *testing.T) {
	ReadLineFile("D:/awesomeProject/src/lsc.com/sshrc/hosts")
	fmt.Println(InfoMap)
}


func ReadLineFile(fileName string) {
	if file, err := os.Open(fileName);err !=nil{
		panic(err)
	}else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan(){
			info := InfoC {
				Host: scanner.Text(),
				User: scanner.Text(),
				Passwd: scanner.Text(),
			}
			InfoMap[info.Host] = info
		}
	}
}


