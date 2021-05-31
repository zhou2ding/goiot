package auth

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ions/models/auth"
	"math"
	"ions/utils"
	"strconv"
	"time"
	"strings"
)

type RoleController struct {
	beego.Controller
}

func (r *RoleController) List()  {

	var roles []auth.Role
	o := orm.NewOrm()


	// 每页显示的条数
	pagePerNum := 8
	// 当前页
	currentPage,err := r.GetInt("page")
	if err != nil {   // 说明没有获取到当前页
		currentPage = 1
	}

	offsetNum := pagePerNum * (currentPage - 1)

	qs := o.QueryTable("sys_role")
	_,_ = qs.Filter("is_del",0).All(&roles)

	count,_ := qs.Filter("is_del",0).Count()
	_,_ = qs.Filter("is_del",0).Limit(pagePerNum).Offset(offsetNum).All(&roles)

	// 总页数
	countPage := int(math.Ceil(float64(count) / float64(pagePerNum)))

	prePage := 1
	if currentPage == 1{
		prePage = currentPage
	}else if currentPage > 1{
		prePage = currentPage -1
	}

	nextPage := 1
	if currentPage < countPage {
		nextPage = currentPage + 1
	}else if currentPage >= countPage {
		nextPage = currentPage
	}


	pageMap := utils.Paginator(currentPage,pagePerNum,count)



	r.Data["roles"] = roles
	r.Data["prePage"] =prePage
	r.Data["nextPage"] = nextPage
	r.Data["currentPage"] = currentPage
	r.Data["countPage"] = countPage
	r.Data["count"] = count
	r.Data["page_map"] = pageMap
	r.TplName = "auth/role-list.html"
}

func (r *RoleController) ToAdd()  {
	r.TplName = "auth/role_add.html"

}

func (r *RoleController) DoAdd()  {

	roleName := r.GetString("role_name")
	desc := r.GetString("desc")
	isActive,_ := r.GetInt("is_active")

	role := auth.Role{RoleName: roleName,Desc:desc,IsActive: isActive,CreateTime:time.Now()}
	o := orm.NewOrm()
	_,err := o.Insert(&role)

	messageMap := map[string]interface{}{}
	if err != nil { // 发生错误
		messageMap["code"] = 10001
		messageMap["msg"] = "添加数据错误，请重新添加"

	}else {
		messageMap["code"] = 200
		messageMap["msg"] = "添加成功"
	}

	r.Data["json"] = messageMap
	r.ServeJSON()

}

// 角色--一用户配置
func (r *RoleController) ToRoleUser()  {
	id,_ := r.GetInt("role_id")

	o := orm.NewOrm()
	role := auth.Role{}
	_ = o.QueryTable("sys_role").Filter("id",id).One(&role)

	// 已绑定的用户
	_,_ = o.LoadRelated(&role,"User")


	// 未绑定的用户
	var users []auth.User
	if len(role.User) > 0 {
		_,_ = o.QueryTable("sys_user").Filter("is_del",0).Filter("is_active",1).Exclude("id__in",role.User).All(&users)

	}else {   // 没有绑定的数据
		_,_ = o.QueryTable("sys_user").Filter("is_del",0).Filter("is_active",1).All(&users)

	}

	r.Data["role"] = role
	r.Data["users"] = users
	r.TplName = "auth/role-user-add.html"

}


// 角色--一用户配置
func (r *RoleController) DoRoleUser()  {
	roleId,_ := r.GetInt("role_id")
	userIds := r.GetString("user_ids")

	//new_user_ids := user_ids[1:len(user_ids)-1]
	userIdArr := strings.Split(userIds,",")

	// "10,12,13"

	o := orm.NewOrm()
	role := auth.Role{Id: roleId}

	// 查询出已绑定的数据
	m2m := o.QueryM2M(&role,"User")
	_,_ = m2m.Clear()

	for _, userId := range userIdArr {
		tmp,_ := strconv.Atoi(userId)
		user := auth.User{Id:tmp}
		m2m := o.QueryM2M(&role,"User")
		_,_ = m2m.Add(&user)

	}

	r.Data["json"] = map[string]interface{}{"code":200,"msg":"添加成功"}
	r.ServeJSON()
}

// 角色--权限配置
func (r *RoleController) ToRoleAuth()  {
	roleId,_ := r.GetInt("role_id")

	o := orm.NewOrm()
	qs := o.QueryTable("sys_role")
	role := auth.Role{}
	_ = qs.Filter("id", roleId).One(&role)
	r.Data["role"] = role
	r.TplName = "auth/role-auth-add.html"

}

func (r *RoleController) GetAuthJson()  {
	roleId,_ := r.GetInt("role_id")


	o := orm.NewOrm()
	qs := o.QueryTable("sys_auth")

	// 已绑定的权限
	role := auth.Role{Id: roleId}
	_,_ = o.LoadRelated(&role,"Auth")

	//[11,14,16]
	var authIdsHas []int
	for _, authData := range role.Auth{
		authIdsHas = append(authIdsHas, authData.Id)
	}




	// 所有的权限
	var auths []auth.Auth
	_,_ = qs.Filter("is_del",0).All(&auths)

	var authArrMap []map[string]interface{} // map数组

	for _, authData := range auths{
		id := authData.Id
		pId := authData.ParentId
		name := authData.AuthName
		if pId == 0 {
			authMap := map[string]interface{}{"id": id,"pId":pId,"name":name,"open":false}
			authArrMap = append(authArrMap, authMap)
		}else {
			authMap := map[string]interface{}{"id": id,"pId":pId,"name":name}
			authArrMap = append(authArrMap, authMap)
		}

	}

	authMaps := map[string]interface{}{}
	authMaps["auth_arr_map"] = authArrMap
	authMaps["auth_ids_has"] = authIdsHas
	r.Data["json"] = authMaps
	r.ServeJSON()

}

func (r *RoleController) DoRoleAuth()  {

	roleId,_ := r.GetInt("role_id")
	authIds := r.GetString("auth_ids")
	//"13,15,16"       "13  15    16"
	//new_auth_ids := auth_ids[1:len(auth_ids)-1]
	idArr := strings.Split(authIds,",")


	o := orm.NewOrm()
	role := auth.Role{Id: roleId}
	m2m := o.QueryM2M(&role,"Auth")
	_,_ = m2m.Clear()

	for _, authId := range idArr {
		authIdInt,_ := strconv.Atoi(authId)
		if authIdInt !=0 {
			authData := auth.Auth{Id: authIdInt}
			m2m := o.QueryM2M(&role,"Auth")
			_,_ = m2m.Add(&authData)
		}

	}

	r.Data["json"] = map[string]interface{}{"code":200,"msg":"添加成功"}
	r.ServeJSON()


}