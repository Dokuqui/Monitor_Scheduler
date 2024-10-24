package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	UserRoutes(r)
	AdminRoutes(r)
	ManagerRoutes(r)

	return r
}
