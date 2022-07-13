package vo

import "errors"

var (
	ErrMaxRange  = errors.New("最大値より大きい")
	ErrMinRange  = errors.New("最小値より小さい")
	ErrMaxLength = errors.New("最大長より長い")
	ErrMinLength = errors.New("最小長より短い")
	ErrOther     = errors.New("予期しないエラー")
)
