package routers

import (
	"AdvancedGo/with_jianyu/go_gin_example/middleware/jwt"
	"AdvancedGo/with_jianyu/go_gin_example/pkg/setting"
	"AdvancedGo/with_jianyu/go_gin_example/routers/api"
	v1 "AdvancedGo/with_jianyu/go_gin_example/routers/api/v1"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func InitRouter() *gin.Engine {

	// 把日志同时写到文件和终端
	logFilePath := setting.LogFilePath + time.Now().Format("20060102") + ".log"
	// 创建日志文件
	f, err := os.Create(logFilePath)
	// 以追加形式写入
	// f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic("create log file error!")
	}
	// 设置日志写入 Writer
	// gin.DefaultWriter = io.MultiWriter(f)	// 输出目标为文件 f
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) // 同时向 文件f 和 终端写入

	r := gin.New()

	// r.Use(gin.Logger())
	// 自定义日志
	// 自定义日志格式中间件
	// 将日志写入默认的 os.Stdout
	r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \" %s \" %s \"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	// 为什么单独把 GetAuth 拿出来？
	// 因为这一步是获取 token 的，就不用经过鉴权的 JWT 中间件去处理了
	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT()) // 仅在 apiv1 组里使用
	{
		apiv1.GET("/tags", v1.GetTags)          // 获取标签列表
		apiv1.POST("/tags", v1.AddTag)          // 新建标签
		apiv1.PUT("/tags/:id", v1.EditTag)      // 更新指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag) // 删除指定标签

		apiv1.GET("/articles", v1.GetArticles)          // 获取文章列表
		apiv1.GET("/articles/:id", v1.GetArticle)       // 获取指定文章
		apiv1.POST("/articles", v1.AddArticle)          // 新建文章
		apiv1.PUT("/articles/:id", v1.EditArticle)      // 修改指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle) // 删除指定文章
	}

	return r
}
