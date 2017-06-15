package db

import "../protobuf"

var newsController *NewsTableController = &NewsTableController{BaseTableController{"NewsDetail",
	[]string{"news_id", "title", "description", "content", "images", "group_id"}}}
var userController *UsersTableController = &UsersTableController{BaseTableController{"UserInfo",
	[]string{"user_id", "name", "gender", "description", "avatar", "age"}}}

func InsertNews(data *protobuf.News) bool {
	return newsController.InsertNews(data)
}

func SelectNews(lastId int64, size int) string {
	return newsController.SelectNews(lastId, size)
}

//////////////////////////////////////------------------
func InsertUser(data *protobuf.User) bool {
	return userController.InsertUser(data)
}

func SelectUser(id int64) *protobuf.User {
	return userController.SelectUser(id)
}
