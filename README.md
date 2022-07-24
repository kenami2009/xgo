##go简易框架

---

###基础功能

####
路由、上下文、中间件、gorm、redis、日志、命令行工具

###使用指南

```
git clone https://github.com/kenami2009/xgo.git

go mod tidy

go test ./...

go run ./main.go app start
```

###数据库迁移
```
//创建迁移文件
xgo db create migrate
//执行迁移
xgo db migrate
```

