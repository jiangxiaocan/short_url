package Model

import (
	"database/sql"
	"fmt"
)

type Increase struct {
	ID int
	Name string
	Value uint64
}

//获取递增数量
func GetIncreaseNum(name string) (int64,error){
	tx, err := Db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		return 0,err
	}

	sqlStr1 := "UPDATE increase SET `value` = `value`+1 WHERE name = ?"
	_,err1 := tx.Exec(sqlStr1, name)

	if err1 != nil {
		_ = tx.Rollback()
		fmt.Printf("exec failed, err:%v\n", err)
		return 0,err
	}

	sqlStr := "select * from increase where name =  ?"
	row := tx.QueryRow(sqlStr,name)

	//声明三个变量
	var id int
	var names string
	var values int64
	//将各个字段中的值读到以上三个变量中
	err2 := row.Scan( &id, &names, & values)

	_ = tx.Commit()

	//没有数据
	if err2 == sql.ErrNoRows{
		return 0,err
	}

	return values,
		nil

}