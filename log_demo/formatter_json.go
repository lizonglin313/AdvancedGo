package log_demo

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"time"
)

type JsonFormatter struct {
	IgnoreBasicFields bool
}

// Format
// @Desc: 	根据格式化要求 对日志实体 进行格式化输出成 json 格式
// @Rece:	f
// @Param:	e
// @Return:	error
// @Notice:
func (f *JsonFormatter) Format(e *Entry) error {
	if !f.IgnoreBasicFields {
		e.Map["level"] = LevelNameMapping[e.Level]
		e.Map["time"] = e.Time.Format(time.RFC3339)
		if e.File != "" {
			e.Map["file"] = e.File + ":" + strconv.Itoa(e.Line)
			e.Map["func"] = e.Func
		}

		switch e.Format {
		case FmtEmptySeparate:
			e.Map["message"] = fmt.Sprint(e.Args...)
		default:
			e.Map["message"] = fmt.Sprintf(e.Format, e.Args...)
		}

		return jsoniter.NewEncoder(e.Buffer).Encode(e.Map)
	}

	switch e.Format {
	case FmtEmptySeparate:
		for _, arg := range e.Args {
			if err := jsoniter.NewEncoder(e.Buffer).Encode(arg); err != nil {
				return err
			}
		}
	default:
		e.Buffer.WriteString(fmt.Sprintf(e.Format, e.Args))
	}
	return nil
}
