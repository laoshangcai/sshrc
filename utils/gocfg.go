package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Config struct {
	filepath string                         //your ini file path directory+file
	Conflist map[string]map[string]string //configuration information slice
}

//Create an empty configuration file
func InitConfig(filepath string) *Config {
	c := new(Config)
	c.filepath = filepath
	c.readList()
	return c
}

func (c *Config) GetValue(section, name string) string {
	_,ok := c.Conflist[section][name]
	if ok{
		return c.Conflist[section][name]
	}else{
		return ""
	}
}

//获取所有配置项
//List all the configuration file
func (c *Config) readList() map[string]map[string]string {
	file, err := os.Open(c.filepath)
	if err != nil {
		CheckErr(err)
	}
	defer file.Close()
	c.Conflist = make(map[string]map[string]string)
	var section string
	var sectionMap map[string]string
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
				c.Conflist[section] = sectionMap
			}else{
				isFirstSection = false
			}
			section = strings.TrimSpace(line[1 : len(line)-1])
			sectionMap = make(map[string]string)
		default:
			i := strings.IndexAny(line, "=")
			if i == -1 {
				continue
			}
			value := strings.TrimSpace(line[i+1 : len(line)])
			sectionMap[strings.TrimSpace(line[0:i])] = value
		}
	}
	c.Conflist[section] = sectionMap
	return c.Conflist
}

func CheckErr(err error) string {
	if err != nil {
		return fmt.Sprintf("Error is :'%s'", err.Error())
	}
	return "Notfound this error"
}

