package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sync"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sshrc",
	Short: "sshrc default implementation",
	Long: `sshrc --host 192.168.1.10 --user root --passwd your-server-password --cmd "ls -a"`,
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		c := BuildInit()
		for _, host := range c.Host {
			wg.Add(1)
			go func(host string) {
				defer wg.Done()
				c.Cmd(host)
			}(host)
		}
		wg.Wait()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringSliceVar(&SSHConfig.Host, "host", []string{}, "servers host for ssh")
	rootCmd.Flags().StringVar(&SSHConfig.User, "user", "root", "servers user name for ssh")
	rootCmd.Flags().StringVar(&SSHConfig.Password, "passwd", "", "password for ssh")
	rootCmd.Flags().StringVar(&SSHConfig.PkFile, "pk", "/root/.ssh/id_rsa", "private key for ssh")
	rootCmd.Flags().StringVar(&CMDS, "cmd", "", "shell command")
}
