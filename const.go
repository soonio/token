package token

const (
	DefaultTokenKeyPrefix = "token" // 默认缓存前缀
	DefaultLength         = 32      // 默认长度
	DefaultTokExpired     = 86400   // 默认有效时常
	DefaultRefExpired     = 129600  // 默认有效时常
)

type Type string

const Tok Type = "tok"
const Ref Type = "ref"
