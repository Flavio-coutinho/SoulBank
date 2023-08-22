package users

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// TODO: Criar usuário
func create(ctx *gin.Context) {
	saltRounds := 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ctx.Request.Body.Password), saltRounds)
	if err != nil {
		ctx.Status(500)
		ctx.String("Algo deu errado, tente novamente")
		return
	}
	user, err := prisma.User.CreateOne(prisma.User.Email.Equals(ctx.Request.Body.Email) && prisma.User.Password.Equals(hashedPassword)).Exec(ctx)
	if err != nil {
		ctx.Status(500)
		ctx.String("Algo deu errado, tente novamente")
		return
	}
	ctx.JSON(200, user)
}