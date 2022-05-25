package controller

import (
	"agridies-golang/entity"
	"agridies-golang/service"
	"github.com/gin-gonic/gin"
)

type QsoController interface {
	Findall() []entity.Qso
	Save(ctx *gin.Context) entity.Qso
}

type controller struct {
	service service.QsoService
}

func New(service service.QsoService) QsoController {
	return controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Qso {
	return c.service.FindAll()
}

func (c *contorller) Save(ctx *gin.Context) entity.Qso {
	var qso entity.Qso
	ctx.BindJSON(&qso)
	c.service.Save(qso)
	return qso
}
