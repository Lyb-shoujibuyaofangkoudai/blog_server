package ctype

import "github.com/goccy/go-json"

type SignStatus int

const (
	SignQQ     SignStatus = iota + 1 // QQ注册
	SignGitee                        // Gitee注册
	SignWechat                       // 微信注册
	SignEmail                        // 邮箱注册
	SignPhone                        // 手机注册
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	switch s {
	case SignQQ:
		return "QQ注册"

	case SignGitee:
		return "gitee注册"

	case SignWechat:
		return "微信注册"
	case SignPhone:
		return "手机注册"
	case SignEmail:
		return "邮箱注册"

	default:
		return "未知注册渠道"
	}
}
