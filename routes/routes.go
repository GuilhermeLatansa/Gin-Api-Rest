package routes

import (
	"gin-api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET(":nome", controllers.Saudacoes)
	r.POST("/alunos", controllers.CriaNovoAluno)

	r.Run()
}
