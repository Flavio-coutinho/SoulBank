package users

import (
	"os"
	"/src/modules/users/services.go"
	"/src/data/data.go"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func login(ctx *gin.Context) {
	email, password := services.DecodeBasicToken(ctx.Request.Header.Get("Authorization"))
	user, err := prisma.User.FindUnique(prisma.User.Email.Equals(email) && prisma.User.Password.Equals(password)).Exec(ctx)
	if err != nil {
		ctx.Status(404)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		ctx.Status(500)
		ctx.String("Algo deu errado, tente novamente", err)
		return
	}
	ctx.JSON(200, gin.H{
		"user":  user,
		"token": tokenString,
	})
}