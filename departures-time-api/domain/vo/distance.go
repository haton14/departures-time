package vo

// 現在地から駅までの距離(km,小数点2桁まで)
type Distance float64

// 現在地から駅までの距離(km,小数点2桁まで)
// // 0未満:ErrMinRange, 20より大きい:ErrMaxRange
func NewDistance(v float64) (*Distance, error) {
	if v < 0 {
		return nil, ErrMinRange
	}
	if v > 20 {
		return nil, ErrMaxRange
	}
	return pointer(Distance(v)), nil
}

func NewDistanceForMeter(vMeter int) (*Distance, error) {
	return NewDistance(float64(vMeter) / 1000)
}

func (v Distance) Value() float64 {
	return float64(v)
}
