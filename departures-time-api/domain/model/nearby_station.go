package model

import (
	"math"

	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/haton14/departures-time/departures-time-api/external"
)

// 最寄駅
type NearbyStation struct {
	Code      external.StationCode // 駅すぱあとWebサービスの駅コード
	Name      vo.StationName
	Longitude vo.Longitude
	Latitude  vo.Latitude
	Distance  vo.Distance
}

// 駅が現在地から指定した範囲内(引数)であるか
func (n NearbyStation) WithinRange(rangeDistance vo.Distance) bool {
	return n.Distance.Value() <= rangeDistance.Value()
}

// 駅と指定した座標の距離を求める
func (n NearbyStation) Difference(lo vo.Longitude, la vo.Latitude) float64 {
	loDifference := n.Longitude.Value() - lo.Value()
	laDifference := n.Latitude.Value() - la.Value()
	return math.Sqrt(loDifference*loDifference + laDifference*laDifference)
}
