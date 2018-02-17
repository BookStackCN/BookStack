package controllers

import (
	"path/filepath"
	"strings"

	"github.com/BookStackCN/BookStack/utils"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.html"
}

func (this *MainController) Read() {
	file := this.GetString(":splat")
	if utils.Debug {
		beego.Debug("读取的文件", file)
	}
	ext := strings.ToLower(filepath.Ext(file))
	cont, err := utils.FileGetContent("books/" + file)
	if ext == ".md" || ext == ".markdown" {
		if err != nil {
			beego.Error(err.Error())
			this.Abort("404")
		}
		summary, _ := utils.FileGetContent("books/" + filepath.Dir(file) + "/summary.md")
		this.Data["Summary"] = summary
		this.Data["Markdown"] = cont
		this.TplName = "read.html"
	} else {
		this.Ctx.WriteString(cont)
	}
}
