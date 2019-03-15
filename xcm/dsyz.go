package xcm

//XCity 广东地市
var XCity = []string{`省公司`, `广州`, `深圳`, `东莞`, `佛山`, `惠州`, `珠海`, `中山`, `江门`, `汕头`, `湛江`, `揭阳`, `肇庆`, `清远`, `韶关`, `潮州`, `茂名`, `河源`, `汕尾`}

//DSYZ 验证地市输入是否正确
func DSYZ(city string) (b bool) {
	for _, v := range XCity {
		if city == v {
			b = true
		}
	}
	return
}
