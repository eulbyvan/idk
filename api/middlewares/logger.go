/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Fri May 12 2023 1:29:18 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		fmt.Printf("%s %s %s %s\n", c.Request.Method, c.Request.URL.Path, c.ClientIP(), duration)
	}
}
