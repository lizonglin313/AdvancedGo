package api

import (
	"AdvancedGo/with_jianyu/go_gin_example/models"
	"AdvancedGo/with_jianyu/go_gin_example/pkg/e"
	"AdvancedGo/with_jianyu/go_gin_example/pkg/logging"
	"AdvancedGo/with_jianyu/go_gin_example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// GetAuth
// @Desc: 	通过用户名密码生成token
// @Param:	c
// @Notice:
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{
		Username: username,
		Password: password,
	}

	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if ok {
		// 是否存在这条数据
		isExist := models.CheckAuth(username, password)
		if isExist {
			// 存在,为其生成token
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			// 不存在这个数据
			code = e.ERROR_AUTH
		}
	} else {
		// 参数错误
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
