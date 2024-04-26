package ctype

import "github.com/goccy/go-json"

type SignStatus int

const (
	SignQQ     SignStatus = iota + 1 // QQ注册
	SignGitee                        // Gitee注册
	SignWechat                       // 微信注册
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	switch s {
	case SignQQ:
		return "管理员"

	case SignGitee:
		return "普通用户"

	case SignWechat:
		return "游客"

	default:
		return "未知注册渠道"
	}
}
