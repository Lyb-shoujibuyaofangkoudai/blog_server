package ctype

import "github.com/goccy/go-json"

type Role int

const (
	PermissionAdmin         Role = iota + 1 // 管理员
	PermissionUser                          // 普通用户
	PermissionVisitor                       // 游客
	PermissionSuperDisabled                 // 被禁用的用户
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	switch s {
	case PermissionAdmin:
		return "管理员"

	case PermissionUser:
		return "普通用户"

	case PermissionVisitor:
		return "游客"

	case PermissionSuperDisabled:
		return "被禁用"

	default:
		return "其他"
	}
}
