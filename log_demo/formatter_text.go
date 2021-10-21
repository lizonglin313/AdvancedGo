package log_demo

import (
	"fmt"
	"time"
)

type TextFormatter struct {
	IgnoreBasicFields bool
}

// Format
// @Desc: 	根据格式化要求对 对日志格式化输出成 Text 形式
// @Rece:	f
// @Param:	e
// @Return:	error
// @Notice:
func (f *TextFormatter) Format(e *Entry) error {

	// 如果没有忽略基础字段
	if !f.IgnoreBasicFields {
		// 写日志时间 和 日志等级
		e.Buffer.WriteString(fmt.Sprintf("%s %s", e.Time.Format(time.RFC3339), LevelNameMapping[e.Level]))
		if e.File != "" {
			short := e.File
			// 截取路径最后的文件名
			for i := len(e.File) - 1; i > 0; i-- {
				if e.File[i] == '/' {
					short = e.File[i+1:]
					break
				}
			}
			// 写 日志定位的 文件名 和 行数
			e.Buffer.WriteString(fmt.Sprintf("%s:%d", short, e.Line))
		}
		e.Buffer.WriteString(" ")
	}

	// 根据格式写日志参数
	switch e.Format {
	case FmtEmptySeparate:
		e.Buffer.WriteString(fmt.Sprint(e.Args...))
	default:
		e.Buffer.WriteString(fmt.Sprintf(e.Format, e.Args...))
	}
	e.Buffer.WriteString("\n")

	return nil
}
