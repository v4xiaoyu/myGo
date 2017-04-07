package db

import "./entities"

var newsController *NewsTableController = &NewsTableController{BaseTableController{"NewsDetail",
										   []string{"id", "title", "description", "content"}}}
var userController *UsersTableController = &UsersTableController{BaseTableController{"MyUsers",
										     []string{"id", "name", "gender", "degree"}}}

func InsertNews(data *entities.NewsEntity) bool {
	return newsController.InsertNews(data)
}

func SelectNews(lastId int64, size int) string {
	return newsController.SelectNews(lastId, size)
}

//////////////////////////////////////------------------
func InsertUser(data *entities.UserEntity) bool {
	return userController.InsertUser(data)
}

func SelectUser(id int64) *entities.UserEntity {
	return userController.SelectUser(id)
}
