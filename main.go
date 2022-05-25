package main

import (
	"agridies-golang/controller"
	"agridies-golang/service"
	"github.com/gin-gonic/gin"
)

func prepareRoutes(router *gin.Engine) {
	// Web Resources
	router.Static("/static", "web/dist")

	// API Routes
	admin := api.Group("/admin")
	qso := qso.Group("/qso")

	prepareMiddleware(admin, qso)

	admin.GET("/operator", getOperators)
	admin.GET("/operator/:id", getOperator)
	admin.POST("/operator/", createOperator)
	admin.PUT("/operator/:id", updateOperator)

	qso.GET("/", getQsos)
	qso.GET("/:id", getQso)
	qso.POST("/", createQso)
	qso.PUT("/:id", updateQso)

	qso.DELETE("/:id", deleteQso)
}

var (
	qsoService    service.QsoService       = service.New()
	qsoController controller.QsoController = controller.New(qsoService)
)

func main() {
	r := prepareRoutes()

	r.GET("/qso", func(ctx *gin.Context) {
		ctx.JSON(200, qsoController.Findall())
	})

	r.POST("/qso", func(ctx *gin.Context) {
		ctx.JSON(201, qsoController.Save(ctx))
	})

	r.Run(":8080")
}
