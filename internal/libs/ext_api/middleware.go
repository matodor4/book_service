package ext_api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	rps = 10
)

func RPSLimiter() gin.HandlerFunc {
	requests := make(chan time.Time, rps)

	go func() {
		for t := range time.Tick(time.Second / rps) {

			requests <- t
		}
	}()

	return func(c *gin.Context) {
		select {
		case <-requests:
			c.Next()
		default:
			c.AbortWithError(http.StatusTooManyRequests, errors.New("too many requests"))
		}
	}
}
