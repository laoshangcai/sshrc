package utils

import (
	"os"
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Config struct {
	filepath string
	configMap map[string][]string
}

func InitConfig(filepath string) *Config {
	c := new(Config)
	c.filepath = filepath
	c.ReadLineFile()
	return c
}

func (c *Config) GetValue(section string) []string  {
	_,ok := c.configMap[section]
	if ok{
		return c.configMap[section]
	}else{
		return nil
	}
}

func (c *Config) ReadLineFile() map[string][]string {
	file, err := os.Open(c.filepath)
	if err != nil {
		CheckErr(err)
	}
	defer file.Close()

	c.configMap = make(map[string][]string)
	var section string

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
			section = strings.TrimSpace(line[1 : len(line)-1])
		default:
			c.configMap[section] = append(c.configMap[section],line)
		}
	}
	return c.configMap
}

func CheckErr(err error) string {
	if err != nil {
		return fmt.Sprintf("Error is :'%s'", err.Error())
	}
	return "Notfound this error"
}