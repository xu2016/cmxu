package xcm

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

const (
	nstr    = "0123456789"
	sdstr   = "klmnopqrwxyzabstuvghijcdef"
	sustr   = "KLMNOPQRWXYZABSTUVGHIJCDEF"
	sastr   = "klmnopqrwxyzabstuvghijcdefKLMNOPQRWXYZABSTUVGHIJCDEF"
	nsdstr  = "kl3mn2opqr1wx9yz8abst5uv4gh7ij6cd0ef"
	nsustr  = "KL3MN2OPQR1WX9YZ8ABST5UV4GH7IJ6CD0EF"
	nsastr  = "kl0mnop9qrwxy7zabs5tuvgh8ijcdefK6LMNOP1QRWX3YZABS4TUVGHI2JCDEF"
	codestr = "0123456789klmnopqrwxyzabstuvghijcdefKLMNOPQRWXYZABSTUVGHIJCDEF"
	keystr  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

//设置生成字符串的格式
const (
	NSTR   = iota //数字字符串
	SDSTR         //小写字母字符串
	SUSTR         //大写字母字符串
	SASTR         //大写和小写字母字符串
	NSDSTR        //数字和小写字母字符串
	NSUSTR        //数字和大写字母字符串
	NSASTR        //数字大写和小写字母字符串
	KEYSTR        //数字大写和小写字母字符串(有序)
)

var num = map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}

//ReturnJSON 返回一个格式化好的JSON字符串
func ReturnJSON(cnt interface{}) string {
	data, _ := json.Marshal(cnt)
	return string(data)
}

//Readcfg 读取以“,”分隔的配置文件的前两个字段
func Readcfg(rs string, mycfg map[string]string) {
	f, err := os.Open(rs)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err == nil || err == io.EOF {
			str := strings.Split(line, ",")
			if str[0] != "" {
				if strings.Contains(str[1], "\r\n") {
					mycfg[str[0]] = strings.Replace(str[1], "\r\n", "", -1)
				} else {
					mycfg[str[0]] = strings.Replace(str[1], "\n", "", -1)
				}
			}
			if err == io.EOF {
				break
			}
		} else {
			break
		}
	}
}
