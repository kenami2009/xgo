package framework

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-ini/ini"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

var levelFlag = map[string]LogLevel{
	"DEBUG": DebugLevel,
	"INFO":  InfoLevel,
	"PANIC": PanicLevel,
	"WARM":  WarnLevel,
	"ERROR": ErrorLevel,
	"TRACE": TraceLevel,
	"FATAL": FatalLevel,
}

type Logger struct {
	level      LogLevel
	output     io.Writer
	formatter  Formatter
	ctxfielder CtxFielder
}

var _ ILogger = &Logger{}

func NewLogger() *Logger {
	conf, err := ini.Load("./framework/config.ini")
	if err != nil {
		log.Fatal("配置文件获取错误")
	}
	log_level := conf.Section("").Key("log_level").String()
	log_path := conf.Section("").Key("log_path").String()
	log_file := filepath.Join(GetExecDirectory(), log_path)
	var out *os.File
	if !Exists(log_file) {
		if out, err = os.OpenFile(log_file, os.O_APPEND|os.O_CREATE, 0777); err != nil {
			log.Fatalln(err)
		}
	} else {
		//output
		out, err = os.OpenFile(log_file, os.O_APPEND|os.O_CREATE, 0777)
		if err != nil {
			log.Fatalln()
		}
	}

	level, ok := levelFlag[log_level]
	if !ok {
		log.Fatal("日志类型有误")
	}

	if err != nil {
		log.Fatal("日志文件创建错误", err)
	}
	return &Logger{
		level:  level,
		output: out,
	}
}

func Prefix(level LogLevel) string {
	prefix := ""
	switch level {
	case PanicLevel:
		prefix = "[Panic]"
	case FatalLevel:
		prefix = "[Fatal]"
	case ErrorLevel:
		prefix = "[Error]"
	case WarnLevel:
		prefix = "[Warn]"
	case InfoLevel:
		prefix = "[Info]"
	case DebugLevel:
		prefix = "[Debug]"
	case TraceLevel:
		prefix = "[Trace]"
	}
	return prefix
}

//默认格式化方法
func defaultLogFormatter(level LogLevel, t time.Time, fields map[string]interface{}, msg ...string) ([]byte, error) {
	bf := bytes.NewBuffer([]byte{})
	Separator := "|"
	// 先输出日志级别
	prefix := Prefix(level)

	bf.WriteString(prefix)
	// 输出时间
	ts := t.Format(time.RFC3339)
	bf.WriteString(ts)
	bf.WriteString(Separator)

	// 请求信息TODO
	//bf.WriteString(fmt.Sprint(fields))

	// 输出msg
	for i, m := range msg {
		if i > 0 {
			bf.WriteString(", ")
		}
		bf.WriteString(m)
	}

	return bf.Bytes(), nil
}

//打印日志
func (l *Logger) logf(loglever LogLevel, ctx context.Context, msg ...string) (err error) {
	//日志级别

	//ctx上下文
	fs := make(map[string]interface{})
	if l.ctxfielder != nil {
		ctx_fields := l.ctxfielder(ctx)
		for k, v := range ctx_fields {
			fs[k] = v
		}
	}
	//格式化
	if l.formatter == nil {
		l.formatter = defaultLogFormatter
	}
	r, err := l.formatter(loglever, time.Now(), fs, msg...)
	if err != nil {
		return
	}
	//输出
	_, err = l.output.Write(r)
	if err != nil {
		fmt.Println(err)
	}
	l.output.Write([]byte("\r\n"))
	return
}
func (l *Logger) SetLevel(level LogLevel) {
	if l.level != level {
		l.level = level
	}
}

func (l *Logger) SetFormatter(formatter Formatter) {
	l.formatter = formatter
}
func (l *Logger) SetOutput(out io.Writer) {
	l.output = out
}
func (l *Logger) SetCtxFielder(handler CtxFielder) {

}
func (l *Logger) Panic(c context.Context, msg ...string) {
	l.logf(PanicLevel, c, msg...)
}
func (l *Logger) Fatal(c context.Context, msg ...string) {
	l.logf(FatalLevel, c, msg...)
}
func (l *Logger) Error(c context.Context, msg ...string) {
	l.logf(ErrorLevel, c, msg...)
}
func (l *Logger) Warn(c context.Context, msg ...string) {
	l.logf(WarnLevel, c, msg...)
}
func (l *Logger) Info(c context.Context, msg ...string) {
	l.logf(InfoLevel, c, msg...)
}
func (l *Logger) Debug(c context.Context, msg ...string) {
	l.logf(DebugLevel, c, msg...)
}
func (l *Logger) Trace(c context.Context, msg ...string) {
	l.logf(TraceLevel, c, msg...)
}
