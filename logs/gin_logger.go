//author bing.xin
//time  2018/10/24 10:45
//desc gin logrus 集成
package logs

import (
	"common-utilities/http_utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func GinLogrusLogger() gin.HandlerFunc {
	skip := make(map[string]string)
	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery
		ctx.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			// Stop timer
			end := time.Now()
			latency := end.Sub(start)
			clientIP := http_utils.GetClientIP(ctx)
			method := ctx.Request.Method
			statusCode := ctx.Writer.Status()
			if raw != "" {
				path = path + "?" + raw
			}
			logrus.Infoln(fmt.Sprintf("[GIN] %v | %3d | %13v | %15s |%s %-7s",
				end.Format("2006/01/02 - 15:04:05"),
				statusCode,
				latency,
				clientIP,
				method,
				path,
			))
		}
	}
}
