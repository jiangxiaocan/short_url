package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/Model"
	"shorturl/public"
)

func Vist(c *gin.Context)  {
	short_url := c.Param("short_url")

	bool,msg,_ :=public.GetOriginUrlByShortUrl(short_url)

	if !bool{
		c.JSON(http.StatusOK, public.ReturnStruct{
			Status: 0,
			Msg: msg,
		})
		return
	}

	urlInfoModel := &Model.UrlInfo{
	}

	uriInfo,err := urlInfoModel.GetUrlByShortUrl(short_url,msg)

	if uriInfo == nil{
		c.JSON(http.StatusOK, public.ReturnStruct{
			Status: 0,
			Msg: "非法字符串",
		})
		return
	}

	if err !=nil{
		c.JSON(http.StatusOK, public.ReturnStruct{
			Status: 0,
			Msg: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, public.ReturnStruct{
		Status: 1,
		Msg: "获取成功",
		Data: public.Data{
			Url: uriInfo.Url,
		},
	})

	return
}