package controllers

import (
	"html/template"

	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

func (ctx *AuthController) LoginPage() {
	currentUser := ctx.GetSession("current_user")
	if currentUser != nil {
		ctx.Redirect("/", 302)
	}

	ctx.Layout = "auth/base.html"   // Main Layout
	ctx.TplName = "auth/login.html" // .LayoutContent
	ctx.Data["xsrf_token"] = ctx.XSRFToken()
	ctx.Data["xsrfdata"] = template.HTML(ctx.XSRFFormHTML())

	ctx.LayoutSections = make(map[string]string)
	ctx.LayoutSections["Head"] = "auth/head.html"      // .Head
	ctx.LayoutSections["Sripts"] = "auth/scripts.html" // .Scripts
}
