package controllers

import (
	"api-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(DB *gorm.DB) PostController {
	return PostController{DB}
}

func (pc *PostController) CreateProgram(ctx *gin.Context) {

	payload := models.CreatePostRequest{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var programs []models.Program

	for _, program := range payload.Items {
		newProgram := models.Program{
			Id:        program.Id,
			Name:      program.Name,
			NameEn:    program.NameEn,
			IsPublic:  program.IsPublic,
			ProjectID: program.ProjectID,
		}
		programs = append(programs, newProgram)
	}

	result := pc.DB.CreateInBatches(programs, 100) // в идеале кол-во бачей передавать откуда ни будь аргументом
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Post with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success"})
}

func (pc *PostController) UpdateProgram(ctx *gin.Context) {
}

func (pc *PostController) FindProgramById(ctx *gin.Context) {}

func (pc *PostController) FindProgram(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var posts []models.Program
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&posts)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(posts), "data": posts})
}

func (pc *PostController) DeletePost(ctx *gin.Context) {

}
