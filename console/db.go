package console

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"xgo/framework"
)

var DbCommand = &cobra.Command{
	Use:     "db",
	Short:   "数据库迁移",
	Long:    "数据库迁移",
	Example: "migrate",
	RunE: func(c *cobra.Command, args []string) error {
		log.Println("migrate控制台")
		return nil
	},
}

func init() {
	DbCommand.AddCommand(migrateCommand)
	DbCommand.AddCommand(generateCommand)
	DbCommand.AddCommand(initCommand)
}

var initCommand = &cobra.Command{
	Use:     "init",
	Short:   "初始化数据表",
	Long:    "初始化数据表",
	Example: "db init",
	RunE: func(c *cobra.Command, args []string) error {
		framework.InitSchemaMigrations()
		return nil
	},
}
var migrateCommand = &cobra.Command{
	Use:     "migrate",
	Short:   "迁移",
	Long:    "迁移命令",
	Example: "db generate",
	RunE: func(c *cobra.Command, args []string) error {
		framework.MigrateDb()
		return nil
	},
}

var generateCommand = &cobra.Command{
	Use:     "generate",
	Short:   "创建迁移文件",
	Long:    "创建迁移文件",
	Example: "db generate create_table_users",
	RunE: func(c *cobra.Command, args []string) error {
		fmt.Println(args)
		framework.GenerateMigrateSql(args[0])
		return nil
	},
}
