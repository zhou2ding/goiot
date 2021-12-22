package utils

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	Id   string
	BS64 string //验证码图片的类型是Base64
	Code int
}

var store = base64Captcha.DefaultMemStore

func GetCaptcha() (id, base64 string, err error) {
	bg := color.RGBA{0, 0, 0, 0}
	fonts := []string{"wqy-microhei.ttc"}
	//生成driver，除了math还有audio、Chinese、digit、language、string风格
	//NewDriverMath接收高、宽，背景文字的干扰，干扰线的条数，背景颜色的指针，字体的列表（切片）
	driver := base64Captcha.NewDriverMath(50, 140, 5, 3, &bg, fonts)
	//使用store和driver生成验证码实例
	captcha := base64Captcha.NewCaptcha(driver, store)
	//通过实例生成验证码
	id, base64, err = captcha.Generate()
	return
}

func VerifyCaptcha(id, answer string) bool {
	//第三个参数为是否从缓存中清除验证码
	return store.Verify(id, answer, true)
}
