package log_demo

import (
	"fmt"
	"io"
	"os"
	"sync"
	"unsafe"
)

var std = New()

type logger struct {
	opt       *options		// 日志选项：包括日志等级、输出目标
	mu        sync.Mutex	// 加锁

	// 保存和复用临时对象，减少内存分配，降低 GC 压力
	// 可伸缩的 并发安全的
	// 仅受制于内存大小
	// 用于存储那些被分配了但是没有被使用的，而且未来可能会使用的值
	// 存放在对象池中的对象如果不活跃了会被自动清理
	entryPool *sync.Pool
}

// New
// @Desc: 	选项模式 根据传入的选项 构造 logger
// @Param:	opts
// @Return:	*logger
// @Notice:
func New(opts ...Option) *logger {
	// 构造选项
	logger := &logger{
		opt: initOptions(opts...),
	}
	// 在对象池创建一个新的对象
	logger.entryPool = &sync.Pool{
		// 这个 New() 属性定义一个 返回任意接口的匿名函数即可
		New: func() interface{} {
			return entry(logger)
		},
	}
	return logger
}

func StdLogger() *logger {
	return std
}

// 设置参数
func SetOptions(opts ...Option) {
	std.SetOptions(opts...)
}

func (l *logger) SetOptions(opts ...Option)  {
	// 上锁
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, opt := range opts {
		opt(l.opt)
	}
}

func Writer() io.Writer {
	return std
}

func (l *logger) Writer() io.Writer {
	return l
}

// logger 实现 io.Write interface
func (l *logger) Write(data []byte) (int, error) {
	l.entry().write(l.opt.stdLevel, FmtEmptySeparate, *(*string)(unsafe.Pointer(&data)))
	return 0, nil
}

func (l *logger) entry() *Entry {
	// 从 对象池中 取出一个对象 注意使用断言 因为出来的对象统一是 interface
	return l.entryPool.Get().(*Entry)
}

func (l *logger) Debug(args ...interface{}) {
	l.entry().write(DebugLevel, FmtEmptySeparate, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.entry().write(InfoLevel, FmtEmptySeparate, args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.entry().write(WarnLevel, FmtEmptySeparate, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.entry().write(ErrorLevel, FmtEmptySeparate, args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.entry().write(PanicLevel, FmtEmptySeparate, args...)
	panic(fmt.Sprint(args...))
}

func (l *logger) Fatal(args ...interface{}) {
	l.entry().write(FatalLevel, FmtEmptySeparate, args...)
	os.Exit(1)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.entry().write(DebugLevel, format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.entry().write(InfoLevel, format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.entry().write(WarnLevel, format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.entry().write(ErrorLevel, format, args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.entry().write(PanicLevel, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.entry().write(FatalLevel, format, args...)
	os.Exit(1)
}

// std logger
func Debug(args ...interface{}) {
	std.entry().write(DebugLevel, FmtEmptySeparate, args...)
}

func Info(args ...interface{}) {
	std.entry().write(InfoLevel, FmtEmptySeparate, args...)
}

func Warn(args ...interface{}) {
	std.entry().write(WarnLevel, FmtEmptySeparate, args...)
}

func Error(args ...interface{}) {
	std.entry().write(ErrorLevel, FmtEmptySeparate, args...)
}

func Panic(args ...interface{}) {
	std.entry().write(PanicLevel, FmtEmptySeparate, args...)
	panic(fmt.Sprint(args...))
}

func Fatal(args ...interface{}) {
	std.entry().write(FatalLevel, FmtEmptySeparate, args...)
	os.Exit(1)
}

func Debugf(format string, args ...interface{}) {
	std.entry().write(DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	std.entry().write(InfoLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	std.entry().write(WarnLevel, format, args...)
}

func Errorf(format string, args ...interface{}) {
	std.entry().write(ErrorLevel, format, args...)
}

func Panicf(format string, args ...interface{}) {
	std.entry().write(PanicLevel, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func Fatalf(format string, args ...interface{}) {
	std.entry().write(FatalLevel, format, args...)
	os.Exit(1)
}