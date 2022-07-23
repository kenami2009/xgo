package main

import (
	"github.com/spf13/cobra"
	"xgo/console"
)

var command = &cobra.Command{
	Use:   "xgo",
	Short: "xgo 命令",
	// 根命令的详细介绍
	Long: "xgo 框架提供的命令行工具", // 根命令的执行函数
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.InitDefaultHelpFlag()
		return cmd.Help()
	},
	// 不需要出现 cobra 默认的 completion 子命令
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

func main() {
	command.AddCommand(console.AppCommand)
	command.Execute()
}
