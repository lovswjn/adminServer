package models

import (
	"fmt"
	"time"
	"github.com/astaxie/beego"
)

// 入库记录表
type StockAddList struct {
	ID             uint    `gorm:"primary_key" json:"id"`
	ItemStock      string `json:"itemStock"`
	NameStock      string `json:"nameStock"`
	UnitusStock    int    `json:"unitusStock"`
	NumStock       int    `json:"numStock"`
	TypeStock      string `json:"typeStock"`
	SnStock        string `json:"snStock"`
	StatusStock    int    `json:"statusStock"`
	SourcesStock   int `json:"sourcesStock"`
	BuyerStock     int `json:"buyerStock"`
	KeeperStock    int `json:"keeperStock"`
	CreatedAt time.Time
	InputTime      string `json:"inputTime"`
	
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
				ID:             v.ID,
				ItemStock:      v.ItemStock,
				NameStock:      v.NameStock,
				UnitusStock:    v.UnitusStock,
				NumStock:       v.NumStock,
				TypeStock:      v.TypeStock,
				SnStock:        v.SnStock,
				StatusStock:    v.StatusStock,
				SourcesStock:   v.SourcesStock,
				BuyerStock:     v.BuyerStock,
				KeeperStock:    v.KeeperStock,
				InputTime:     v.CreatedAt.Format("2006-01-02 15:04:05") ,
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
		ItemStock:      param["itemStock"].(string),
		NameStock:      param["nameStock"].(string),
		UnitusStock:    int(param["unitusStock"].(float64)),
		NumStock:    
	   int(param["numStock"].(float64)),
		TypeStock:      param["typeStock"].(string),
		SnStock:        param["snStock"].(string),
		StatusStock:    int(param["statusStock"].(float64)), 
		SourcesStock:   int(param["sourcesStock"].(float64)),
		BuyerStock:     int(param["buyerStock"].(float64)),
		KeeperStock:    int(param["keeperStock"].(float64)),
	}
	if err := db.Create(&stock).Error; err != nil {
		beego.Error(err)
		fmt.Println("添加错误，主键已回滚")
		return err
	}
	permissionIds := param["rules"].([]interface{})
	sql := "INSERT INTO `auth_role_permission_access` (`role_id`,`permission_id`) VALUES"
	for key, value := range permissionIds {
		if len(permissionIds)-1 == key {
			//最后一条数据以分号结尾
			fmt.Println("---haha-----------------")
			sql += fmt.Sprintf("(%d, %d);", stock.ID, int(value.(float64)))
		} else {
			fmt.Println("---sooooooo-----------------")
			sql += fmt.Sprintf("(%d,%d),", stock.ID, int(value.(float64)))
		}
	}
	db.Exec(sql)
	return nil
}

// 角色删除
func (t *StockAddList) Delete(stockId int) error {
	tx := db.Begin()

	err := db.Where("id = ?", stockId).Delete(&StockAddList{}).Error
	if err != nil {
		beego.Error(err)
		tx.Rollback()
		return err
	}
	err = db.Where("role_id = ?", stockId).Delete(&AuthRolePermissionAccess{}).Error
	if err != nil {
		beego.Error(err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
