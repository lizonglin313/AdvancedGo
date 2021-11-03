package routers

import (
	"AdvancedGo/with_jianyu/go_gin_example/middleware/jwt"
	"AdvancedGo/with_jianyu/go_gin_example/pkg/setting"
	"AdvancedGo/with_jianyu/go_gin_example/routers/api"
	v1 "AdvancedGo/with_jianyu/go_gin_example/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	// 为什么单独把 GetAuth 拿出来？
	// 因为这一步是获取 token 的，就不用经过鉴权的 JWT 中间件去处理了
	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())	// 仅在 apiv1 组里使用
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
