package cmd

import (
	"fmt"
	"github.com/pkg/sftp"
	"github.com/wonderivan/logger"
	"golang.org/x/crypto/ssh"
	"net"
	"os"
	"path"
	"strconv"
	"syscall"
	"time"
)

const (
	oneKBByte = 1024
    oneMBByte = 1024 * 1024
    )

// 实例初始化，赋值
func BuildInit() *SSHRConfig {
   c := &SSHRConfig{
	   Host:     SSHConfig.Host,
	   User:     SSHConfig.User,
	   Password: SSHConfig.Password,
   }
   return c
}

//远程执行命令
func (s *SSHRConfig) Cmd(host string) []byte {
	logger.Info("[%s]exec cmd is : %s", host, CMDS)
	session, err := s.Connect(host)
	defer func() {
		if r := recover(); r != nil {
			logger.Error("[%s]Error create ssh session failed,%s", host, err)
		}
	}()

	if err != nil {
		panic(1)
	}
	defer session.Close()

	b, err := session.CombinedOutput(CMDS)
	logger.Info("[%s]command msg :\n %s", host, string(b))
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

//拷贝本地主机文件到远程主机
func (s *SSHRConfig) Copy(host string) {
	sftpClient, err := s.sftpconnect(host)
	defer func() {
		if r := recover(); r != nil {
			logger.Error("[ssh][%s]scpCopy: %s", host, err)
		}
	}()
	if err != nil {
		panic(1)
	}
	defer sftpClient.Close()

	srcFile, err := os.Open(SourcePath)   //本地
	defer func() {
		if r := recover(); r != nil {
			logger.Error("[ssh][%s]scpCopy: %s", host, err)
		}
	}()
	if err != nil {
		panic(1)
	}
	defer srcFile.Close()

	var remoteFileName = path.Base(SourcePath)
	dstFile, err := sftpClient.Create(path.Join(RemotePath, remoteFileName))  //远程
	defer func() {
		if r := recover(); r != nil {
			logger.Error("[ssh][%s]scpCopy: %s", host, err)
		}
	}()
	if err != nil {
		panic(1)
	}
	defer dstFile.Close()

	buf := make([]byte, 100*oneMBByte) //100mb
	total := 0
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		length, _ := dstFile.Write(buf[0:n])
		total += length
	}
	totalLength, totalUnit := toSizeFromInt(total)
	logger.Alert("[ssh][%s]transfer total size is: %.2f%s", host, totalLength, totalUnit)
}

//拷贝远程主机文件到本地主机
func (s *SSHRConfig) Fetch(host string) {
	sftpClient, err := s.sftpconnect(host)
	defer func() {
		if r := recover(); r != nil {
			logger.Error("[ssh][%s]scpCopy: %s", host, err)
		}
	}()
	if err != nil {
		panic(1)
	}
	defer sftpClient.Close()

	//创建sftp连接，并读取文件内容
	srcFile, err := sftpClient.Open(SourcePath)  //远程
	defer func() {
		if r := recover(); r != nil {
			logger.Error("[ssh][%s]scpCopy: %s", host, err)
		}
	}()
	if err != nil {
		panic(1)
	}
	defer srcFile.Close()

    //本地文件写入内容
	var DestfileName = (path.Join(DestPath, path.Base(SourcePath)))
	dstFile, err := os.OpenFile(DestfileName, syscall.O_CREAT|syscall.O_RDWR, 0664) //本地
	if err != nil {
		panic(1)
	}
	defer dstFile.Close()

	buf := make([]byte, 100*oneMBByte) //100mb
	total := 0
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		length, _ := dstFile.Write(buf[0:n])
		total += length
	}
	totalLength, totalUnit := toSizeFromInt(total)
	logger.Alert("fetch file [%s][%s] total size is: %.2f%s", host, SourcePath, totalLength, totalUnit)
}

// 创建sftp连接
func (s *SSHRConfig) sftpconnect(host string) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(s.Password))
	clientConfig = &ssh.ClientConfig{
		User:    s.User,
		Auth:    auth,
		Timeout: 30 * time.Second,
		//验证远程主机，以保证安全性。这里不验证
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = AddrReformat(host)
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}
	return sftpClient, nil
}

func toSizeFromInt(length int) (float64, string) {
	isMb := length/oneMBByte > 1
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(length)/oneMBByte), 64)
	if isMb {
		return value, "MB"
	} else {
		value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(length)/oneKBByte), 64)
		return value, "KB"
	}
}