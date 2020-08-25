/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"sync"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "fetch the remote host file",
	Long: `sshrc fetch --host 192.168.1.10 --passwd your-server-passwor --src /opt/kube.tar.gz --dest /opt`,
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		c := BuildInit()
		for _, host := range c.Host {
			wg.Add(1)
			go func(host string) {
				defer wg.Done()
				c.Fetch(host)
			}(host)
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().StringSliceVar(&SSHConfig.Host, "host", []string{}, "servers host for ssh")
	fetchCmd.Flags().StringVar(&SSHConfig.User, "user", "root", "servers user name for ssh")
	fetchCmd.Flags().StringVar(&SSHConfig.Password, "passwd", "", "password for ssh")
	fetchCmd.Flags().StringVar(&SSHConfig.PkFile, "pk", "/root/.ssh/id_rsa", "private key for ssh")
	fetchCmd.Flags().StringVar(&SourcePath, "src",  "", "The remote file path")
	fetchCmd.Flags().StringVar(&DestPath, "dest", "", "The local file path")
}
