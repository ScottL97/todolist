package controllers

import (
	"fmt"
	"time"
	"todolist/models"
	"todolist/utils"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	db := beego.AppConfig.String("mongodb::mongo_db")
	coll_items := beego.AppConfig.String("mongodb::mongo_coll_items")
	coll_history := beego.AppConfig.String("mongodb::mongo_coll_history")

	var res []models.Item
	err := models.FindAll(db, coll_items, nil, &res)
	if err != nil {
		fmt.Println("error to FindAll:", err)
	}
	c.Data["items"] = res
	c.Data["count"] = len(res)
	// 获取总共拖延时长
	var delay int64 = 0
	var diff int64 = 0
	for _, item := range res {
		dt, err := time.ParseInLocation("2006-01-02 15:04:05", item.Deadline, time.Local)
		if err != nil {
			fmt.Println("error to parse deadline:", err)
			continue
		}
		diff = time.Now().Unix() - dt.Unix()
		if diff > 0 {
			delay += diff
		}
	}
	delayDay := delay / (3600 * 24)
	delayHour := delay % (3600 * 24) / 3600
	delayMinute := delay % 3600 / 60
	c.Data["delayDay"] = delayDay
	c.Data["delayHour"] = delayHour
	c.Data["delayMinute"] = delayMinute
	c.Data["historyCount"], err = models.Count(db, coll_history)
	if err != nil {
		fmt.Println("get history count failed:", err)
		c.Data["historyCount"] = 0
	}
	c.Data["now"] = time.Now().Format("2006-01-02 15:04:05")
	c.Data["website"] = "starguard.cn"
	c.Data["email"] = "liuxinhao4@huawei.com"
	c.Data["CheckDeadline"] = utils.CheckDeadline
	c.Data["GetRemainTime"] = utils.GetRemainTime
	c.TplName = "index.tpl"
}
