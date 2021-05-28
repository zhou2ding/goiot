package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ions/models/auth"
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
		_,_ = o.LoadRelated(&rol,"Auth")
		for _,au := range rol.Auth {
			auIds = append(auIds,au.Id)
		}
	}
	qs := o.QueryTable("sys_auth")
	//当前用户的所有的一级菜单
	var aus []auth.Auth
	_, _ = qs.Filter("parent_id", 0).Filter("id__in",auIds).OrderBy("-weight").All(&aus)

	var trees []*auth.MenuTree
	for _, au := range aus {
		treeData := auth.MenuTree{Id: au.Id, AuthName: au.AuthName, UrlFor: au.UrlFor, Weight: au.Weight, Children: []*auth.MenuTree{}}
		GetChildNode(au.Id, &treeData)
		trees = append(trees,&treeData)
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
	_ = o.QueryTable("sys_user").Filter("id",uid).One(&usr)
	h.Data["trees"] = trees
	h.Data["user"] = usr
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
