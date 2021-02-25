package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	loger "webproject/src/log"
)
//func init() {
//	//log.SetFormatter(&log.JSONFormatter{})
//	log.SetFormatter(&log.TextFormatter{})
//
//	log.SetOutput(os.Stdout)
//	log.SetLevel(log.InfoLevel)
//}
func main() {
	//simpleGinForLog()
	//logrus()
	ginForLogToFile()
}



//func logrus(){
//	log.WithFields(log.Fields{
//		"name":"yh",
//	}).Info("info hh")
//
//	log.WithFields(log.Fields{
//		"company":"yestae",
//	}).Debug("debug hh")
//
//}

//ginForLog 存文件
func ginForLogToFile()  {
	// 禁用控制台颜色
	gin.DisableConsoleColor()
	router :=gin.Default()
	router.Use(loger.LoggerToFile())
	router.GET("/q", func(context *gin.Context) {
		//rand.Seed(time.Now().Unix())
		randomString :=strconv.Itoa(rand.Intn(100))
		//Info级别的日志
		loger.Logger().WithFields(logrus.Fields{
			"name": randomString,
		}).Info("记录一下日志", "Info")
	})
	router.Run()
}


//简单的ginForLog
//func simpleGinForLog()  {
//	f,_:=os.Create("gin.log")
//	gin.DefaultWriter = io.MultiWriter(f,os.Stdout)
//
//	rand.Seed(time.Now().Unix())
//	router :=gin.Default()
//	router.GET("/index", func(c *gin.Context) {
//		randomString :=strconv.Itoa(rand.Intn(100))
//		c.String(200,randomString)
//		fmt.Println("rand:",randomString)
//	})
//	router.Run(":8080")
//}

//简单的gin例子
//func simpleGin()  {
//	r := gin.Default()
//	rand.Seed(time.Now().Unix())
//	r.GET("/ping", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": rand.Intn(100),
//		})
//	})
//	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
//}

func exampleX()  {
	router:=gin.Default()

	get := router.Group("/get")
	{
		get.GET("/someGet",GetDataD)
	}

	router.GET("someGet",getting())
	router.POST("somePost",posting())
	router.PUT("somePut",putting())
	router.DELETE("someDelete",deleting())

	router.Run(":8080")

}

func GetDataD(c *gin.Context) {

}


func getting() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}
func posting() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}
func putting() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}
func deleting() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}
