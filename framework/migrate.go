package framework

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

var migrationTable = "schema_migrations"

type version int64

func initDefaultDb() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_test")

	if err != nil {
		log.Panic(err)
	}
	return db
}

func GenerateMigrateSql(name string) {
	migrateDir := filepath.Join(GetExecDirectory(), "db/migrate")
	newName := time.Now().Format("20060102150405")
	migrateFileName := fmt.Sprintf("%v_%s.sql", newName, name)
	migrateFile := filepath.Join(migrateDir, migrateFileName)
	if !Exists(migrateFile) {
		if _, err := os.OpenFile(migrateFile, os.O_APPEND|os.O_CREATE, 0777); err != nil {
			log.Fatalln(err)
		}
	}
}

//初始化迁移版本表
func InitSchemaMigrations() {
	db := initDefaultDb()

	sql := `-- auto-generated definition
create table schema_migrations
(
    version varchar(260) not null,
    constraint schema_migrations_version_key
        unique (version)
)
    charset = utf8;

`
	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("create schema_migrations success.")
	}
}

func MigrateDb() {
	db := initDefaultDb()
	defer db.Close()
	versions := findNotActiveVersion()
	fmt.Println(versions)
	migrateDir := filepath.Join(GetExecDirectory(), "db/migrate")
	for _, v := range versions {
		migrateFile := filepath.Join(migrateDir, v)
		if !Exists(migrateFile) {
			continue
		}
		fmt.Println(versions)
		b, err := os.ReadFile(migrateFile)
		if err != nil {
			fmt.Println("迁移文件读取失败", err)
		}
		_, err = db.Exec(string(b))
		if err != nil {
			fmt.Println("迁移文件执行错误", err, migrateFile)
		} else {
			_, err = db.Exec("INSERT INTO schema_migrations (version) value (?)", v[:14])
			if err != nil {
				fmt.Println("迁移版本入库失败", err)
			}
			fmt.Println("迁移成功", migrateFile)
		}
	}

}
func findNotActiveVersion() []string {
	var notActives []string
	versions := getVersionFromMigrateDir()
	db := initDefaultDb()
	defer db.Close()
	fmt.Println(versions)
	for _, v := range versions {
		var count int
		row := db.QueryRow("select count(*) as count from "+migrationTable+" where version= ?", v[:14])
		if row.Scan(&count) != nil {
			continue
		} else {
			notActives = append(notActives, v)
		}
	}
	fmt.Println(notActives)
	return notActives
}

//获取所有迁移文件版本号
func getVersionFromMigrateDir() []string {
	var files []string

	migrateDir := filepath.Join(GetExecDirectory(), "db/migrate")

	err := filepath.Walk(migrateDir, func(filePath string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if path.Ext(filePath) == ".sql" {
				//files = append(files, strings.TrimSuffix(info.Name(), ".sql"))
				files = append(files, info.Name())
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return files
}
