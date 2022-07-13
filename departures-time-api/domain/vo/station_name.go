package vo

// 駅名 1文字以上
type StationName string

// 駅名 1文字以上
// 0文字以下: ErrMinLength
func NewStationName(v string) (*StationName, error) {
	if len(v) <= 0 {
		return nil, ErrMinLength
	}
	return pointer(StationName(v)), nil
}

func (v StationName) Value() string {
	return string(v)
}
