package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	Version    = "v1.1.0"
	Build      = ""
	VersionStr = fmt.Sprintf("sshrc version %v, build %v %v", Version, Build, runtime.Version())
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show sshrc version",
	Long: `show sshrc version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VersionStr)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
