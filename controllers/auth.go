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

	// Check if redirected from sign up page and successfully created one
	userCreated := ctx.GetSession("user_created")
	ctx.Data["UserCreated"] = userCreated

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
				"errors": err.Error(),
			},
		}
		ctx.ServeJSON()
		return
	}

	user, err := auth.LogIn(authData)
	if err != nil {
		log.Println("[ERROR] DoLogin:", err.Error())
		ctx.Data["json"] = models.JSONResponse{
			HttpRes: 200,
			Status:  "Invalid Credentials",
			Data: map[string]interface{}{
				"response": "invalid username/password",
				"errors":   err.Error(),
			},
		}
		ctx.ServeJSON()
		return
	}

	ctx.SetSession("current_user", user)
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

func (ctx *AuthController) DoLogout() {
	ctx.SetSession("current_user", nil)

	ctx.Redirect("/", 302)
}

func (ctx *AuthController) ForgetPassword() {
	// Check if successfully generated password forget token or not
	forgotten := ctx.GetSession("forget_password")
	ctx.Data["Forgotten"] = forgotten
	ctx.SetSession("forget_password", false)

	ctx.Layout = "auth/base.html"            // Main Layout
	ctx.TplName = "auth/forgetPassword.html" // .LayoutContent
	ctx.Data["xsrf_token"] = ctx.XSRFToken()
	ctx.Data["xsrfdata"] = template.HTML(ctx.XSRFFormHTML())

	ctx.LayoutSections = make(map[string]string)
	ctx.LayoutSections["Scripts"] = "auth/scripts.html" // .Scripts
}

func (ctx *AuthController) DoForgetPassword() {
	type UserRequest struct {
		Username string `json:"username"`
	}

	username := UserRequest{}

	err := json.Unmarshal(ctx.Ctx.Input.RequestBody, &username)
	if err != nil {
		log.Println("[ERROR] DoForgetPassword:", err.Error())
		ctx.Data["json"] = models.JSONResponse{
			HttpRes: 500,
			Status:  "Internal Server Error",
			Data: map[string]interface{}{
				"errors": err.Error(),
			},
		}
		ctx.ServeJSON()
		return
	}

	token, err := auth.ForgetPassword(username.Username)
	if err != nil {
		log.Println("[ERROR] DoForgetPassword:", err.Error())
		if err.Error() == "user is not found" {
			ctx.Data["json"] = models.JSONResponse{
				HttpRes: 400,
				Status:  "User is not found",
				Data: map[string]interface{}{
					"response": "not found",
					"errors":   err.Error(),
				},
			}
		} else {
			ctx.Data["json"] = models.JSONResponse{
				HttpRes: 500,
				Status:  "Internal Server Error",
				Data: map[string]interface{}{
					"errors": err.Error(),
				},
			}
		}
		ctx.ServeJSON()
		return
	}

	ctx.SetSession("forget_password", true)
	log.Printf("[INFO] Token generated for %s: %s", username.Username, token)

	response := models.JSONResponse{
		HttpRes: 302,
		Data: map[string]interface{}{
			"response": "ok",
		},
		Status: "Success",
	}

	ctx.Data["json"] = &response
	ctx.ServeJSON()
}

func (ctx *AuthController) ResetPassword() {}

func (ctx *AuthController) DoResetPassword() {}

func (ctx *AuthController) SuccessResetPassword() {}

func (ctx *AuthController) SignUp() {
	// Clear User Creation Session
	ctx.SetSession("user_created", false)

	ctx.Layout = "auth/base.html"    // Main Layout
	ctx.TplName = "auth/signup.html" // .LayoutContent
	ctx.Data["xsrf_token"] = ctx.XSRFToken()
	ctx.Data["xsrfdata"] = template.HTML(ctx.XSRFFormHTML())

	ctx.LayoutSections = make(map[string]string)
	ctx.LayoutSections["Scripts"] = "auth/scripts.html" // .Scripts
}

func (ctx *AuthController) DoSignUp() {
	new_user := models.User{}
	err := json.Unmarshal(ctx.Ctx.Input.RequestBody, &new_user)

	if err != nil {
		log.Println("[ERROR] DoSignUp:", err.Error())
		ctx.Data["json"] = models.JSONResponse{
			HttpRes: 500,
			Status:  "Internal Server Error",
			Data: map[string]interface{}{
				"errors": err.Error(),
			},
		}
		ctx.ServeJSON()
		return
	}

	err = auth.SignUp(new_user)
	if err != nil {
		log.Println("[ERROR] DoSignUp:", err.Error())
		ctx.Data["json"] = models.JSONResponse{
			HttpRes: 500,
			Status:  "Internal Server Error",
			Data: map[string]interface{}{
				"errors": err.Error(),
			},
		}
		ctx.ServeJSON()
		return
	}

	ctx.SetSession("user_created", true)
	response := models.JSONResponse{
		HttpRes: 302,
		Data: map[string]interface{}{
			"response": "ok",
		},
		Status: "User Created",
	}

	log.Println("[INFO] User Creation : OK")
	ctx.Data["json"] = &response
	ctx.ServeJSON()
}
