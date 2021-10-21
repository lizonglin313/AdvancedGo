package log_demo

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)


const FmtEmptySeparate = ""

type Level uint8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

var LevelNameMapping = map[Level]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	PanicLevel: "PANIC",
	FatalLevel: "FATAL",
}

var errUnmarshalNilLevel = errors.New("can't unmarshal a nil *Level")

// 将 log 可读等级 映射成 数字： DEBUG -> 0
func (l *Level) unmarshalText(text []byte) bool {
	switch string(text) {
	case "debug", "DEBUG":
		*l = DebugLevel
	case "info", "INFO":
		*l = InfoLevel
	case "warn", "WARN":
		*l = WarnLevel
	case "error", "ERROR":
		*l = ErrorLevel
	case "panic", "PANIC":
		*l = PanicLevel
	case "fatal", "FATAL":
		*l = FatalLevel
	default:
		return false
	}
	return true
}

// 对 unmarshalText 再次封装
func (l *Level) UnmarshalText(text []byte) error {
	if l == nil {
		return errUnmarshalNilLevel
	}
	if !l.unmarshalText(text) || !l.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("unrecognized level: %q", text)
	}
	return nil
}

// 日志选项
type options struct {
	output   io.Writer // 输出目标
	level    Level
	stdLevel Level

	//	type Formatter interface {
	//	Format(entry *Entry) error
	//}
	formatter Formatter

	disableCaller bool
}

type Option func(*options)

// 如何使用?
// initOptions
// @Desc: 	设置日志选项
// @Param:	opts
// @Return:	o
// @Notice:	这里的匿名函数的使用可以概况为这样的场景下
//			1. 知道要向一个对象的属性赋值
//			2. 这个值是确定的
//			3. 不知道具体向哪个对象赋值
//			4. 使用匿名函数，使其带着值过来
//			此外：
//			函数作为一等公民，作为参数传入
func initOptions(opts ...Option) (o *options) {
	o = &options{}
	for _, opt := range opts {
		opt(o)
	}

	if o.output == nil {
		o.output = os.Stderr
	}

	if o.formatter == nil {
		o.formatter = &TextFormatter{}
	}

	return
}

func WithOutput(output io.Writer) Option {
	return func(o *options) {
		o.output = output
	}
}

func WithLevel(level Level) Option {
	return func(o *options) {
		o.level = level
	}
}

func WithStdLevel(level Level) Option {
	return func(o *options) {
		o.stdLevel = level
	}
}

func WithFormatter(formatter Formatter) Option {
	return func(o *options) {
		o.formatter = formatter
	}
}

func WithDisableCaller(caller bool) Option {
	return func(o *options) {
		o.disableCaller = caller
	}
}
