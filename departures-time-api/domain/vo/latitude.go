package vo

// 緯度: -180~180
type Latitude float64

// 緯度
// -180未満:ErrMinRange, 180より大きい:ErrMaxRange
func NewLatitude(v float64) (*Latitude, error) {
	if v < -180 {
		return nil, ErrMinRange
	}
	if v > 180 {
		return nil, ErrMaxRange
	}
	return pointer(Latitude(v)), nil
}

func (v Latitude) Value() float64 {
	return float64(v)
}
