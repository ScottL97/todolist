package controllers

import (
	"fmt"
	"strings"
	"todolist/models"
	"todolist/utils"

	"github.com/astaxie/beego"
	"github.com/globalsign/mgo/bson"
	"github.com/siddontang/go/num"
)

type ItemController struct {
	beego.Controller
}

func (c *ItemController) Get() {
	db := beego.AppConfig.String("mongodb::mongo_db")
	coll_items := beego.AppConfig.String("mongodb::mongo_coll_items")

	idParam := c.Ctx.Input.Param(":id")
	if idParam != "" {
		id, err := num.ParseInt(idParam)
		if err != nil {
			fmt.Printf("wrong id = %d, err: %s", id, err)
		}
		var item models.Item
		err = models.FindOne(db, coll_items, bson.M{"id": id}, &item)
		if err != nil {
			fmt.Printf("find id = %d failed, err: %s", id, err)
		}
		// 时间格式修改为表单默认的
		item.Deadline = strings.ReplaceAll(item.Deadline, " ", "T")
		c.Data["item"] = item
	}
	c.Data["JoinPeople"] = utils.JoinPeople
	c.TplName = "item.tpl"
}

func (c *ItemController) Post() {
	db := beego.AppConfig.String("mongodb::mongo_db")
	coll_items := beego.AppConfig.String("mongodb::mongo_coll_items")

	// 解析并修改要插入数据库的item结构体
	item := models.Item{}
	err := c.ParseForm(&item)
	if err != nil {
		fmt.Println("ParseForm failed, err:", err)
		return
	}
	item.Deadline = strings.ReplaceAll(item.Deadline, "T", " ")
	if strings.Count(item.Deadline, ":") == 1 {
		item.Deadline += ":00"
	}
	item.People = strings.Split(c.Input().Get("people"), ";")
	// 如果URL为/item，说明是插入；如果URL为/item/xx，表示修改id为xx的记录
	idParam := c.Ctx.Input.Param(":id")
	if idParam == "" {
		item.Id = utils.GetUsableId()
		if item.Id == -1 {
			fmt.Println("GetUsableId failed")
			return
		}
		err = models.Insert(db, coll_items, item)
		if err != nil {
			fmt.Println("insert todolist failed:", err)
			return
		}
		fmt.Println("插入事项：", item)
	} else {
		item.Id, err = num.ParseInt(idParam)
		if err != nil {
			fmt.Printf("parse id = %s failed, err: %s\n", idParam, err)
			return
		}
		err = models.Update(db, coll_items, bson.M{"id": item.Id}, item)
		if err != nil {
			fmt.Printf("idParam: %s, update todolist failed: %s\n", idParam, err)
			return
		}
		fmt.Println("更新事项：", item)
	}

	c.Ctx.Redirect(302, "/")
}
