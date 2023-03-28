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
	requests := make(chan struct{}, rps)

	for i := 0; i < rps; i++ {
		select {
		case requests <- struct{}{}:
		default:
		}
	}

	ticker := time.NewTicker(time.Second)
	go func() {
		for _ = range ticker.C {
			for i := 0; i < rps; i++ {
				select {
				case requests <- struct{}{}:
				default:
				}
			}
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
