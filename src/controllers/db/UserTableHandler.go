package db

import (
	"./protobuf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type UsersTableController struct {
	BaseTableController
}

func (this *UsersTableController) InsertUser(data *protobuf.User) bool {
	sql := fmt.Sprintf("insert into %s(%s,%s,%s) value (?,?,?)", this.TableName, this.Indexs[1], this.Indexs[2], this.Indexs[3])
	stmt := GetStmt(sql)

	_, err := stmt.Exec(data.Name, data.Gender, data.Degree)

	if err != nil {
		return false
	} else {
		return true
	}
}

func (this *UsersTableController) UpdateUser(data *protobuf.User) {
	sql := fmt.Sprintf("update %s set %s=?,%s=?,%s=? where %s=?", this.TableName, this.Indexs[1], this.Indexs[2], this.Indexs[3], this.Indexs[0])
	stmt := GetStmt(sql)

	stmt.Exec(data.GetName(), data.GetGender(), data.GetDegree(), data.GetId())
	//for _, entity := range data {
	//	stmt.Exec(entity.GetUserName(), entity.GetId())
	//}
}

func (this *UsersTableController) DeleteUser(data *protobuf.User) {
	sql := fmt.Sprintf("delete from %s where %s=?", this.TableName, this.Indexs[0])
	stmt := GetStmt(sql)
	stmt.Exec(data.GetId())
}

func (this *UsersTableController) SelectUser(id int64) *protobuf.User {
	//sql := fmt.Sprintf("select %s,%s,%s,%s from %s where %s=?", this.Indexs[0], this.Indexs[1], this.Indexs[2], this.Indexs[3], this.TableName, this.Indexs[0])
	sql := fmt.Sprintf("select * from %s where %s=?", this.TableName, this.Indexs[0])
	var user protobuf.User

	row := db.QueryRow(sql, id)
	row.Scan(&user.Id, &user.Name, &user.Gender, &user.Degree)
	return &user
}
