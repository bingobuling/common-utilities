//author xinbing
//time 2018/10/10 10:02
//
package http_utils

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GetClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	if ip == "" {
		RemoteAddr := ctx.Request.Header.Get("Remote_addr")
		if RemoteAddr == "" {
			addr := strings.Split(ctx.Request.RemoteAddr, ":")
			return addr[0]
		} else {
			return RemoteAddr
		}
	}
	return ip
}