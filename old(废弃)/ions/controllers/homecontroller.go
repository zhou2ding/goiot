package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions/models/auth"
	"ions/utils"
	"math"
	"time"
)

type HomeController struct {
	beego.Controller
}

func (h *HomeController) Get() {
	o := orm.NewOrm()

	var auIds []int
	uid := h.GetSession("id").(int)
	usr := auth.User{Id: uid}
	_, _ = o.LoadRelated(&usr, "Role")
	//把每个用户的所有角色反向查询出来，再遍历这些角色把id存进去，会有重复的，所以后面要用expr表达式的in
	for _, rol := range usr.Role {
		rol := auth.Role{Id: rol.Id}
		_, _ = o.LoadRelated(&rol, "Auth")
		for _, au := range rol.Auth {
			auIds = append(auIds, au.Id)
		}
	}
	qs := o.QueryTable("sys_auth")
	//当前用户的所有的一级菜单
	var aus []auth.Auth
	_, _ = qs.Filter("is_del", 0).Filter("parent_id", 0).Filter("id__in", auIds).OrderBy("-weight").All(&aus)

	var trees []*auth.MenuTree
	for _, au := range aus {
		treeData := auth.MenuTree{Id: au.Id, AuthName: au.AuthName, UrlFor: au.UrlFor, Weight: au.Weight, Children: []*auth.MenuTree{}}
		GetChildNode(au.Id, &treeData)
		trees = append(trees, &treeData)
	}
	//原始写法，把root这个空的Menu当做根菜单传进去递归，不利于后续修改
	//auNames := make(map[string]int, 100)
	//for _, rol := range usr.Role {
	//	tmp := auth.Role{Id: rol.Id}
	//	_, _ = o.LoadRelated(&tmp, "Auth")
	//	for i, au := range tmp.Auth {
	//		auNames[au.AuthName] = i
	//	}
	//}
	//root := auth.MenuTree{}
	//GetChildNode(root.Id, &root)
	//for _, child := range root.Children {
	//	fmt.Println(child.AuthName)
	//	if _, ok := auNames[child.AuthName]; !ok {
	//		*child = auth.MenuTree{}
	//	}
	//}

	// 消息通知,发送消息，使用定时任务优化
	qs1 := o.QueryTable("sys_cars_apply")
	var carsApply []auth.CarsApply
	_, _ = qs1.Filter("user_id", uid).Filter("return_status", 0).Filter("notify_tag", 0).All(&carsApply)

	curTime, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))

	for _, apply := range carsApply {
		returnDate := apply.ReturnDate
		ret := curTime.Sub(returnDate)
		content := fmt.Sprintf("%s用户，你借的车辆归还时间为%v,已经预期，请尽快归还!!", usr.UserName, returnDate.Format("2006-01-02"))
		if ret > 0 { // 已经逾期
			messageNotify := auth.MessageNotify{
				Flag:    1,
				Title:   "车辆归还逾期",
				Content: content,
				User:    &usr,
				ReadTag: 0,
			}
			_, _ = o.Insert(&messageNotify)
		}

		apply.NotifyTag = 1

		_, _ = o.Update(&apply)

	}

	// 展示消息,使用websocket优化

	qs2 := o.QueryTable("sys_message_notify")
	notifyCount, _ := qs2.Filter("read_tag", 0).Count()
	_ = o.QueryTable("sys_user").Filter("id", uid).One(&usr)
	h.Data["trees"] = trees
	h.Data["user"] = usr
	h.Data["notify_count"] = notifyCount
	h.TplName = "index.html"
}

func (h *HomeController) Welcome() {
	h.TplName = "welcome.html"
}

func GetChildNode(parentId int, parentTree *auth.MenuTree) {
	var auths []auth.Auth
	o := orm.NewOrm()
	qs := o.QueryTable("sys_auth")
	_, err := qs.Filter("parent_id", parentId).All(&auths)
	if err != nil {
		return
	}
	for i, v := range auths {
		treeData := auth.MenuTree{Id: auths[i].Id, AuthName: auths[i].AuthName, UrlFor: auths[i].UrlFor, Weight: auths[i].Weight, Children: []*auth.MenuTree{}}
		parentTree.Children = append(parentTree.Children, &treeData)
		GetChildNode(v.Id, &treeData)
	}
}

func (h *HomeController) NotifyList() {
	o := orm.NewOrm()

	qs := o.QueryTable("sys_message_notify")

	var nofities []auth.MessageNotify
	// 每页显示的条数
	pagePerNum := 8
	// 当前页
	currentPage, err := h.GetInt("page")

	offsetNum := pagePerNum * (currentPage - 1)

	kw := h.GetString("kw")
	var count int64 = 0

	ret := fmt.Sprintf("当前页;%d,查询条件：%s", currentPage, kw)
	logs.Info(ret)
	if kw != "" { // 有查询条件的
		// 总数
		count, _ = qs.Filter("title__contains", kw).Count()
		_, _ = qs.Filter("title__contains", kw).Limit(pagePerNum).Offset(offsetNum).All(&nofities)
	} else {
		count, _ = qs.Count()
		_, _ = qs.Limit(pagePerNum).Offset(offsetNum).All(&nofities)

	}
	if err != nil { // 说明没有获取到当前页
		currentPage = 1
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

	h.Data["nofities"] = nofities
	h.Data["prePage"] = prePage
	h.Data["nextPage"] = nextPage
	h.Data["currentPage"] = currentPage
	h.Data["countPage"] = countPage
	h.Data["count"] = count
	h.Data["page_map"] = pageMap
	h.Data["kw"] = kw

	h.TplName = "notify_list.html"

}

func (h *HomeController) ReadNotify() {
	id, _ := h.GetInt("id")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_message_notify")
	_, _ = qs.Filter("id", id).Update(orm.Params{
		"read_tag": 1,
	})
	h.Redirect(beego.URLFor("HomeController.NotifyList"), 302)

}
