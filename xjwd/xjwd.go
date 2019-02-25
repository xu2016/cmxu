package xjwd

import "math"

//ER 地球半径
const ER = 6378137.0 // 6371000

//Rad ...
const Rad = math.Pi / 180.0

//EarthDistance 地球上两个经纬度之间的直线距离,单位：米
func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	lat1 = lat1 * Rad
	lat2 = lat2 * Rad
	theta := lng2*Rad - lng1*Rad
	return (math.Acos(math.Sin(lat1)*math.Sin(lat2)+math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta)) * ER) / 1000
}
