package users

import "github.com/gin-gonic/gin"

// TODO: Listar usuários
func list(ctx *gin.Context) {
	users, err := prisma.User.FindMany().Exec(ctx)
	if err != nil {
		ctx.Status(500)
		ctx.String("Algo deu errado, tente novamente")
		return
	}
	ctx.JSON(200, users)
}