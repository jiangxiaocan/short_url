package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/Model"
	"shorturl/public"
)

//创建接口
func Create(c *gin.Context)  {
	//1、接收用户的参数
	url := c.PostForm("url")
	//url := "http://www.baidu.com/20/"+public.RandChar(10)//进行压力测试

	if url == ""{
		c.JSON(http.StatusOK, public.ReturnStruct{
			Status: 0,
			Msg: "url不能为空",
		})
		return
	}

	urlInfoModel := &Model.UrlInfo{
	}
	oldInfo,err := urlInfoModel.GetUrlByUrl(url)

	if err !=nil{
		c.JSON(http.StatusOK, public.ReturnStruct{
			Status: 0,
			Msg: "系统错误:"+err.Error(),
			Data: public.Data{
				Url: url,
				ShortUrl: oldInfo.ShortUrl,
			},
		})
		return
	}

	if oldInfo != nil{
		c.JSON(http.StatusOK, public.ReturnStruct{
			Status: 1,
			Msg: "获取成功",
			Data: public.Data{
				Url: url,
				ShortUrl: oldInfo.ShortUrl,
				VisitUrl: public.HOST+oldInfo.ShortUrl,
			},
		})
		return
	}

	//获取递增数量
	increaseNum,err:= Model.GetIncreaseNum("jxc")

	if err != nil{
		c.JSON(http.StatusOK, public.ReturnStruct{
			Status: 0,
			Msg: "系统异常请稍后再试",
		})
		return
	}

	tailNum := public.GetUrlTableRand(url)
	shortUrl := public.CreateShortUrl(increaseNum,tailNum)

	value,err :=urlInfoModel.InsertUrl(url,shortUrl)

	if value == 0 || err !=nil{
		c.JSON(http.StatusOK, public.ReturnStruct{
			Status: 0,
			Msg: "系统异常请稍后再试",
		})
		return
	}

	c.JSON(http.StatusOK, public.ReturnStruct{
		Status: 1,
		Msg: "获取成功",
		Data: public.Data{
			Url: url,
			ShortUrl: shortUrl,
			VisitUrl: public.HOST+shortUrl,
		},
	})
	return
}
