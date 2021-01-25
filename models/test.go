package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

/**
 *
 *@author MaiShuRen
 *@date 2020/6/19 09:19
 *@version v1.0
 */
type Payment struct {
	Id         int
	Name       string
	Price      float64
	CreateTime time.Time
	Serial     string
}

const (
	//格式化时间类型
	timeFormat = "2006-01-02 15:04:05"
)

func init() {
	orm.RegisterDataBase("default", "mysql",
		"root:mai1208142397@(127.0.0.1:3306)/test?charset=utf8")
	orm.RegisterModel(new(Payment))
	orm.RunSyncdb("default", false, true)
}
func main() {
	orm.Debug = true
	o := orm.NewOrm()
	var qs orm.QuerySeter
	var users []Payment
	qs = o.QueryTable("payment")
	c, err := qs.All(&users)
	if err != nil {
		logs.Error(err.Error())
	}
	if c > 0 {
		logs.Info("查出数据")
		jsonUsers, _ := json.Marshal(users)
		fmt.Println(string(jsonUsers))
		for _, user := range users {
			fmt.Println(user)
		}
	} else {
		logs.Info("未查出数据")
	}
}
