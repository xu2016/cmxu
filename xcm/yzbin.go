package xcm

//IsOne 检查64位整数n第i位是否为1
func IsOne(n int64, i int) (b bool) {
	if n>>uint(i)&1 == 1 {
		b = true
	}
	return
}
