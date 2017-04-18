package controllers

import (
	"./db/entities"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) ParseJson(data entities.BaseEntity) string {
	return ""
}

func (this *BaseController) ParseXml(data entities.BaseEntity) string {
	return ""
}

func (this *BaseController) GainDataFromJson(jsonStr string) entities.BaseEntity {
	entity := entities.BaseEntity{}
	return entity
}

func (this *BaseController) GainDataFromXml(jsonStr string) entities.BaseEntity {
	entity := entities.BaseEntity{}
	return entity
}
