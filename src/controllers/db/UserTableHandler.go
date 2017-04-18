package db

import (
	"./entities"
	"log"
	"encoding/json"
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

func SelectAllUser() string {
	sql := "select id,name,gender,degree from MyUsers"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	//判断err是否有错误的数据，有err数据就显示panic的数据
	defer rows.Close()

	rows.ColumnTypes()

	result := "{"
	for rows.Next() {
		var user entities.UserEntity
		rerr := rows.Scan(&user.Id, &user.Name, &user.Gender, &user.Degree) //数据指针，会把得到的数据，往刚才id 和 lvs引入

		content, _ := json.MarshalIndent(user, "", "  ")

		str := string(content[:])
		//if rows.Next() {
		result += str + ",\n"
		//}

		if rerr != nil {
			log.Fatal(rerr)
			continue
		}
	}
	result += "}"
	return result
}
