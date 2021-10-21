package log_demo

import (
	"bytes"
	"runtime"
	"strings"
	"time"
)

// 一个日志实体
type Entry struct {
	logger *logger

	// 字符串输出的缓冲区
	// json输出的对map marshal 后的缓冲区
	Buffer *bytes.Buffer

	Map    map[string]interface{} // 该字段主要是格式化输出成Json格式所用
	Level  Level
	Time   time.Time
	File   string
	Line   int
	Func   string
	Format string
	Args   []interface{}
}

// entry
// @Desc: 	通过logger构造一个Entry
// @Param:	logger
// @Return:	*Entry
// @Notice:
func entry(logger *logger) *Entry {
	return &Entry{
		logger: logger,
		Buffer: new(bytes.Buffer),
		Map:    make(map[string]interface{}, 5), // 对应 options 中的 5 个选项
	}
}

func (e *Entry) write(level Level, format string, args ...interface{})  {
	if e.logger.opt.level > level {
		return
	}

	e.Time = time.Now()
	e.Level = level
	e.Format = format
	e.Args = args

	if !e.logger.opt.disableCaller {
		// Caller(层数) 表示文件栈的深度
		// 0 代表当前层
		// 1 代表上一次调用者
		// 2 3 4 依此类推
		if pc, file, line, ok := runtime.Caller(2); !ok {
			e.File="???"
			e.Func="???"
		} else {
			e.File, e.Line, e.Func = file, line, runtime.FuncForPC(pc).Name()
			e.Func = e.Func[strings.LastIndex(e.Func, "/")+1:]
		}
	}

	e.format()
	e.writer()
	// 把 log 数据写完之后 释放掉 再存入 对象池中即可
	e.release()
}

// 根据 Format 格式化写到缓存buffer中
func (e *Entry) format() {
	_ = e.logger.opt.formatter.Format(e)
}

// 然后按照输出目标 从缓存中 输出到 目标
func (e *Entry) writer() {
	e.logger.mu.Lock()
	_, _ = e.logger.opt.output.Write(e.Buffer.Bytes())
	e.logger.mu.Unlock()
}

// 对象用完后 放回对象池
func (e *Entry) release() {
	e.Args, e.Line, e.File, e.Format, e.Func = nil, 0, "", "", ""
	e.Buffer.Reset()
	// 放回池子
	e.logger.entryPool.Put(e)
}
