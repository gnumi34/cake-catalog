package controllers

import (
	"cake-catalog/models"
	"cake-catalog/models/auth"
	"encoding/json"
	"html/template"
	"log"

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
	ctx.LayoutSections["Scripts"] = "auth/scripts.html" // .Scripts
}

func (ctx *AuthController) DoLogin() {
	if ctx.GetSession("current_user") != nil {
		ctx.Redirect("/", 302)
	}

	authData := models.AuthRequest{}
	err := json.Unmarshal(ctx.Ctx.Input.RequestBody, &authData)

	if err != nil {
		log.Println(err.Error())
		ctx.Data["json"] = models.JSONResponse{
			HttpRes: 500,
			Status:  "Internal Server Error",
			Data: map[string]interface{}{
				"error": err.Error(),
			},
		}
		ctx.ServeJSON()
		return
	}

	user, err := auth.LogIn(authData)
	if err != nil {
		log.Println(err.Error())
		ctx.Data["json"] = models.JSONResponse{
			HttpRes: 200,
			Status:  "Invalid Credentials",
			Data: map[string]interface{}{
				"response": "not found",
				"error":    err.Error(),
			},
		}
		ctx.ServeJSON()
		return
	}

	ctx.SetSession("current_user", user.ID)
	response := models.JSONResponse{
		HttpRes: 302,
		Data: map[string]interface{}{
			"response": "ok",
		},
		Status: "Credential found",
	}

	log.Println("[INFO] Login Session : OK for", user.ID)
	ctx.Data["json"] = &response
	ctx.ServeJSON()
}

func (ctx *AuthController) DoLogout() {}

func (ctx *AuthController) ForgetPassword() {}

func (ctx *AuthController) DoForgetPassword() {}

func (ctx *AuthController) ResetPassword() {}

func (ctx *AuthController) DoResetPassword() {}

func (ctx *AuthController) SuccessResetPassword() {}

func (ctx *AuthController) SignUp() {}

func (ctx *AuthController) DoSignUp() {}
