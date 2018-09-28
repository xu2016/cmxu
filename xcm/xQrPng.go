package xcm

import (
	"errors"
	"image/png"
	"os"

	ebar "github.com/boombuler/barcode"
	eqr "github.com/boombuler/barcode/qr"
)

//CreateQrPng 生成二维码图片
func CreateQrPng(filepath, shopid, wdlink string) (err error) {
	base64 := wdlink + "?shopid=" + shopid
	code, err := eqr.Encode(base64, eqr.L, eqr.Unicode)
	if err != nil {
		return
	}
	if base64 != code.Content() {
		err = errors.New("编码不正确")
		return
	}
	code, err = ebar.Scale(code, 400, 400)
	if err != nil {
		return
	}
	file, err := os.Create(filepath + shopid + ".png")
	defer file.Close()
	if err != nil {
		return
	}
	err = png.Encode(file, code)
	if err != nil {
		return
	}
	err = nil
	return
}
