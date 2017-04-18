package controllers

import (
	"./db"
	"./db/entities"
)

type UserController struct {
	BaseController

	Id     int64
	Name   string
	Gender int8
}

func (this *UserController) GainData() {
	this.Id = 12
	this.Name = "laffey"
	this.Gender = 1

	data := &entities.UserEntity{entities.BaseEntity{11}, "laffey", 0, 4}
	db.InsertUser(data)
}

func (this *UserController) Get() {
	this.GainData()
	this.Ctx.WriteString("enter UserController")

	user := db.SelectUser(1)

	//jsonString, err := json.Marshal(user)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}

	this.Ctx.WriteString("\n")
	this.Ctx.Output.JSON(user, true, true)

	this.Ctx.WriteString("\n")
	this.Ctx.WriteString(db.SelectAllUser())
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
