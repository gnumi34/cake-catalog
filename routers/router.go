package routers

import (
	"cake-catalog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	auth := beego.NewNamespace("/auth",
		beego.NSRouter("/login", &controllers.AuthController{}, "get:LoginPage"),
		beego.NSRouter("/login", &controllers.AuthController{}, "post:DoLogin"),
		beego.NSRouter("/logout", &controllers.AuthController{}, "post:DoLogout"),
		beego.NSRouter("/forgetPassword", &controllers.AuthController{}, "get:ForgetPassword"),
		beego.NSRouter("/forgetPassword", &controllers.AuthController{}, "post:DoForgetPassword"),
		beego.NSRouter("/resetPassword/:token", &controllers.AuthController{}, "get:ResetPassword"),
		beego.NSRouter("/resetPassword/:token", &controllers.AuthController{}, "post:DoResetPassword"),
		beego.NSRouter("/resetPasswordSuccess", &controllers.AuthController{}, "get:SuccessResetPassword"),
	)

	main := beego.NewNamespace("/",
		beego.NSRouter("", &controllers.MainController{}, "get:Homepage"),
	)

	catalog := beego.NewNamespace("/catalog",
		beego.NSRouter("", &controllers.CatalogController{}, "get:GetCatalog"),
		beego.NSRouter("cake/:id", &controllers.CatalogController{}, "get:GetCake"),
		beego.NSRouter("cake/:id/update", &controllers.CatalogController{}, "put:UpdateCake"),
		beego.NSRouter("cake/:id/delete", &controllers.CatalogController{}, "delete:DeleteCake"),
		beego.NSRouter("cake/add", &controllers.CatalogController{}, "get:AddCake"),
		beego.NSRouter("cake/add", &controllers.CatalogController{}, "post:DoAddCake"),
	)

	beego.AddNamespace(main)
	beego.AddNamespace(auth)
	beego.AddNamespace(catalog)
}
