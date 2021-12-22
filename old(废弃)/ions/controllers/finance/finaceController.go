package finance

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions/models/salary"
	"ions/utils"
	"math"
	"strconv"
	"time"
)

type FinaController struct {
	beego.Controller
}

func (f *FinaController) Get() {
	o := orm.NewOrm()
	var sal []salary.Salary
	qs := o.QueryTable("sys_salary")
	curPage, err := f.GetInt("page")
	if err != nil {
		curPage = 1
	}
	pageSize := 6
	kw := f.GetString("month")
	month := time.Now().Format("2006-01")
	if len(kw) > 0 {
		month = kw
	}
	_, _ = qs.Filter("pay_date", month).Limit(pageSize, (curPage-1)*pageSize).All(&sal)
	all, _ := qs.Filter("pay_date", month).Count()
	allPage := math.Ceil(float64(all) / float64(pageSize))
	var prevPage int
	var nextPage int
	if curPage == 1 {
		prevPage = curPage
	} else {
		prevPage = curPage - 1
	}
	if curPage == int(allPage) {
		nextPage = curPage
	} else {
		nextPage = curPage + 1
	}
	pageMap := utils.Paginator(curPage, pageSize, all)
	f.Data["prePage"] = prevPage
	f.Data["nextPage"] = nextPage
	f.Data["all"] = all
	f.Data["allPage"] = allPage
	f.Data["curPage"] = curPage
	f.Data["pageMap"] = pageMap
	f.Data["sal"] = sal
	f.Data["month"] = month
	f.TplName = "finance/salary-list.html"
}

func (f *FinaController) ToImport() {
	f.TplName = "finance/salary-import.html"
}

func (f *FinaController) DoImport() {
	var err error
	fi, head, err := f.GetFile("upload_file")
	resp := make(map[string]interface{}, 3)
	if err != nil {
		resp["code"] = 1001
		resp["msg"] = "文件上传失败"
		return
	}
	defer func() {
		_ = fi.Close()
	}()
	unixStr := strconv.Itoa(int(time.Now().Unix()))
	fileName := head.Filename
	filePath := "static/upload/salary_upload/" + unixStr + " - " + fileName
	_ = f.SaveToFile("upload_file", filePath)

	//读取文件把文件内容导入数据库
	excl, err := excelize.OpenFile(filePath)
	if err != nil {
		logs.Error("open excel failed, error: ", err)
		return
	}
	errData := make([]string, 0, 10) //存放插入失败的工号
	rows := excl.GetRows("Sheet1")
	o := orm.NewOrm()
	for i := 1; i < len(rows); i++ {
		CardId := rows[i][2]
		BaseSalary, _ := strconv.ParseFloat(rows[i][3], 64)
		WorkDays, _ := strconv.ParseFloat(rows[i][4], 64)
		DaysOff, _ := strconv.ParseFloat(rows[i][5], 64)
		DaysLeave, _ := strconv.ParseFloat(rows[i][6], 64)
		Bonus, _ := strconv.ParseFloat(rows[i][7], 64)
		RentSubsidy, _ := strconv.ParseFloat(rows[i][8], 64)
		TransSubsidy, _ := strconv.ParseFloat(rows[i][9], 64)
		SocialSec, _ := strconv.ParseFloat(rows[i][10], 64)
		HouseFund, _ := strconv.ParseFloat(rows[i][11], 64)
		Tax, _ := strconv.ParseFloat(rows[i][12], 64)
		Fine, _ := strconv.ParseFloat(rows[i][13], 64)
		NetSalary, _ := strconv.ParseFloat(rows[i][14], 64)
		PayDate := rows[i][15]
		sal := salary.Salary{
			CardId:       CardId,
			BaseSalary:   BaseSalary,
			WorkDays:     WorkDays,
			DaysOff:      DaysOff,
			DaysLeave:    DaysLeave,
			Bonus:        Bonus,
			RentSubsidy:  RentSubsidy,
			TransSubsidy: TransSubsidy,
			SocialSec:    SocialSec,
			HouseFund:    HouseFund,
			Tax:          Tax,
			Fine:         Fine,
			NetSalary:    NetSalary,
			PayDate:      PayDate,
			CreateTime:   time.Now(),
		}
		_, err := o.Insert(&sal)
		if err != nil {
			errData = append(errData, sal.CardId)
			continue
		}
	}
	if len(errData) == 0 {
		resp["code"] = 200
		resp["msg"] = "数据导入成功！"
	} else {
		resp["code"] = 1002
		resp["msg"] = "部分数据导入失败！"
		resp["errData"] = errData
	}
	f.Data["json"] = resp
	f.ServeJSON()
}
