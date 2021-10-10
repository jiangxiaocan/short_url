package admin

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestPublishUnreachable(t *testing.T) {
	api := "http://localhost:8080/admin/create"
	for i:=0;i<10000;i++{
		err := GetInfo(api)
		if err != nil {
			t.Errorf("GetInfo() return an error %s",err)
		}
	}

}


func GetInfo(api string) ( error) {
	resp, _ := http.Post(api, "application/x-www-form-urlencoded", strings.NewReader("url=http://www.baidu.com/"+RandChar(3)))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("get info didn’t respond 200 OK: %s", resp.Status)
	}

	return nil
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