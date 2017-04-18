package db

import (
	"./entities"
)

func CreateUserTable() {

}

func InsertUser(data *entities.UserEntity) {
	sql := "insert into MyUsers(name,gender,degree) value (?,?,?)"
	stmt := GetStmt(sql)

	stmt.Exec(data.Name, data.Gender, data.Degree)
}

func UpdateUser(data *entities.UserEntity) {
	sql := "update MyUsers set name=?,gender=?,degree=? where id=?"
	stmt := GetStmt(sql)

	stmt.Exec(data.Name, data.Gender, data.Degree, data.GetId())
	//for _, entity := range data {
	//	stmt.Exec(entity.GetUserName(), entity.GetId())
	//}
}

func DeleteUser(data *entities.UserEntity) {
	sql := "delete from MyUsers where id=?"
	stmt := GetStmt(sql)
	stmt.Exec(data.GetId())
}

func SelectUser(id int64) *entities.UserEntity {
	sql := "select id,name,gender,degree from MyUsers where id=?"
	var user entities.UserEntity

	row := db.QueryRow(sql, id)
	row.Scan(&user.Id, &user.Name, &user.Gender, &user.Degree)
	return &user
}

//func SelectAllUser() {
//	sql := "select id,name,gender,degree from MyUsers"
//	rows, err := db.Query(sql, 10)
//	if err != nil {
//		panic(err)
//		log.Fatal(err)
//	}
//	//判断err是否有错误的数据，有err数据就显示panic的数据
//	defer rows.Close()
//
//	var users []entities.UserEntity
//	for rows.Next() {
//		var user entities.UserEntity
//		rerr := rows.Scan(&user.Id, &user.Name, &user.Gender, &user.Degree)  //数据指针，会把得到的数据，往刚才id 和 lvs引入
//		u
//
//	}
//	return &users
//}
