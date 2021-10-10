package public

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/Unknwon/goconfig"
	"hash/crc32"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)
/**
	计算字符串的crc32位
 */
func GetCrc32Key(strKey string) uint32 {
	table := crc32.MakeTable(crc32.IEEE)
	ret := crc32.Checksum([]byte(strKey), table)
	return ret
}

//获得url属于哪个表
func GetUrlTableRand(url string) uint32{
	crc32 := GetCrc32Key(url)

	feat := crc32%TABLE_NUM

	return feat
}

func GetStringMd5(s string) string {
	md5 := md5.New()
	md5.Write([]byte(s))
	md5Str := hex.EncodeToString(md5.Sum(nil))
	return md5Str
}

var chars string = "824356719ABCDEFGHIJKLMONPQRSTVUWXYZabcdefghikjlmnopqrstuvwxyz"

/**
	将数字转成61位
 */
func Encode61(num int64) string {
	bytes := []byte{}
	for num > 0 {
		bytes = append(bytes, chars[num%61])
		num = num / 61
	}
	reverse(bytes)
	return string(bytes)
}

/**
	将61位数字转成字符串
 */
func Decode61(str string) int64 {
	var num int64
	n := len(str)
	for i := 0; i < n; i++ {
		if str[i] == '0'{
			continue
		}
		pos := strings.IndexByte(chars, str[i])
		num += int64(math.Pow(61, float64(n-i-1)) * float64(pos))
	}
	return num
}

func reverse(a []byte) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}

//补多少个0
func GetPos(pos string,num string) string{
	newStr:=fmt.Sprintf("%0"+num+"s", pos)
	return newStr
}

//去除
func TrimPos(pos string)  string{
	return strings.TrimLeft(pos, "0")
}

var cfg *goconfig.ConfigFile

func GetConfigIni(filepath string) (err error) {
	config, err := goconfig.LoadConfigFile(filepath)
	if err != nil {
		fmt.Println("配置文件读取错误,找不到配置文件", err)
		return err
	}
	cfg = config
	return nil
}

func GetDatabase() (types, local, online string, err error) {
	if types, err = cfg.GetValue("database", "types"); err != nil {
		fmt.Println("配置文件中不存在types", err)
		return types, local, online, err
	}
	if local, err = cfg.GetValue("database", "local"); err != nil {
		fmt.Println("配置文件中不存在local", err)
		return types, local, online, err
	}
	if online, err = cfg.GetValue("database", "online"); err != nil {
		fmt.Println("配置文件中不存在online", err)
		return types, local, online, err
	}
	return types, local, online, nil
}

func GetSelf() (port ,host string, err error) {
	if port, err = cfg.GetValue("self", "port"); err != nil {
		fmt.Println("配置文件中不存在port", err)
		return port, host, err
	}

	host, err = cfg.GetValue("self", "host")
	if err != nil {
		fmt.Println("配置文件中不存在tag", err)
		return port, host, err
	}

	return port, host, err

}

//创建短链接
func CreateShortUrl(num int64,tableNum uint32)string{
	base61 := Encode61(num)
	middle := GetPos(base61,"5")//5位代表一张表的数量

	//得出url在哪张表上面
	pre := GetPos(Encode61(int64(tableNum)),"2")//前面2位得出短链接在哪个位置

	end := GetStringMd5(pre+middle+SHORT_URL_CHECKCODE)
	end = end[len(end)-1:]
	
	return pre+middle+end
}

/**
	获取原来的短链接
 */
func GetOriginUrlByShortUrl(shorturl string)(bool,string,error){
	if len(shorturl)<8{
		return false,"字符串长度错误",nil
	}
	end := shorturl[len(shorturl)-1:]
	preNoCheck := shorturl[:len(shorturl)-1]
	checkCode := GetStringMd5(preNoCheck+SHORT_URL_CHECKCODE)
	checkCodeEnd := checkCode[len(checkCode)-1:]
	if checkCodeEnd != end{
		return false,"非法短链接",nil
	}

	//得到前两位数据库位置
	tableNum := shorturl[0:2]
	tableNumInt := Decode61(tableNum)

	return true,strconv.Itoa(int(tableNumInt)),nil
}

const char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandChar(size int) string {
	rand.NewSource(time.Now().UnixNano()) // 产生随机种子
	var s bytes.Buffer
	for i := 0; i < size; i ++ {
		s.WriteByte(char[rand.Int63() % int64(len(char))])
	}
	return s.String()
}