package main

import (
	"os"
	"v2ray-manager/app/cmd/serve"
	"v2ray-manager/app/runtime/app"
	"v2ray-manager/app/runtime/flags"
	"v2ray-manager/bootstrap"
	"v2ray-manager/config"
	configPkg "v2ray-manager/pkg/config"

	"github.com/spf13/cobra"
)

func main() {
	os.Mkdir("etc", os.ModePerm)
	app.AbsolutePath, _ = os.Getwd()
	config.Initialize()

	// 主入口，调用 cmd.CmdServe 命令
	var rootCmd = &cobra.Command{
		Use:   "v2ray-manager",
		Short: "my project describe",
		Long:  `默认会使用 "serve" 命令，输入 -h 查看帮助`,

		// 定义所有注册命令都会执行的func
		PersistentPreRun: func(command *cobra.Command, args []string) {
			configPkg.InitConfig("", "")

			bootstrap.SetupLogger()
			bootstrap.SetupNodeConfig()
			bootstrap.SetupConnections()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		serve.CmdServe,
	)

	// 注册指定命令，非全局的flag
	flags.RegisterCommandFlags()

	// 注册全局参数
	flags.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
