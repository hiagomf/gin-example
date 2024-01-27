package contact

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func findAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"method": "GET_FIND_ALL",
	})
}

func findByID(ctx *gin.Context) {
	contactID := ctx.Param("contact_id")
	ctx.JSON(http.StatusOK, gin.H{
		"contact_id": contactID,
		"method":     "GET",
	})
}

func create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"method": "POST_CREATE",
	})
}

func update(ctx *gin.Context) {
	contactID := ctx.Param("contact_id")
	ctx.JSON(http.StatusOK, gin.H{
		"contact_id": contactID,
		"method":     "PUT",
	})
}

func delete(ctx *gin.Context) {
	contactID := ctx.Param("contact_id")
	ctx.JSON(http.StatusOK, gin.H{
		"contact_id": contactID,
		"method":     "DELETE",
	})
}
