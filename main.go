package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	port string
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	port = os.Getenv("web_ide_manager_port")
	if port == "" {
		port = "3070"
	}
}

func main() {
	logrus.SetReportCaller(true)
	r := gin.Default()
	defer r.Run(":" + port)
	collectRoutes(r)
}
