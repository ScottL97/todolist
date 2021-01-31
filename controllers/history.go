package controllers

import (
	"fmt"
	"todolist/models"
	"todolist/utils"

	"github.com/astaxie/beego"
	"github.com/globalsign/mgo/bson"
	"github.com/siddontang/go/num"
)

type HistoryController struct {
	beego.Controller
}

func (c *HistoryController) Get() {
	db := beego.AppConfig.String("mongodb::mongo_db")
	coll_history := beego.AppConfig.String("mongodb::mongo_coll_history")

	var res []models.Item
	err := models.FindAll(db, coll_history, nil, &res)
	if err != nil {
		fmt.Println("error to FindAll:", err)
	}
	c.Data["items"] = res
	c.Data["CheckDeadline"] = utils.CheckDeadline
	c.Data["GetRemainTime"] = utils.GetRemainTime
	c.TplName = "history.tpl"
}

func (c *HistoryController) Post() {
	db := beego.AppConfig.String("mongodb::mongo_db")
	coll_items := beego.AppConfig.String("mongodb::mongo_coll_items")
	coll_history := beego.AppConfig.String("mongodb::mongo_coll_history")

	id, err := num.ParseInt(c.Input().Get("id"))
	if err != nil {
		fmt.Println("parse id failed:", err)
		return
	}
	// 如果查询记录不存在，不进行插入和删除操作
	var item models.Item
	err = models.FindOne(db, coll_items, bson.M{"id": id}, &item)
	if err != nil {
		fmt.Printf("find id = %d failed: %s\n", id, err)
		return
	}
	models.Insert(db, coll_history, item)
	models.Remove(db, coll_items, bson.M{"id": id})
}
