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

func TestReadfile(t *testing.T) {
	ReadLineFile("D:/awesomeProject/src/lsc.com/sshrc/hosts")
}


func ReadLineFile(fileName string) {
	if file, err := os.Open(fileName);err !=nil{
		panic(err)
	}else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan(){
			fmt.Println(scanner.Text())
		}
	}
}


