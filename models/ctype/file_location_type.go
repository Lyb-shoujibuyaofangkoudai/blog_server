package ctype

import "github.com/goccy/go-json"

type FileLocationType int

const (
	Local FileLocationType = iota + 1 // 本地保存
	QiNiu                             // 七牛保存
)

func (s FileLocationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s FileLocationType) String() string {
	switch s {
	case Local:
		return "本地保存"

	case QiNiu:
		return "七牛保存"

	default:
		return "其他渠道保存"
	}
}
