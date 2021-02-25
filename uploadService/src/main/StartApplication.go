package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router :=gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"pong",
		})
	})
	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	router.GET("/user/:name", func(c *gin.Context) {
		name :=c.Param("name")
		c.String(http.StatusOK,"hello %s",name)
	})

	// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
	// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name :=c.Param("name")
		action :=c.Param("action")
		message :=name + " is " + action
		c.String(http.StatusOK,message)
	})

	//获取Get参数
	router.GET("/welcome", func(c *gin.Context) {
		firstName :=c.DefaultQuery("firstname","guest")

		lastName :=c.Query("lastname")//是 c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK,"hello %s %s",firstName,lastName)
	})

	//获取post参数
	router.POST("/form_post", func(c *gin.Context) {
		message :=c.PostForm("message")
		nick := c.DefaultPostForm("nick","anonymous")//此方法可以设置默认值

		c.JSON(200,gin.H{
			"status":"posted",
			"message": message,
			"nick": nick,
		})
	})

	//上传文件(单个文件)
	router.POST("/upload/single", func(c *gin.Context) {
		file,_ :=c.FormFile("file")
		log.Print(file.Filename)

		// 上传文件到指定的路径
		// c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK,fmt.Sprintf("'%s' uploaded!",file.Filename))
	})

	router.POST("/upload/multi", func(c *gin.Context) {
		form,_ := c.MultipartForm()
		files := form.File["upload"]

		for _,file := range files{
			log.Println(file.Filename)
			// 上传文件到指定的路径
			// c.SaveUploadedFile(file, dst)
		}

		c.String(http.StatusOK,fmt.Sprintf("%d files upload!", len(files)))
	})




	router.Run()
}
