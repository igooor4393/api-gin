package routes

import (
	"api-gin/controllers"
	"github.com/gin-gonic/gin"
)

type PostRouteController struct {
	postController controllers.PostController
}

func NewRoutePostController(postController controllers.PostController) PostRouteController {
	return PostRouteController{postController}
}

func (pc *PostRouteController) PostRoute(rg *gin.RouterGroup) {

	// Достучаться до хендлеров можно через host:port/api/posts[ / ; uuid ]
	router := rg.Group("posts")
	router.POST("/", pc.postController.CreateProgram)
	router.GET("/:postId", pc.postController.FindProgramById)

	// ------------------------ Просто для примера ------------------------
	router.GET("/", pc.postController.FindProgram)
	router.PUT("/:postId", pc.postController.UpdateProgram)
	router.DELETE("/:postId", pc.postController.DeletePost)
}
