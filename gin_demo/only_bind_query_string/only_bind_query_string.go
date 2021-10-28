package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Person struct {
	Name string `form:"name"`
	Address string `form:"address"`
}

func startPage(c *gin.Context)  {
	var person Person
	// ShouldBindQuery is a shortcut for c.ShouldBindWith(obj, binding.Query)
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}

func main() {
	route := gin.Default()
	// Any 适用于任何请求方法 包括但不限于 GET POST...
	route.Any("/testing", startPage)
	route.Run(":8123")
}
