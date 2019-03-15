package xcm

import "strings"

var lmtStr = []string{` `, `/`, `%`, `.`, `[`, `]`, `{`, `}`, `(`, `)`, `'`, `"`, `#`, `&`, `^`, `*`, `+`,
	`=`, `\`, `?`, `>`, `<`, `$`, `!`}

//LimitStr 把特殊字符替换成startstr+特殊字符+endstr
func LimitStr(limitstr, startstr, endstr string) (str string) {
	str = limitstr
	for _, v := range lmtStr {
		str = strings.Replace(str, v, startstr+v+endstr, -1)
	}
	return
}

//LimitStrReplaceBlank 去特殊字符
func LimitStrReplaceBlank(limitstr string) (str string) {
	str = limitstr
	for _, v := range lmtStr {
		str = strings.Replace(str, v, "", -1)
	}
	return
}
