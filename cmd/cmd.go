package cmd

import (
	"github.com/wonderivan/logger"
)


func Cmd(host string, user, passwd string, cmd string) []byte {
	logger.Info("[%s]exec cmd is : %s", host, cmd)
	session, err := Connect(user, passwd, "", host)
	defer func() {
		if r := recover(); r != nil {
			logger.Error("[%s]Error create ssh session failed,%s", host, err)
		}
	}()
	if err != nil {
		panic(1)
	}
	defer session.Close()

	b, err := session.CombinedOutput(cmd)
	logger.Debug("[%s]command result is:\n %s", host, string(b))
	defer func() {
		if r := recover(); r != nil {
			logger.Error("[%s]Error exec command failed: %s", host, err)
		}
	}()
	if err != nil {
		panic(1)
	}

	return b
}

