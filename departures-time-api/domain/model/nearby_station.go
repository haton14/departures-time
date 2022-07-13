package model

import "github.com/haton14/departures-time/departures-time-api/domain/vo"

// 最寄駅
type NearbyStation struct {
	Name      vo.StationName
	Longitude vo.Longitude
	Latitude  vo.Latitude
	Distance  vo.Distance
}

// 駅が現在地から指定した範囲内(引数)であるか
func (n NearbyStation) WithinRange(rangeDistance vo.Distance) bool {
	return n.Distance.Value() <= rangeDistance.Value()
}
