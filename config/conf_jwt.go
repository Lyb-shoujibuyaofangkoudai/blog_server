package config

type Jwt struct {
	Expires    int    `yaml:"expires"`     // 过期时间 单位小时
	Issuer     string `yaml:"issuer"`      // 颁发人
	GrantScope string `yaml:"grant_scope"` // 授权范围
	Subject    string `yaml:"subject"`     // 主题
}

const (
	PRI_KEY = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDOEYGlPZuYBVDRFe6gXl/03MJE2q+k1/+PcVzubWAlF6RUHHvl
Om4IQcQH1g3D1mQT6QX8RQL7cPgcLnGFZ7qaTiJfKN8kexuS6sETlNcYOIMPl6zA
grxsa8k5kM16cCK2S8xeKHCNNBD+rUWDHflZYsQvvDEg1Ys+GCauYajLxQIDAQAB
AoGAXpcUnsgX2wFdpoxdvAl2HI0VM8v6Yj2wFqUf1mYogv5GNUHZ8VAP4ARoOnyc
Vu/bgnQthi4bf1XM3grHm0gRFB6eyDgdR+UbiyxGZn7kpgtwkYWycMhg8YPJm+HV
UIh2Rvy2JkCzDNUUSk1cw4bcIUqlxhB1t886vhHJ5rAyFoUCQQDtPV3JHyVdoWxo
PI/xeA+gLY8IohMECZKz0cSKzuhVGWOmX58hLUwbmXRladPMJLH6pLR5IY0dVX2Q
e1SYnzlXAkEA3l0dKMxLCUM+LVaY6yz0OMxBsdarmUkZp6tZLyM4Ii2ArP0j/5E5
lr4gKHcIMMpDHummqO7yvYYdsMwih8lGQwJAY1eGFTkAmZOF5KQvlmqzCFzrfy73
DYLAtqHJTmLT8QafrsRtyyO/sfLxRaIp+VsIWC9uDycYg0cQPFcYloxeIwJBAKK3
NoBFRl9n0lbw+IOXaLsrVKNjODy6DkjwjRl+RzRTYca0kqQQTDjvta6Gs/qn94fm
aGtUN1LSkmVua5I7iesCQQCGaDRm5miWNQafoZ+RV3BRvoZ0tAOUHNMOhOWNNZmv
/yGj+Uzbjg6N7+ulsUXcIXZtIDaXjL9WR4gk/wBXdMFr
-----END RSA PRIVATE KEY-----
`
	PUB_KEY = `-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAM4RgaU9m5gFUNEV7qBeX/TcwkTar6TX/49xXO5tYCUXpFQce+U6bghB
xAfWDcPWZBPpBfxFAvtw+BwucYVnuppOIl8o3yR7G5LqwROU1xg4gw+XrMCCvGxr
yTmQzXpwIrZLzF4ocI00EP6tRYMd+VlixC+8MSDViz4YJq5hqMvFAgMBAAE=
-----END RSA PUBLIC KEY-----
`
)
