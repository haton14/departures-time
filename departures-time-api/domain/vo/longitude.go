package vo

// 経度: -180~180
type Longitude float64

// 経度
// -180未満:ErrMinRange, 180より大きい:ErrMaxRange
func NewLongitude(v float64) (*Longitude, error) {
	if v < -180 {
		return nil, ErrMinRange
	}
	if v > 180 {
		return nil, ErrMaxRange
	}
	return pointer(Longitude(v)), nil
}

func (v Longitude) Value() float64 {
	return float64(v)
}
