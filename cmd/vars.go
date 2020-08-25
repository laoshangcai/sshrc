package cmd

import "time"

type SSHRConfig struct {
	//sshconfig
	Host     []string
	User     string
	Password string
	PkFile   string
	Timeout  *time.Duration
}

type SSH struct {
	Host       []string
	User     string
	Password string
	PkFile   string
}

var (
	SSHConfig  SSH
	// shell cmd
	CMDS       string

	// cpoy path
	SourcePath string
	RemotePath string
	// fetch ; localpath
	DestPath   string
)
