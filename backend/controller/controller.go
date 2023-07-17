package controller

import (
	"golang_task/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Return the negative images
// @Tags images
// @Accept  json
// @Produce json
// @Param image body model.InputImage true "The image in format base64"
// @Success 201 {string} model.InputImage
// @Failure 500
// @Router /negative_image [post]
func NegativeImage(context *gin.Context) {
	var input model.InputImage
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	entity, err := input.CreateEntity()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	saved_entity, err := entity.SaveEntity()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"data": saved_entity})
}

// @Summary Return three last images
// @Tags images
// @Accept  json
// @Produce json
// @Success 201 {object} model.StoredImage
// @Failure 500
// @Router /get_last_images [get]
func GetLastImages(context *gin.Context) {
	ret, err := model.GetLastThree()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusCreated, gin.H{"data": ret})
}
