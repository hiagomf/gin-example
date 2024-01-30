package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hiagomf/gin-example/application/user/base"
)

func findAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"method": "GET_FIND_ALL",
	})
}

func findByID(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	newUserID, err := uuid.Parse(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	response, err := base.FindByID(ctx, &newUserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, response)
}

func create(ctx *gin.Context) {
	var request base.CreateRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	response, err := base.Create(ctx, &request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, response)
}

func update(ctx *gin.Context) {
	var request base.UpdateRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	userID := ctx.Param("user_id")
	newUserID, err := uuid.Parse(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	request.ID = &newUserID

	if err := base.Update(ctx, &request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func delete(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	newUserID, err := uuid.Parse(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := base.Delete(ctx, &newUserID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusNoContent, nil)
}
