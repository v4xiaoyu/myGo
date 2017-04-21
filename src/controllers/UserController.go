package controllers

import (
	"./db"
	"strconv"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {

	s := this.Ctx.Request.URL.String()
	v := getParams(s)

	//this.GainData()
	this.Ctx.WriteString("enter UserController")

	id, _ := strconv.ParseInt(v["id"][0], 10, 64)

	user := db.SelectUser(id)

	//jsonString, err := json.Marshal(user)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}

	this.Ctx.WriteString("\n")
	this.Ctx.Output.JSON(user, true, true)

	//this.Ctx.WriteString("\n")
	//user.Name = "dan"
	//user.Gender = 2
	//user.Degree = 8
	//db.UpdateUser(user)
	//
	//
	//user2 := db.SelectUser(3)
	//
	//this.Ctx.WriteString(fmt.Sprintf("name is : %s\ngender is : %d\ndegree is: %f",user2.Name,user2.Gender,user2.Degree))
}
