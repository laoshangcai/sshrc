package cmd

import (
	"github.com/spf13/cobra"
	"sync"
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy local files to remote hosts",
	Long: `sshrc copy --host 192.168.1.10 --passwd your-server-passwor --src /opt/kube.tar.gz --dest /opt`,
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		c := BuildInit()
		for _, host := range c.Host {
			wg.Add(1)
			go func(host string) {
				defer wg.Done()
				c.Copy(host)
			}(host)
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.Flags().StringSliceVar(&SSHConfig.Host, "host", []string{}, "servers host for ssh")
	copyCmd.Flags().StringVar(&SSHConfig.User, "user", "root", "servers user name for ssh")
	copyCmd.Flags().StringVar(&SSHConfig.Password, "passwd", "", "password for ssh")
	copyCmd.Flags().StringVar(&SSHConfig.PkFile, "pk", "/root/.ssh/id_rsa", "private key for ssh")
	copyCmd.Flags().StringVar(&SourcePath, "src",  "", "Specifies which file to pull from the managed host")
	copyCmd.Flags().StringVar(&RemotePath, "dest", "", "Specifies the location where the file will be stored after the file is pulled locally")
}