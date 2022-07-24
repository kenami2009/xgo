package console

import (
	"github.com/spf13/cobra"
	"log"
	"xgo/framework"
	"xgo/framework/middleware"
)

var appAddr = ":8000"

var AppCommand = &cobra.Command{
	Use:     "app",
	Short:   "app控制台",
	Long:    "app控制台命令[start|stop|state|restart]",
	Example: "xgo app start",
	RunE: func(c *cobra.Command, args []string) error {
		log.Println("app控制台")
		return nil
	},
}

func init() {
	AppCommand.AddCommand(stopAppCommand)
	AppCommand.AddCommand(startAppCommand)
	AppCommand.AddCommand(stateAppCommand)
	AppCommand.AddCommand(restartAppCommand)
}

//启动App
var startAppCommand = &cobra.Command{
	Use:   "start",
	Short: "app启动",
	RunE: func(c *cobra.Command, args []string) error {
		var x = framework.NewEngine()

		x.Use(middleware.Recovery())
		x.Use(middleware.Logger())

		framework.InitDb()
		framework.InitRedis()

		x.GET("/", framework.IndexController)
		log.Fatalln(x.Run(":8000"))
		return nil
	},
}

//停止App
var stopAppCommand = &cobra.Command{
	Use:   "stop",
	Short: "app停止",
	RunE: func(c *cobra.Command, args []string) error {
		log.Println("app stop")
		return nil
	},
}

//App状态
var stateAppCommand = &cobra.Command{
	Use:   "state",
	Short: "app状态",
	RunE: func(c *cobra.Command, args []string) error {
		log.Println("app state")
		return nil
	},
}

//重启App服务
var restartAppCommand = &cobra.Command{
	Use:   "restart",
	Short: "app重启",
	RunE: func(c *cobra.Command, args []string) error {
		log.Println("app restart")
		return nil
	},
}
