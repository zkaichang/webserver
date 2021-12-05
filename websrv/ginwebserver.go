package websrv

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zkaichang/webserver/websrv/getinfo"
)

func WebServer() {

	ipaddr := getinfo.GetIpAddress()
	hostname := getinfo.Gethostname()
	version := "v1.0"
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "This is sample web.\tversion: %s\tIP address is: %s\tHostname is：%s\n", version, ipaddr, hostname)
	})
	r.GET("/readness", func(he *gin.Context) {
		he.String(http.StatusOK, "This web server is ready!")
	})
	r.GET("/healthy", func(he *gin.Context) {
		he.String(http.StatusOK, "This web server is healthy!")
	})
	r.Run(":8888")

}
