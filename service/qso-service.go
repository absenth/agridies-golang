package service

import "agridies-golang/entity"

type QsoService interface {
	Save(entity.Qso) entity.Qso
	FindAll() []entity.Qso
}

type qsoService struct {
	qsos []entity.Qso
}

func New() QsoService {
	return &qsoService{}
}

func (service *qsoService) Save(qso entity.Qso) entity.Qso {
	service.qsos = append(service.qsos, qso)
	return qso
}

func (service *qsoService) FindAll() []entity.Qso {
	return service.qsos
}
