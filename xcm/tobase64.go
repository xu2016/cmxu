package xcm

import (
	"encoding/base64"
)

//ToBase64 转换成base64
func ToBase64(file []byte) (b64 []byte) {
	b64 = make([]byte, 5000000)          //数据缓存
	base64.StdEncoding.Encode(b64, file) // 文件转base64
	return
}
