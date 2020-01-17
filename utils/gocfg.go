package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Confi struct {
	filepath string
	Conflist map[string]map[string]string
}

//Create an empty configuration file
func InitConfig(filepath string) *Confi {
	c := new(Confi)
	c.filepath = filepath
	c.readList()
	return c
}

func (c *Confi) GetValue(section, name string) string {
	_,ok := c.Conflist[section][name]
	if ok{
		return c.Conflist[section][name]
	}else{
		return ""
	}
}

//获取所有配置项
//List all the configuration file
func (c *Confi) readList() map[string]map[string]string {
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
		fmt.Println(string(line))
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
			fmt.Println(line)
			i := strings.IndexAny(line, ",")
			if i == -1 {
				continue
			}
			value := strings.TrimSpace(line[i : ])
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

