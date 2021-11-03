package jwt

import (
	"AdvancedGo/with_jianyu/go_gin_example/pkg/e"
	"AdvancedGo/with_jianyu/go_gin_example/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// JWT
// @Desc: 	JWT中间件,从URL中拿token并鉴权
// @Return:	gin.HandlerFunc
// @Notice:
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			// 没有token
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				// 解析token失败
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				log.Printf("token %s don't have auth", token)
			} else if time.Now().Unix() > claims.ExpiresAt {
				// 如果token过期
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		// 如果鉴权失败就不再继续做了，直接截断本次的请求
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		// 鉴权成功过这个中间件
		c.Next()
	}
}
