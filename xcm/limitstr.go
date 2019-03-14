package xcm

import "strings"

var lmtStr = []string{` `, `/`, `%`, `.`, `[`, `]`, `{`, `}`, `(`, `)`, `'`, `"`, `#`, `&`, `^`, `*`, `+`,
	`=`, `\`, `?`, `>`, `<`, `$`, `!`, `|`}

//LimitStr 去掉一些特殊字符
func LimitStr(limitstr string) (str string) {
	str = limitstr
	for _, v := range lmtStr {
		str = strings.Replace(str, v, "【"+v+"】", -1)
	}
	return
}
