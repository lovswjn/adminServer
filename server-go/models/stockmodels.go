package models

import (
	"fmt"

	"github.com/astaxie/beego"
)

// 入库记录表
type StockAddList struct {
	Id             int    `gorm:"primary_key" json:"id"`
	QuickCodeStock string `json:"quickCodeStock"`
	NameStock      string `json:"nameStock"`
	ItemStock      string `json:"itemStock"`
	UnitusStock    int    `json:"unitusStock"`
	//	UnitusStock    string `json:"unitusStock"`
	NumStock       string `json:"numStock"`
	TypeStock      string `json:"typeStock"`
	SnStock        string `json:"snStock"`
	StatusStock    int    `json:"statusStock"`
	SourcesStock   int    `json:"sourcesStock"`
	BuyerStock     int    `json:"buyerStock"`
	KeeperStock    int    `json:"keeperStock"`
	InputDataStock string `json:"inputDataStock"`
}

type stocklistParam struct {
	StockAddList
	Permissions []int `json:"permissions"`
}

// 角色列表
func (t *StockAddList) Stock(page int, pageSize int) (map[string]interface{}, error) {
	var stockAddLists []StockAddList
	retData := make(map[string]interface{})
	offset := (page - 1) * pageSize
	var err error
	//	err = db.Table("stock_add_list").Find(&stockAddLists).Error
	err = db.Table("stock_add_list").Offset(offset).Limit(pageSize).Order("id desc").Find(&stockAddLists).Error
	if err != nil {
		fmt.Println("查询错误")
		beego.Error(err)
		return nil, err
	} else {
		fmt.Println("查询OK")
	}
	fmt.Println(stockAddLists)
	var stocklistParams []stocklistParam
	for _, v := range stockAddLists {
		var ids []int
		stocklistParams = append(stocklistParams, stocklistParam{
			StockAddList: StockAddList{
				Id:             v.Id,
				QuickCodeStock: v.QuickCodeStock,
				NameStock:      v.NameStock,
				ItemStock:      v.ItemStock,
				UnitusStock:    v.UnitusStock,
				NumStock:       v.NumStock,
				TypeStock:      v.TypeStock,
				SnStock:        v.SnStock,
				StatusStock:    v.StatusStock,
				SourcesStock:   v.SourcesStock,
				BuyerStock:     v.BuyerStock,
				KeeperStock:    v.KeeperStock,
				InputDataStock: v.InputDataStock,
			},
			Permissions: ids,
		})
	}
	pagination := make(map[string]int)
	pagination["current"] = page
	pagination["pageSize"] = pageSize
	var total int
	db.Model(&StockAddList{}).Count(&total)
	pagination["total"] = total
	retData["data"] = stocklistParams
	retData["pagination"] = pagination
	warm := make(map[string]interface{})
	warm["roles"] = retData
	// 规则
	var permissionModel AuthPermission
	permissions := permissionModel.GetAllPermissionsTree(1, 1000)
	warm["rules"] = permissions
	fmt.Println(stocklistParams)
	fmt.Println("测试")
	fmt.Println(warm)
	fmt.Println("测试")
	return warm, nil
}

// 角色添加
func (t *StockAddList) Add(param map[string]interface{}) error {
	stock := StockAddList{
		QuickCodeStock: param["quickCodeStock"].(string),
		NameStock:      param["nameStock"].(string),
		ItemStock:      param["itemStock"].(string),
		UnitusStock:    int(param["unitusStock"].(float64)),
		NumStock:       param["numStock"].(string),
		TypeStock:      param["typeStock"].(string),
		SnStock:        param["snStock"].(string),
		StatusStock:    int(param["unitusStock"].(float64)),
		SourcesStock:   int(param["sourcesStock"].(float64)),
		BuyerStock:     int(param["buyerStock"].(float64)),
		KeeperStock:    int(param["keeperStock"].(float64)),
		//InputDataStock: param["inputDataStock"].(string),
	}
	if err := db.Create(&stock).Error; err != nil {
		beego.Error(err)
		return err
	}
	permissionIds := param["rules"].([]interface{})
	sql := "INSERT INTO `auth_role_permission_access` (`role_id`,`permission_id`) VALUES"
	for key, value := range permissionIds {
		if len(permissionIds)-1 == key {
			//最后一条数据以分号结尾
			sql += fmt.Sprintf("(%d, %d);", stock.Id, int(value.(float64)))
		} else {
			sql += fmt.Sprintf("(%d,%d),", stock.Id, int(value.(float64)))
		}
	}
	db.Exec(sql)
	return nil
}

//库存记录修改
func (t *StockAddList) Edit(roleId int, param map[string]interface{}) error {
	var role StockAddList
	db.Where("id = ?", roleId).First(&role)
	err := db.Model(&role).UpdateColumns(StockAddList{
		QuickCodeStock: param["quickCodeStock"].(string),
		NameStock:      param["nameStock"].(string),
		ItemStock:      param["itemStock"].(string),
		UnitusStock:    int(param["unitusStock"].(float64)),
		NumStock:       param["numStock"].(string),
		TypeStock:      param["typeStock"].(string),
		SnStock:        param["snStock"].(string),
		StatusStock:    int(param["unitusStock"].(float64)),
		SourcesStock:   int(param["sourcesStock"].(float64)),
		BuyerStock:     int(param["buyerStock"].(float64)),
		KeeperStock:    int(param["keeperStock"].(float64)),
	}).Error
	if err != nil {
		beego.Error(err)
		return err
	}
	return nil
}
