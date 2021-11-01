package util

import (
	"AdvancedGo/with_jianyu/go_gin_example/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetPage
// @Desc: 	从URL中拿到页数
//			根据设置的每页显示多少文返回数量
//			设置分页页码
// @Param:	c
// @Return:	int
// @Notice:
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
