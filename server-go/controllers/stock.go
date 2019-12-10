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
	var err error
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
	fmt.Println("---库存记录查询成功-get:stock-contonllers内-----------------")
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

// 入库记录修改
func (t *StockbacController) StockUpdate() {
	stockId, err := t.GetInt(":id", 0)
	if err != nil || stockId == 0 {
		t.Error("paramError")
	}
	fmt.Println("---库存记录更新id-----------------")
	param := make(map[string]interface{})
	t.GetJsonParam(&param)
	fmt.Println(param)
	var stockModel models.StockAddList
	err = stockModel.Edit(stockId, param)
	if err != nil {
		t.Error("修改失败, 唯一标识重复")
	}
	t.Success("修改成功")
}
