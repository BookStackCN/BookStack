package utils

import (
	"io/ioutil"

	"github.com/astaxie/beego"
)

//是否是调试模式
var Debug = true

func init() {
	if beego.AppConfig.String("runmode") == "prod" {
		Debug = false
	}
}

//获取文件内容字符串
//@param			file			文件
//@return			content			返回的字符串内容
//@return			err				错误
func FileGetContent(file string) (content string, err error) {
	var b []byte
	if b, err = ioutil.ReadFile(file); err == nil {
		content = string(b)
	}
	return
}
