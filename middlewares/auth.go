package middlewares

import (
	"fmt"
	// "log"
	// "os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/ntnghiatn/rest-go-gin-api/utils"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		data, err := utils.ParseToken(tokenString)
		if err != nil {
			fmt.Println("Err: ", err)
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		if data != nil { // có thể parse data tại đây để check tiệp
			context.Next()
			return
		}
		context.JSON(401, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
}

// Override time value for tests.  Restore default value after.
func at(t time.Time, fn func()) {
	// fmt.Println(t)
	jwt.TimeFunc = func() time.Time {
		return t
	}
	fn()
	jwt.TimeFunc = time.Now
}
