package log

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

func Logger() *logrus.Logger {
	logFilePath :=""
	if dir,err := os.Getwd();err==nil{
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath,0777);err !=nil{
		fmt.Println(err.Error())
	}
	now :=time.Now()
	logFileName :=now.Format("2006-1-2") + ".log"
	fileName := path.Join(logFilePath,logFileName)
	if _,err:=os.Stat(fileName);err!=nil{
		if _,err := os.Create(fileName);err!=nil{
			fmt.Println(err.Error())
		}
	}
	src,err :=os.OpenFile(fileName,os.O_APPEND|os.O_WRONLY,os.ModeAppend)
	if err!=nil{
		fmt.Println("err",err)
	}

	logger :=logrus.New()
	logger.Out = src

	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})

	return logger
}

func LoggerToFile() gin.HandlerFunc {
	logger :=logrus.New()
	return func(c *gin.Context) {
		startTime :=time.Now()
		c.Next()
		endTime :=time.Now()
		//执行时间
		latencyTime :=endTime.Sub(startTime)
		//请求方法
		reqMethod := c.Request.Method
		//请求路由
		reqUri :=c.Request.RequestURI
		// 状态码
		statusCode :=c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		//日志格式
		logger.Infof("| %3d | %13v | %15s |%s | %s|",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
			)
	}
}
