package flags

import (
	"os"
	"v2ray-manager/app/cmd/serve"
	"v2ray-manager/pkg/helpers"

	"github.com/spf13/cobra"
)

// RegisterGlobalFlags 注册全局选项（flag）
func RegisterGlobalFlags(rootCmd *cobra.Command) {

}

// 注册其他命令的flag
func RegisterCommandFlags() {
	serve.CmdServe.Flags().BoolVarP(&serve.RunServerDaemon, "daemon", "d", false, "run server as a daemon")
}

// RegisterDefaultCmd 注册默认命令
func RegisterDefaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	firtArg := helpers.FirstElement(os.Args[1:])
	if err == nil && cmd.Use == rootCmd.Use && firtArg != "-h" && firtArg != "--help" {
		args := append([]string{subCmd.Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
}
