package utils

import (
	"testing"
	"os"
	"bufio"
	"fmt"
	"io"
	"strings"
)


var group List

type List struct {
	Server string
}
type Config struct {
	Filepath string
	configMap map[string]string
}


func TestRead(t *testing.T) {
	var conf= InitConfi("D:/awesomeProject/src/lsc.com/sshrc/hosts")
	value := conf.GetValue("server1")
	fmt.Println(value)

}

func InitConfi(filepath string) *Config {
	c := new(Config)
	c.Filepath = filepath
	c.ReadLineFile()
	return c
}

func (c *Config) GetValue(section string) string {
	_,ok := c.configMap[section]
	if ok{
		return c.configMap[section]
	}else{
		return ""
	}
}

func (c *Config) ReadLineFile() map[string]string {
	file, err := os.Open(c.Filepath)
	if err != nil {
		CheckErr(err)
	}
	defer file.Close()
	c.configMap = make(map[string]string)
	isFirstSection := true
	buf := bufio.NewReader(file)

	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				CheckErr(err)
			}
			if len(line) == 0 {
				break
			}
		}
		switch {
		case len(line) == 0:
		case string(line[0]) == "#":	//增加配置文件备注
		case line[0] == '[' && line[len(line)-1] == ']':
			if !isFirstSection{
				continue
			}else{
				isFirstSection = false
			}
			section := strings.TrimSpace(line[1 : len(line)-1])
			group = List {Server: section,}
		default:
			c.configMap[group.Server] = line
		}
	}
	return c.configMap

}


