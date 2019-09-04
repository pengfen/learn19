package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"path"
	"time"
	"wel/models"
)

// 文章控制器
type ArticleController struct {
	beego.Controller
}

// 文章列表
func (c *ArticleController) Index() {
	// 查询文章列表数据
	o := orm.NewOrm()
	var articles []models.Article
	_, err := o.QueryTable("Article").All(&articles)
	if err != nil {
		beego.Info("查询所有文章信息出错")
		return
	}

	// --------------- 处理分页 -----------
	id,_ := c.GetInt("select")
	//获得数据总数，总页数，当前页码
	count,err := o.QueryTable("Article").Count()//RelatedSel("ArticleType").Filter("ArticleType__Id",id).Count()
	if err != nil {
		beego.Info("查询失败",err)
		return
	}
	pagesize := int64(2)  //每页显示数据条目

	index,err := c.GetInt("pageIndex")  //当前页码
	if err != nil{
		index = 1
	}


	pageCount := math.Ceil(float64(count) / float64(pagesize))   //总页数

	if index <=0 || index > int(pageCount){
		index = 1
	}

	start := (int64(index)  -1 ) * pagesize
	// inner   left join
	var artis []models.Article
	o.QueryTable("Article").Limit(pagesize,start).All(&artis)//RelatedSel("ArticleType").Filter("ArticleType__Id",id).All(&artis)
	c.Data["pageCount"] = pageCount
	c.Data["count"] = count
	c.Data["pageIndex"] = index
	c.Data["typeid"] = id    //文章类型ID
	// --------------- 处理分页 -----------

	//获取类型数据
	var artiTypes []models.ArticleType
	_,err = o.QueryTable("ArticleType").All(&artiTypes)
	if err != nil{
		beego.Info("获取类型错误")
		return
	}

	c.Data["articleType"] = artiTypes

	c.Data["artiles"] = articles
	c.TplName = "article/index.html"
}

// 显示添加文章界面
func (c *ArticleController) Add() {
	c.TplName = "article/add.html"
}

// 添加文章逻辑
func (c *ArticleController) HandlerAdd() {
	// 1. 获取参数
	artName := c.GetString("title")
	artContent := c.GetString("content")
	f,h,err := c.GetFile("uploadname")
	defer f.Close()

	//---------------------- 文件上伟注意格式大小等 --------------------
	// 限定格式
	fileext := path.Ext(h.Filename)
	if fileext != ".jpg" && fileext != ".png" {
		beego.Info("上传文件格式错误")
		return
	}
	// 限制大小
	if h.Size > 50000000 {
		beego.Info("上传文件过大")
		return
	}
	// 需要对文件重命名 防止文件名重复
	//filename := time.Now().Format("2006-2-15-15-36-10")
	filename := time.Now().Format("2006-01-02 15:04:05") + fileext
	//---------------------- 文件上伟注意格式大小等 --------------------
	beego.Info(filename)

	if err != nil {
		beego.Info("上传文件失败")
		return
	} else {
		c.SaveToFile("uploadname", "./static/img/" + filename)
	}
	beego.Info(artName, artContent)

	// 2. 判断数据是否合法
	if artName == "" || artContent == "" {
		beego.Info("添加文章数据错误")
		return
	}
	// 3. 插入数据
	o := orm.NewOrm()
	art := models.Article{}
	art.Aname = artName
	art.Acontent = artContent
	art.Atime = time.Now()
	art.Aimg = "/static/img/" + h.Filename
	_,err = o.Insert(&art)
	if err != nil {
		beego.Info("插入数据库数据", err)
		return
	}
	// 4. 返回文章界面
	c.Redirect("/index", 302)
}

// 文章详情
func (c *ArticleController) Detail() {
	// 1 获取文章ID
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("获取文章ID错误", err)
		return
	}

	// 2. 查询数据并获取数据
	o := orm.NewOrm()
	art := models.Article{Id:id}
	err = o.Read(&art)
	if err != nil {
		beego.Info("查询错误", err)
		return
	}

	// 3. 显示详情页面
	c.Data["article"] = art
	c.TplName = "article/detail	.html"
}

// 显示修改页面
func (c *ArticleController) Update() {
	// 1. 获取文章ID
	id,err := c.GetInt("id")
	if err != nil {
		beego.Info("获取文章ID错误", err)
		return
	}

	// 2. 获取对应数据
	o := orm.NewOrm()
	art := models.Article{Id:id}
	err = o.Read(&art)
	if err != nil {
		beego.Info("查询错误", err)
		return
	}

	// 3. 渲染视图
	c.Data["article"] = art
	c.TplName = "article/update.html"
}

// 修改逻辑
func (c *ArticleController) HandlerUpd() {
	//1.拿到数据
	id,_ := c.GetInt("id")
	artiName := c.GetString("articleName")
	content := c.GetString("content")
	f,h,err:=c.GetFile("uploadname")
	var filename string
	if err != nil{
		beego.Info("上传文件失败")
		return
	}else {
		defer f.Close()


		//1.要限定格式
		fileext := path.Ext(h.Filename)
		if fileext != ".jpg" && fileext != "png"{
			beego.Info("上传文件格式错误")
			return
		}
		//2.限制大小
		if h.Size > 50000000 {
			beego.Info("上传文件过大")
			return
		}

		//3.需要对文件重命名，防止文件名重复
		filename = time.Now().Format("2006-01-02 15:04:05") + fileext  //6-1-2 3:4:5
		c.SaveToFile("uploadname","./static/img/"+filename)
	}

	//2.对数据进行一个处理
	if artiName == "" || content ==""{
		beego.Info("更新数据获取失败")
		return
	}

	//3.更新操作
	o := orm.NewOrm()
	art := models.Article{Id:id}
	err = o.Read(&art)
	if err != nil{
		beego.Info("查询数据错误")
		return
	}
	art.Aname = artiName
	art.Acontent = content
	art.Aimg = "./static/img/"+filename


	_,err = o.Update(&art,"ArtiName","Acontent","Aimg")
	if err != nil{
		beego.Info("更新数据显示错误")
		return
	}
	//4.返回列表页面
	c.Redirect("/index",302)
}

// 删除逻辑
//删除操作
func (c *ArticleController) HandleDelete(){
	//1.拿到数据
	id,err:=c.GetInt("id")
	if err != nil{
		beego.Info("获取id数据错误")
		return
	}


	//2.执行删除操作
	o := orm.NewOrm()
	arti := models.Article{Id:id}
	err = o.Read(&arti)
	if err != nil{
		beego.Info("查询错误")
		return
	}
	o.Delete(&arti)

	//3.返回列表页面
	c.Redirect("/index",302)
}