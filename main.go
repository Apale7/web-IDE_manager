package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetReportCaller(true)
	r := gin.Default()
	defer r.Run(":3070")
	collectRoutes(r)
}
