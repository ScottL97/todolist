package utils

import (
	"fmt"
	"sort"
	"strings"
	"time"
	"todolist/models"

	"github.com/astaxie/beego"
	"github.com/globalsign/mgo/bson"
)

func init() {
	//err := beego.AddFuncMap("CheckDeadline", CheckDeadline)
	//if err != nil {
	//  fmt.Println("AddFuncMap failed:", err)
	//}
	//err = beego.AddFuncMap("GetRemainTime", GetRemainTime)
	//if err != nil {
	//  fmt.Println("AddFuncMap failed:", err)
	//}
}

func CheckDeadline(deadline string) string {
	dt, err := time.ParseInLocation("2006-01-02 15:04:05", deadline, time.Local)
	if err != nil {
		fmt.Println("error to parse deadline:", err)
		return ""
	}
	if time.Now().After(dt) {
		return "-alert"
	}
	return ""
}

func GetRemainTime(deadline string) string {
	dt, err := time.ParseInLocation("2006-01-02 15:04:05", deadline, time.Local)
	if err != nil {
		fmt.Println("error to parse deadline:", err)
		return ""
	}
	diff := time.Now().Unix() - dt.Unix()
	if diff > 0 {
		return fmt.Sprintf("超期：%d 时 %d 分", diff/3600, diff%3600/60)
	}
	return fmt.Sprintf("剩余：%d 时 %d 分", -diff/3600, -diff%3600/60)
}

func JoinPeople(people []string) string {
	return strings.Join(people, ";")
}

func GetUsableId() int {
	db := beego.AppConfig.String("mongodb::mongo_db")
	coll_items := beego.AppConfig.String("mongodb::mongo_coll_items")
	coll_history := beego.AppConfig.String("mongodb::mongo_coll_history")

	var items []models.Item
	var history []models.Item
	id := -1

	err := models.FindAllWithSelect(db, coll_items, nil, bson.M{"id": 1}, &items)
	if err != nil {
		fmt.Println("failed to get collection items:", err)
		return id
	}
	err = models.FindAllWithSelect(db, coll_history, nil, bson.M{"id": 1}, &history)
	if err != nil {
		fmt.Println("failed to get collection items:", err)
		return id
	}

	var ids []int
	for _, item := range items {
		ids = append(ids, item.Id)
	}
	for _, hist := range history {
		ids = append(ids, hist.Id)
	}
	if len(ids) == 0 {
		return 1
	}

	sort.Ints(ids)
	for i := 0; i < len(ids)-1; i++ {
		if ids[i+1]-ids[i] != 1 {
			id = ids[i] + 1
			break
		}
	}
	if id == -1 {
		id = len(ids) + 1
	}

	return id
}
