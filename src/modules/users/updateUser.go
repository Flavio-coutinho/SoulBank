package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: Actualizar usuário
func update(ctx *gin.Context) {
	var user User
	ctx.BindJSON(&user)
	id := ctx.Param("id")
	_, err := db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, id)
	if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Algo deu errado, tente novamente")
			return
	}
	ctx.JSON(http.StatusOK, user)
}