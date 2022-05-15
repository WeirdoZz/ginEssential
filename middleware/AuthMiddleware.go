package middleware

import (
	"ginEssential/common"
	"ginEssential/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			context.JSON(http.StatusUnauthorized, gin.H{"msg": "not authorized,token string err"})
			context.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"msg": "not authorized,parse failed"})
			context.Abort()
			return
		}

		// get user id from token
		userId := claims.UserId
		db := common.GetDB()
		var user model.User
		db.First(&user, userId)

		if userId == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{"msg": "not authorized,user not exist"})
			context.Abort()
			return
		}

		context.Set("user", user)
		context.Next()
	}
}
