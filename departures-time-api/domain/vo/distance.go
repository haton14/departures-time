package vo

// 現在地から駅までの距離(m)
type Distance int

// 現在地から駅までの距離(m)
// // 0未満:ErrMinRange
func NewDistance(v int) (*Distance, error) {
	if v < 0 {
		return nil, ErrMinRange
	}
	return pointer(Distance(v)), nil
}

func (v Distance) Value() int {
	return int(v)
}
