package Model

import (
	"database/sql"
	"shorturl/public"
	"strconv"
	"time"
)

type UrlInfo struct {
	ID int
	Url string
	UrlMd5 string
	Createtime int
	ShortUrl string
}

//获取数据库的值
func(urlinfo * UrlInfo) GetUrlByMd5Url(urlMd5 string,tableTailNum string,oidId int)( *UrlInfo, error) {
	//写 sql 语句
	tablename := public.TABLE_BASE_NAME+tableTailNum
	sqlStr := "select * from "+tablename+" where url_md5 = ? AND id !=?"

	//执行 sql
	row := Db.QueryRow(sqlStr,urlMd5,oidId)

	//声明三个变量
	var id int
	var urlstring string
	var urlmd5 string
	var createtime int
	var short_url string
	//将各个字段中的值读到以上三个变量中
	err := row.Scan( &id, &urlstring, & urlmd5, & createtime, &short_url)

	//没有数据
	if err == sql.ErrNoRows{
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	//将三个变量的值赋给 User 结构体
	u := &UrlInfo {
			ID: id,
			Url: urlstring,
			UrlMd5: urlmd5,
			Createtime: createtime,
			ShortUrl: short_url,
		}
	return u,
		nil
}

//根据url获取原来的记录
func(urlinfo * UrlInfo) GetUrlByUrl(url string)(*UrlInfo,error){
	tailNum := public.GetUrlTableRand(url)
	md5 := public.GetStringMd5(url)
	var info *UrlInfo
	var err error
	info,err =urlinfo.GetUrlByMd5Url(md5,strconv.Itoa(int(tailNum)),0)

	if err != nil{
		return nil,err
	}

	if info != nil && info.Url != url{
		info,err = urlinfo.GetUrlByMd5Url(md5,strconv.Itoa(int(tailNum)),info.ID)//再次寻找
	}

	return info,err
}

//插入一条数据记录
func (urlinfo * UrlInfo) InsertUrl(url string,shorturl string) (int64,error){
	tailNum := public.GetUrlTableRand(url)
	md5 := public.GetStringMd5(url)

	tablename := public.TABLE_BASE_NAME+strconv.Itoa(int(tailNum))

	stmt, err := Db.Prepare("INSERT "+tablename+" (url,url_md5,createtime,short_url) values (?,?,?,?)")
	if err != nil{
		return 0,err
	}
	timeUnix:=time.Now().Unix()
	res, err := stmt.Exec(url, md5, timeUnix,shorturl)

	if err != nil{
		return 0,err
	}

	return res.LastInsertId()
}

//获取数据库的值
func(urlinfo * UrlInfo) GetUrlByShortUrl(shortUrl string,tableTailNum string)( *UrlInfo, error) {
	//写 sql 语句
	tablename := public.TABLE_BASE_NAME+tableTailNum
	sqlStr := "select * from "+tablename+" where short_url = ?"

	//执行 sql
	row := Db.QueryRow(sqlStr,shortUrl)

	//声明三个变量
	var id int
	var urlstring string
	var urlmd5 string
	var createtime int
	var short_url string
	//将各个字段中的值读到以上三个变量中
	err := row.Scan( &id, &urlstring, & urlmd5, & createtime, &short_url)

	//没有数据
	if err == sql.ErrNoRows{
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	//将三个变量的值赋给 User 结构体
	u := &UrlInfo {
		ID: id,
		Url: urlstring,
		UrlMd5: urlmd5,
		Createtime: createtime,
		ShortUrl: short_url,
	}
	return u,
		nil
}

