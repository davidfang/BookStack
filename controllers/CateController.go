package controllers

import (
	"github.com/TruthHun/BookStack/models"
	"github.com/astaxie/beego"
)

type CateController struct {
	BaseController
}

func (this *CateController) Index() {
	cid, _ := this.GetInt("cid")
	if cid > 0 {
		this.Redirect(beego.URLFor("HomeController.Index")+this.Ctx.Request.RequestURI, 302)
	}
	this.List()
}

//分类
func (this *CateController) List() {
	if cates, err := new(models.Category).GetCates(-1, 1); err == nil {
		this.Data["Cates"] = cates
	} else {
		beego.Error(err.Error())
	}
	this.GetSeoByPage("cate", map[string]string{
		"title":       "书籍分类",
		"keywords":    this.Option["SITE_KEYWORDS"],
		"description": this.Sitename + this.Option["SITE_DESCRIPTION"],
	})
	this.Data["IsCate"] = true
	this.Data["Friendlinks"] = new(models.FriendLink).GetList(false)
	this.Data["Recommends"], _, _ = models.NewBook().HomeData(1, 12, models.OrderLatestRecommend, "", 0)
	this.TplName = "cates/list.html"
}
