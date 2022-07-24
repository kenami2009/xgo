# go-start-kit框架

---

本框架主要是仿gin框架，并结合了主流PHP框架使用经验，旨在打造一个开箱即用的简介web框架。

## 基础功能

* 路由：使用字典树（trie）实现
* 上下文：
```go
type XContext struct {
	index          int //中间件计数器
	handlers       ControllerHandlerChain //中间件管道
	request        *http.Request //http请求
	responseWriter http.ResponseWriter //http响应
	ctx            context.Context  //上下文
	logger         *Logger //日志
}
```
* 中间件

```go
var x = framework.NewEngine()

//使用全局Recovery中间件
x.Use(middleware.Recovery())
```
* gorm
* redis
* 日志
* 命令行工具(cobra)

## 后续功能
* 缓存服务
* 持续集成
* 进程守护
* 前后端分离（Vue）


## 使用指南

```
git clone https://github.com/kenami2009/xgo.git

go mod tidy

go test ./...

go run ./main.go app start
```

## 数据库迁移
```
//创建迁移文件
xgo db create migrate
//执行迁移
xgo db migrate
```

