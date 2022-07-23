package model

import (
	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/haton14/departures-time/departures-time-api/external"
)

// 目的駅
type Destination struct {
	Code      external.StationCode // 駅すぱあとWebサービスの駅コード
	Name      vo.StationName
	Longitude vo.Longitude
	Latitude  vo.Latitude
}
