package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {

	router := gin.Default()

	// 返回 JSON
	// gin.H 实际上就是 map[string]interface{}
	router.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hey",
			"status": http.StatusOK,
		})
	})

	// struct 形式返回 JSON
	router.GET("/struct_json", func(c *gin.Context) {
		var msg struct{
			Name string `json:"user"`
			Message string `json:"msg"`
			Number int
		}
		msg.Name = "lzl"
		msg.Message = "a message"
		msg.Number = 100
		c.JSON(http.StatusOK, msg)
	})

	// 安全的 防止劫持的 json
	// while(1);["lena","austin","foo"]
	router.GET("/secure_json", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		c.SecureJSON(http.StatusOK, names)
	})

	// jsonp callback模式
	router.GET("/jsonp", func(c *gin.Context) {
		data := gin.H{
			"foo": "bar",
		}
		c.JSONP(http.StatusOK, data)
	})

	// 对 json 数据中的 非 ascii 进行转义
	// {"name":"go\u8bed\u8a00","tag":"\u003cbr\u003e"}
	router.GET("/ascii_json", func(c *gin.Context) {
		data := gin.H{
			"name": "go语言",
			"tag": "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	// 保持原始的 数据 不对 Unicode  进行编码
	// Serves unicode entities
	// {"html":"\u003cb\u003eHello, world!\u003c/b\u003e"}
	router.GET("/unicode_json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	// Serves literal characters
	// {"html":"<b>Hello, world!</b>"}
	router.GET("/pure_json", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 返回 xml
	// <map>
	// <message>hey xml</message>
	// <status>200</status>
	// </map>
	router.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"message": "hey xml",
			"status": http.StatusOK,
		})
	})

	// 返回 yaml
	// message: hey yaml
	// status: 200
	router.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"message": "hey yaml",
			"status": http.StatusOK,
		})
	})

	// 返回 protobuf
	//
	router.GET("/protobuf", func(c *gin.Context) {
		resp := []int64{int64(1), int64(2)}
		label := "test"

		data := &protoexample.Test{
			Label: &label,
			Reps: resp,
		}
		// 响应数据以二进制的形式传递
		// 序列化输出
		c.ProtoBuf(http.StatusOK, data)
	})

	router.Run(":8123")
}
