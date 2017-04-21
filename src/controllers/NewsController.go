package controllers

import (
	"./db"
	"github.com/bitly/go-simplejson"
	"encoding/json"
)

type NewsController struct {
	BaseController
}

func (this *NewsController) Get() {
	jsoninfo := this.GetString("data")
	var id int64
	var ierr error
	if jsoninfo == "" {
		this.Ctx.WriteString("data is empty")
		id = 0
	} else {

		data, err := json.Marshal(jsoninfo)
		if err != nil {
			return
		}
		jsonObj, jerr := simplejson.NewJson(data)
		if jerr != nil {
			return
		}
		this.Ctx.WriteString("\n")
		id, ierr = jsonObj.Get("id").Int64()
		if ierr != nil {
			id = 0
		}
	}
	this.Ctx.WriteString(db.SelectNews(id, 10))
}
