package main

import "github.com/gin-gonic/gin"

func prepareRoutes(router *gin.Engine) {
	// Web Resources
	router.Static("/static", "web/dist")

	// API Routes
	admin := api.Group("/admin")
	qso := qso.Group("/qso")

	prepareMiddleware(admin, qso)

	admin.GET("/operator", listOperator)
	admin.POST("/operator", createOperator)
	admin.PUT("/operator", updateOperator)

	qso.GET("/:id", getQso)
	qso.POST("/", createQSO)
	qso.PUT("/:id", updateQSO)
	qso.GET("/", getQsos)
	qso.DELETE("/:id", deleteQSO)
}

func main() {
	r := prepareRoutes()
	r.Run(":8080")
}
