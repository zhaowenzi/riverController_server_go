package util

import "time"

func GenerateCurrentTimeStamp() int64 {
	return time.Now().Unix()
}

func GetRiverTypeFromId(id string) string {
	if id == "0" {
		return "河道"
	}
	if id == "1" {
		return "水库"
	}
	if id == "2" {
		return "排污口"
	}
	if id == "3" {
		return "管网"
	}
	if id == "4" {
		return "楼宇蓄水池"
	}
	if id == "5" {
		return "养殖鱼塘"
	}
	return ""
}
