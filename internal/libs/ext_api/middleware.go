package ext_api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	// RequestIDHeader stores HTTP-header name to pass the requestID.
	RequestIDHeader = "x-request-id"
	rps             = 10
	timeout         = time.Second * 5
)

func RPSLimiter() gin.HandlerFunc {
	requests := make(chan time.Time, 10)

	go func() {
		for t := range time.Tick(time.Second / rps) {
			requests <- t
		}
	}()

	return func(c *gin.Context) {
		select {
		case <-requests:
			c.Next()
		case <-time.After(timeout):
			c.AbortWithError(http.StatusGatewayTimeout, errors.New("timeout waiting"))
		default:
			c.AbortWithError(http.StatusTooManyRequests, errors.New("too many requests"))
		}
	}
}
