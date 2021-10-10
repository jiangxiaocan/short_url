package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"shorturl/Model"
	"shorturl/admin"
	"shorturl/home"
	"shorturl/public"
	"time"
)

func main()  {
	router := gin.Default()

	//读取数据库配置
	err := public.GetConfigIni("./conf/conf.ini")
	if err != nil {
		panic(err)
	}
	datatype,local,_,_ := public.GetDatabase()

	//初始化数据库
	Model.Db, err = sql.Open(datatype, local)
	if err !=nil{
		panic(err)
	}

	Model.Db.SetConnMaxLifetime(time.Duration(100) * time.Second)
	Model.Db.SetMaxIdleConns(10)

	var port string
	port,public.HOST, err =  public.GetSelf()
	if err !=nil{
		panic(err)
	}

	//访问后台,创建URl
	v1 := router.Group("admin")
	{
		//访问为http://127.0.0.1:8080/admin/create
		v1.POST("/create", admin.Create)
	}

	router.GET("/:short_url", home.Vist)

	router.Run(":"+port) // listen and serve on 0.0.0.0:8080
}
