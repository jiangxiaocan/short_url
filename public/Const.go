package public
//前面两位是数据表 中间5位  后面1位

const TABLE_NUM = 4
const TABLE_BASE_NAME = "short_url_"
const SHORT_URL_CHECKCODE = "3E32DSFGS#%%^~"

var HOST string

//结构体返回，灵活使用tag来对结构体字段做定制化操作
type ReturnStruct struct {
	Status    int
	Msg string
	Data     interface{}
}

type Data struct {
	Url string
	ShortUrl string `json:"short_url"`
	VisitUrl string `json:"visit_url"`
}