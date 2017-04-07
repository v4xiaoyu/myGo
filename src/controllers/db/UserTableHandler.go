package db

import (
	"./entities"
	"fmt"
)

type UsersTableController struct {
	BaseTableController
}

func (this *UsersTableController) InsertUser(data *entities.UserEntity) bool {
	sql := fmt.Sprintf("insert into %s(%s,%s,%s) value (?,?,?)", this.TableName, this.Indexs[1], this.Indexs[2], this.Indexs[3])
	stmt := GetStmt(sql)

	_, err := stmt.Exec(data.Name, data.Gender, data.Degree)

	if err != nil {
		return false
	} else {
		return true
	}
}

func (this *UsersTableController) UpdateUser(data *entities.UserEntity) {
	sql := fmt.Sprintf("update %s set %s=?,%s=?,%s=? where %s=?", this.TableName, this.Indexs[1], this.Indexs[2], this.Indexs[3], this.Indexs[0])
	stmt := GetStmt(sql)

	stmt.Exec(data.Name, data.Gender, data.Degree, data.GetId())
	//for _, entity := range data {
	//	stmt.Exec(entity.GetUserName(), entity.GetId())
	//}
}

func (this *UsersTableController) DeleteUser(data *entities.UserEntity) {
	sql := fmt.Sprintf("delete from %s where %s=?", this.TableName, this.Indexs[0])
	stmt := GetStmt(sql)
	stmt.Exec(data.GetId())
}

func (this *UsersTableController) SelectUser(id int64) *entities.UserEntity {
	sql := fmt.Sprintf("select %s,%s,%s,%s from %s where %s=?", this.Indexs[0], this.Indexs[1], this.Indexs[2], this.Indexs[3], this.TableName, this.Indexs[0])
	var user entities.UserEntity

	row := db.QueryRow(sql, id)
	row.Scan(&user.Id, &user.Name, &user.Gender, &user.Degree)
	return &user
}
