package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: Deletar usuário
func remove(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Algo deu errado, tente novamente")
			return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}