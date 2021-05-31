package finance

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"math"
	"ions/models/profit"
	"ions/utils"
	"strconv"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego/logs"
)

type CaiWuEchartDataController struct {
	beego.Controller
}

func (c *CaiWuEchartDataController) Get()  {


	o := orm.NewOrm()
	qs := o.QueryTable("sys_caiwu_data")
	// 每页显示的条数
	pagePerNum := 8
	// 当前页
	currentPage,err := c.GetInt("page")
	if err != nil {   // 说明没有获取到当前页
		currentPage = 1
	}

	offsetNum := pagePerNum * (currentPage - 1)


	month := c.GetString("month")
	var count int64 = 0
	var caiwuDatas []profit.Profit
	if month != ""{   // 有查询条件的
		// 总数
		count,_ = qs.Filter("mon",month).Count()
		_,_ = qs.Filter("mon",month).Limit(pagePerNum).Offset(offsetNum).All(&caiwuDatas)
	}else {
		month = time.Now().Format("2006-01")

		count,_ = qs.Filter("mon",month).Count()

		_,_ = qs.Filter("mon",month).Limit(pagePerNum).Offset(offsetNum).All(&caiwuDatas)

	}

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

	c.Data["caiwu_datas"] = caiwuDatas
	c.Data["prePage"] =prePage
	c.Data["nextPage"] = nextPage
	c.Data["currentPage"] = currentPage
	c.Data["countPage"] = countPage
	c.Data["count"] = count
	c.Data["page_map"] = pageMap
	c.Data["month"] = month

	c.TplName = "finance/echart_data_list.html"

}

func (c *CaiWuEchartDataController) ToImportExcel()  {
	c.TplName = "finance/echart_data_import.html"

}


func (c *CaiWuEchartDataController) DoImportExcel()  {
	f,h,err := c.GetFile("upload_file")

	messageMap := map[string]interface{}{}
	var errDataArr []string

	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		messageMap["code"] = 10001
		messageMap["msg"] = "文件上传失败"
		c.Data["json"] = messageMap
		c.ServeJSON()
		return
	}

	fileName := h.Filename

	timeUnixInt := time.Now().Unix()
	timeUnitStr := strconv.FormatInt(timeUnixInt,10)

	filePath := "static/upload/echart_data_upload/"+ timeUnitStr + "-" + fileName

	_ = c.SaveToFile("upload_file", filePath)


	// 读取数据并插入数据库
	file,err1 := excelize.OpenFile(filePath)
	logs.Error(err1)
	rows := file.GetRows("Sheet1")

	o := orm.NewOrm()


	i := 0
	for _,row := range rows {
		caiwuDate := row[0]
		salesVolume,_ := strconv.ParseFloat(row[1],64)
		studentIncress,_ := strconv.Atoi(row[2])
		django,_ := strconv.Atoi(row[3])
		vueDjango,_ := strconv.Atoi(row[4])
		celery,_ := strconv.Atoi(row[5])

		echartData := profit.Profit{
			Mon:       caiwuDate,
			Sales:     salesVolume,
			StuIncr:   studentIncress,
			Django:    django,
			VueDjango: vueDjango,
			Celery:    celery,
		}

		if i == 0 {
			i ++
			continue
		}


		// 重复导入相同月份的数据：先删除已有的工资月份，再导入
		qs := o.QueryTable("sys_caiwu_data")
		isExist := qs.Filter("mon", caiwuDate).Exist()

		if isExist {
			_,_ = qs.Filter("mon", caiwuDate).Delete()
		}



		// 精确到导入失败的数据信息提示
		_,err := o.Insert(&echartData)

		if err != nil {  // 报错的数据
			errDataArr = append(errDataArr, caiwuDate)
		}
		i ++

	}

	if len(errDataArr) <= 0 {
		messageMap["code"] = 200
		messageMap["msg"] = "导入成功"
	} else {
		messageMap["code"] = 10002
		messageMap["msg"] = "导入失败"
		messageMap["err_data"] = errDataArr

	}

	c.Data["json"] = messageMap
	c.ServeJSON()

}
