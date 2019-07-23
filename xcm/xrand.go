package xcm

import (
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

/*GetRandomInt 生成随机数字
 */
func GetRandomInt(min, max int) (rd int, err error) {
	err = nil
	if max < min {
		err = errors.New("max 必须大于 min")
		return
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rd = r.Intn(max-min-1) + min
	return
}

/*GetRandomSeedInt 生成随机数字
 */
func GetRandomSeedInt(min, max int, seed int64) (rd int, err error) {
	err = nil
	if max < min {
		err = errors.New("max 必须大于 min")
		return
	}
	r := rand.New(rand.NewSource(seed))
	rd = r.Intn(max-min-1) + min
	return
}

/*GetRandomString 生成随机字符串
slen:生成的随机数长度
stp:加密所选择的类型
    NSTR   = iota //数字字符串
	SDSTR         //小写字母字符串
	SUSTR         //大写字母字符串
	SASTR         //大写和小写字母字符串
	NSDSTR        //数字和小写字母字符串
	NSUSTR        //数字和大写字母字符串
	NSASTR        //数字大写和小写字母字符串
	KEYSTR        //数字大写和小写字母字符串(有序)
*/
func GetRandomString(slen int64, stp int) string {
	var mstr string
	switch stp {
	case NSTR:
		mstr = nstr
	case SDSTR:
		mstr = sdstr
	case SUSTR:
		mstr = sustr
	case SASTR:
		mstr = sastr
	case NSDSTR:
		mstr = nsdstr
	case NSUSTR:
		mstr = nsustr
	case KEYSTR:
		mstr = keystr
	default:
		mstr = nsastr
	}
	bytes := []byte(mstr)
	blen := len(bytes)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := int64(0); i < slen; i++ {
		result = append(result, bytes[r.Intn(blen)])
	}
	return string(result)
}

//GetMD5 生成32位的MD5
func GetMD5(str string) (md5str string) {
	md5 := md5.New()
	io.WriteString(md5, str)
	md5str = fmt.Sprintf("%x", md5.Sum(nil))
	return
}

/*GetNumStr 生成数字加密字符
nm:需要加密的字符串
stp:加密所选择的类型
    NSTR   = iota //数字字符串
	SDSTR         //小写字母字符串
	SUSTR         //大写字母字符串
	SASTR         //大写和小写字母字符串
	NSDSTR        //数字和小写字母字符串
	NSUSTR        //数字和大写字母字符串
	NSASTR        //数字大写和小写字母字符串
	KEYSTR        //数字大写和小写字母字符串(有序)
salt:噪声
*/
func GetNumStr(nm string, stp int, salt int) (str string) {
	var mstr string
	switch stp {
	case NSTR:
		mstr = nstr
	case SDSTR:
		mstr = sdstr
	case SUSTR:
		mstr = sustr
	case SASTR:
		mstr = sastr
	case NSDSTR:
		mstr = nsdstr
	case NSUSTR:
		mstr = nsustr
	case KEYSTR:
		mstr = keystr
	default:
		mstr = nsastr
	}
	bytes := []byte(mstr)
	blen := len(bytes)
	result := []byte{}
	for _, i := range nm {
		result = append(result, bytes[(num[string(i)]+salt)%blen])
	}
	str = string(result)
	return
}

//GetSHA256 生成256位的SHA256
func GetSHA256(str string) (sha256str string) {
	sha256 := sha256.New()
	io.WriteString(sha256, str)
	sha256str = fmt.Sprintf("%x", sha256.Sum(nil))
	return
}

/*GetNumCodeStr 对小于62的数进行1位编码
0->"0"   1->"1"   2->"2"   3->"3"   4->"4"   5->"5"   6->"6"   7->"7"   8->"8"   9->"9"   10->"a"  11->"b"  12->"c"  13->"d"  14->"e"
15->"f"  16->"g"  17->"h"  18->"i"  19->"j"  20->"k"  21->"l"  22->"m"  23->"n"  24->"o"  25->"p"  26->"q"  27->"r"  28->"s"  29->"t"
30->"u"  31->"v"  32->"w"  33->"x"  34->"y"  35->"z"  36->"A"  37->"B"  38->"C"  39->"D"  40->"E"  41->"F"  42->"G"  43->"H"  44->"I"
45->"J"  46->"K"  47->"L"  48->"M"  49->"N"  50->"O"  51->"P"  52->"Q"  53->"R"  54->"S"  55->"T"  56->"U"  57->"V"  58->"W"  59->"X"
60->"Y"  61->"Z"
*/
func GetNumCodeStr(n []int) (NumStr string, err error) {
	bytes := []byte(codestr)
	err = nil
	for _, v := range n {
		if v > 61 {
			err = errors.New(strconv.Itoa(v) + "大于61")
			NumStr = ""
			return
		}
		NumStr += string(bytes[v])
	}
	return
}

//GetRandomID 生成随机数ID
func GetRandomID(s1, s2 string, min, max int) string {
	rd, _ := GetRandomInt(min, max)
	return s1 + s2 + strconv.Itoa(rd)
}

//NumCodeStr 对小于62的数进行1位编码
func NumCodeStr(n []int) (NumStr string, err error) {
	bytes := []byte(keystr)
	err = nil
	for _, v := range n {
		if v > 61 {
			err = errors.New("number too big")
			NumStr = ""
			return
		}
		NumStr += string(bytes[v])
	}
	return
}
