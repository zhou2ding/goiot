package news

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions/models/news"
	"ions/utils"
	"math"
	"strconv"
	"time"
)

type NewsController struct {
	beego.Controller
}

func (n *NewsController) Get() {

	o := orm.NewOrm()

	qs := o.QueryTable("sys_news")

	var newsData []news.News
	// 每页显示的条数
	pagePerNum := 8
	// 当前页
	currentPage, err := n.GetInt("page")
	if err != nil { // 说明没有获取到当前页
		currentPage = 1
	}

	offsetNum := pagePerNum * (currentPage - 1)

	kw := n.GetString("kw")

	var count int64 = 0

	ret := fmt.Sprintf("当前页;%d,查询条件：%s", currentPage, kw)
	logs.Info(ret)
	if kw != "" { // 有查询条件的
		// 总数
		count, _ = qs.Filter("is_delete", 0).Filter("title__contains", kw).Count()
		_, _ = qs.Filter("is_delete", 0).Filter("title__contains", kw).Limit(pagePerNum).Offset(offsetNum).RelatedSel().All(&newsData)
	} else {
		count, _ = qs.Filter("is_delete", 0).Count()
		_, _ = qs.Filter("is_delete", 0).Limit(pagePerNum).Offset(offsetNum).RelatedSel().All(&newsData)

	}

	// 总页数
	countPage := int(math.Ceil(float64(count) / float64(pagePerNum)))

	prePage := 1
	if currentPage == 1 {
		prePage = currentPage
	} else if currentPage > 1 {
		prePage = currentPage - 1
	}

	nextPage := 1
	if currentPage < countPage {
		nextPage = currentPage + 1
	} else if currentPage >= countPage {
		nextPage = currentPage
	}

	pageMap := utils.Paginator(currentPage, pagePerNum, count)

	n.Data["news_data"] = newsData
	n.Data["prePage"] = prePage
	n.Data["nextPage"] = nextPage
	n.Data["currentPage"] = currentPage
	n.Data["countPage"] = countPage
	n.Data["count"] = count
	n.Data["page_map"] = pageMap
	n.Data["kw"] = kw

	n.TplName = "news/news_list.html"

}

func (n *NewsController) ToAdd() {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_category")
	var categories []news.Category
	_, _ = qs.Filter("is_delete", 0).All(&categories)
	n.Data["categories"] = categories
	n.TplName = "news/news_add.html"

}

func (n *NewsController) DoAdd() {
	content := n.GetString("content")
	title := n.GetString("title")
	categoryId, _ := n.GetInt("category_id")
	isActive, _ := n.GetInt("is_active")

	category := news.Category{Id: categoryId}
	o := orm.NewOrm()
	newsData := news.News{
		Content:  content,
		Title:    title,
		Category: &category,
		IsActive: isActive,
	}
	_, err := o.Insert(&newsData)

	messageMap := map[string]interface{}{}
	if err != nil {
		messageMap["code"] = 10001
		messageMap["msg"] = "添加失败"
	}
	messageMap["code"] = 200
	messageMap["msg"] = "添加成功"

	n.Data["json"] = messageMap
	n.ServeJSON()
}

func (n *NewsController) UploadImg() {

	f, h, err := n.GetFile("file")
	messageMap := map[string]interface{}{}
	if err != nil {
		messageMap["code"] = 10001
		messageMap["msg"] = "文件上传失败"
		return
	}

	defer func() {
		_ = f.Close()
	}()

	fileName := h.Filename

	timeUnixInt := time.Now().Unix()
	timeUnitStr := strconv.FormatInt(timeUnixInt, 10)

	filePath := "upload/news_img/" + timeUnitStr + "-" + fileName

	imgLink := "http://127.0.0.1:8080/" + filePath

	_ = n.SaveToFile("file", filePath)

	messageMap["code"] = 200
	messageMap["msg"] = "文件上传成功"
	messageMap["link"] = imgLink

	n.Data["json"] = messageMap
	n.ServeJSON()

}

func (n *NewsController) ToEdit() {
	nId, _ := n.GetInt("id")

	o := orm.NewOrm()
	qs := o.QueryTable("sys_news")

	newsData := news.News{}
	_ = qs.Filter("id", nId).RelatedSel().One(&newsData)
	n.Data["news_data"] = newsData

	var categories []news.Category
	_, _ = o.QueryTable("sys_category").Exclude("id", newsData.Category.Id).All(&categories)

	n.Data["news_data"] = newsData
	n.Data["categories"] = categories
	n.TplName = "news/news_edit.html"

}

func (n *NewsController) DoEdit() {

	newsId, _ := n.GetInt("news_id")
	content := n.GetString("content")
	title := n.GetString("title")
	categoryId, _ := n.GetInt("category_id")
	isActive, _ := n.GetInt("is_active")

	o := orm.NewOrm()
	qs := o.QueryTable("sys_news")
	_, err := qs.Filter("id", newsId).Update(orm.Params{
		"title":       title,
		"content":     content,
		"category_id": categoryId,
		"is_active":   isActive,
	})
	messageMap := map[string]interface{}{}
	if err != nil {
		messageMap["code"] = 10001
		messageMap["msg"] = "更新失败"
	}
	messageMap["code"] = 200
	messageMap["msg"] = "更新成功"

	n.Data["json"] = messageMap
	n.ServeJSON()

}
