package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("开始执行")
		c.Set("request", "中间件")
		c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time: ", t2)
	}
}

type Login struct {
	User     string `json:"user" form:"user" uri:"user" binding:"required"`
	Password string `json:"password" form:"password" uri:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/api/v1/measure/:user/:password", MiddleWare(), func(c *gin.Context) {
		var json Login
		err := c.ShouldBindUri(&json)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.User != "xize" || json.Password != "123456" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "用户名或密码错误",
				"status":  "304",
			})
			return
		}

		request, _ := c.Get("request")
		fmt.Println("request: ", request)

		var result struct {
			Status  int
			Message string
		}
		result.Status = 200
		result.Message = "OK"

		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "notset"
			c.SetCookie("key_cookie", "value_cookie", 60, "/", "localhost", false, true)
		}
		fmt.Println("cookie: ", cookie)
		fmt.Println("hello")
		c.JSON(http.StatusOK, result)
	})

	err := r.Run(":8003")
	if err != nil {
		return
	}
}
