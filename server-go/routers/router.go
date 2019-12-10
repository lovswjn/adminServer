package routers

import (
	"server-go/controllers"

	"github.com/astaxie/beego"
)

func init() {
	user := beego.NewNamespace("user",
		beego.NSRouter("/login", &controllers.AuthController{}, "post:DoLogin"),
		beego.NSRouter("/info", &controllers.AuthController{}, "get:Info"),
	)
	beego.Router("/auth/logout", &controllers.AuthController{}, "post:Logout")
	beego.Router("/upload", &controllers.BaseController{}, "post:Upload")
	beego.Router("/settings", &controllers.AuthController{}, "put:Setting") // 修改资料

	rbac := beego.NewNamespace("auth",
		beego.NSBefore(),
		// 规则
		beego.NSRouter("/rule", &controllers.RbacController{}, "get:Rules"),
		beego.NSRouter("/ruleAdd", &controllers.RbacController{}, "post:RuleAdd"),
		beego.NSRouter("/rule/:id", &controllers.RbacController{}, "put:RuleEdit"),
		beego.NSRouter("/ruleDelete/:id", &controllers.RbacController{}, "delete:RuleDelete"),
		beego.NSRouter("/tree", &controllers.RbacController{}, "get:Tree"),
		// 角色
		beego.NSRouter("/role", &controllers.RbacController{}, "get:Roles"),
		beego.NSRouter("/roleAdd", &controllers.RbacController{}, "post:RoleAdd"),
		beego.NSRouter("/role/:id", &controllers.RbacController{}, "put:RoleEdit"),
		beego.NSRouter("/roleDelete/:id", &controllers.RbacController{}, "delete:RoleDelete"),
		// 管理员
		beego.NSRouter("/user", &controllers.RbacController{}, "get:Users"),
		beego.NSRouter("/userAdd", &controllers.RbacController{}, "post:UserAdd"),
		beego.NSRouter("/userDelete/:id", &controllers.RbacController{}, "delete:UserDelete"),
		beego.NSRouter("/user/:id", &controllers.RbacController{}, "put:UserEdit"),
	)
	stock := beego.NewNamespace("stock",
		beego.NSBefore(),
		//库存记录查看
		beego.NSRouter("/recordstock", &controllers.StockbacController{}, "get:Stock"),
		//库存添加
		beego.NSRouter("/addStock", &controllers.StockbacController{}, "post:StockAdd"),
		//库存更新
		beego.NSRouter("/stok/updateStock/:id", &controllers.StockbacController{}, "put:StockUpdate"),
	)
	beego.AddNamespace(user, rbac, stock)

}
