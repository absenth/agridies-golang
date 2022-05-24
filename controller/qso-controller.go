package controller

import (
	"agridies-golang/entity"
	"agridies-golang/service"
	"github.com/gin-gonic/gin"
)

type QsoController interface {
	Findall() []entity.Qso
	Save(ctx *gin.Context)
}

type controller struct {
	service service.QsoService
}

func New(service service.QsoService) QsoController {
	return controller{
		service: service,
	}
}
