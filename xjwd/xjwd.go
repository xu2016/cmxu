package xjwd

import "math"

//ER 地球半径
const ER = 6378137.0 // 6371000

//Rad ...
const Rad = math.Pi / 180.0

//ARad ...
const ARad = 180.0 / math.Pi

//EarthDistance 地球上两个经纬度之间的直线距离,单位：米
func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	lat1 = lat1 * Rad
	lat2 = lat2 * Rad
	theta := lng2*Rad - lng1*Rad
	return (math.Acos(math.Sin(lat1)*math.Sin(lat2)+math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta)) * ER) / 1000
}

//LngLatRange 获取距离经纬度为(lng,lat)的点距离为d的圆的外接正方形上下经纬度边界
func LngLatRange(lat, lng, d float64) (minLat, maxLat, maxLng, minLng float64) {
	rangex := ARad * d / 6372797.0 //里面的 d就代表搜索 dm 之内，单位m
	lngR := rangex / math.Cos(lat*Rad)
	maxLat = lat + rangex //最大纬度
	minLat = lat - rangex //最小纬度
	maxLng = lng + lngR   //最大经度
	minLng = lng - lngR   //最小经度
	return
}
