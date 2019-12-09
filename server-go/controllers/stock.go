package controllers

import (
	"fmt"
	"server-go/models"

	"github.com/astaxie/beego"
)

type StockbacController struct {
	BaseController
}

// 入库记录查询
var stockAddListModel models.StockAddList

func (t *StockbacController) Stock() {
	fmt.Println("---stock-3tep-----------------")
	var err error
	fmt.Println("---stock-4tep-----------------")
	page, err := t.GetInt("page", 1)
	pageSize, err := t.GetInt("pageSize", 10)
	if err != nil {
		beego.Error(err)
		t.Error(err.Error())
	}
	stock, err := stockAddListModel.Stock(page, pageSize)
	if err != nil {
		beego.Error(err)
		t.Error(err.Error())
	}

	fmt.Println(stock)
	t.Success("success", stock)
}

// 角色添加
func (t *StockbacController) StockAdd() {
	param := make(map[string]interface{})
	t.GetJsonParam(&param)
	var stockModel models.StockAddList
	err := stockModel.Add(param)
	if err == nil {
		t.Success("新增成功")
	} else {
		beego.Error(err)
		t.Error("新增失败, 唯一标识重复")
	}
}

// 角色删除
func (t *StockbacController) StockDelete() {
	stockId, err := t.GetInt(":id", 0)
	if err != nil || stockId == 0 {
		t.Error("paramError")
	}
	err = stockAddListModel.Delete(stockId)
	if err != nil {
		t.Abort("500")
	}
	t.Success("删除成功")
}
