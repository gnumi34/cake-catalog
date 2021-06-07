package controllers

import "github.com/astaxie/beego"

type CatalogController struct {
	beego.Controller
}

func (ctx *CatalogController) GetCatalog() {}

func (ctx *CatalogController) GetCake() {}

func (ctx *CatalogController) UpdateCake() {}

func (ctx *CatalogController) DeleteCake() {}

func (ctx *CatalogController) AddCake() {}

func (ctx *CatalogController) DoAddCake() {}
