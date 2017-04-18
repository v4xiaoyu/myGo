package controllers

import (
	"fmt"
)

type NewsController struct {
	BaseController

	Title string
	Idx   int64
	Desc  string
}

func (this *NewsController) GainData() {
	this.Title = "trerst"
	this.Idx = 11
	this.Desc = "3213fsaf"
}

func (this *NewsController) Get() {
	this.GainData()
	this.Ctx.WriteString(fmt.Sprintf("title : %s\nid : %d\ndesc : %s", this.Title, this.Idx, this.Desc))
}
