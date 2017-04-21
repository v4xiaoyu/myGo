package db

import (
	"./protobuf"
	"log"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type NewsTableController struct {
	BaseTableController
}

func (this *NewsTableController) InsertNews(data *protobuf.News) bool {
	sql := fmt.Sprintf("insert into %s(%s,%s, %s) value (?,?,?)", this.TableName, this.Indexs[1], this.Indexs[2], this.Indexs[3])

	_, err := insert(sql, data.Title, data.Desc, data.Content)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (this *NewsTableController) UpdateNews(data *protobuf.News) {
	sql := fmt.Sprintf("update %s set %s=?,%s=?,%s=? where %s=?", this.TableName, this.Indexs[1], this.Indexs[2], this.Indexs[3], this.Indexs[0])
	exec(sql, data.Title, data.Desc, data.GetId())
}

func (this *NewsTableController) DeleteNews(data *protobuf.News) {
	sql := fmt.Sprintf("delete from %s where %s=?", this.TableName, this.Indexs[0])
	exec(sql, data.GetId())
}

func (this *NewsTableController) SelectNews(lastId int64, size int) string {
	//sql := fmt.Sprintf("select %s,%s,%s,%s from %s limit %d,%d", this.Indexs[0], this.Indexs[1], this.Indexs[2], this.Indexs[3], this.TableName, lastId, size)

	sql := fmt.Sprintf("select * from %s limit %d,%d", this.TableName, lastId, size)
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	//判断err是否有错误的数据，有err数据就显示panic的数据
	defer rows.Close()

	rows.ColumnTypes()
	result := "{"
	var count int = 0
	for rows.Next() {
		var news protobuf.News
		rerr := rows.Scan(&news.Id, &news.Title, &news.Desc, &news.Content) //数据指针，会把得到的数据，往刚才id 和 lvs引入

		content, _ := json.MarshalIndent(news, "", "  ")

		str := string(content[:])
		//if rows.Next() {
		result += str + ",\n"
		//}

		if rerr != nil {
			log.Fatal(rerr)
			continue
		}
		count++
	}
	if count < size {
		size = count
	}
	result += fmt.Sprintf("last_id:%d}", lastId+int64(size))
	return result
}
