package token

import "context"

// 当前驱动禁

type Store interface {
	SetCtx(ctx context.Context, key1 string, key2 string, value string, seconds int64, seconds2 int64) (val bool, err error)
	GetCtx(ctx context.Context, key string) (val string, err error)
}
