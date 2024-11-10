package token

import (
	"context"
	"strings"
)

type Token[T any] struct {
	store   Store      // 存储驱动
	prefix  string     // 缓存前缀
	length  int64      // token长度
	tokExp  int64      // 正常Token有效时常 秒
	refExp  int64      // 刷新Token的有效时长 秒
	encoder Encoder[T] // 编码器
}

func NewToken[T any](store Store, opt ...Apply[T]) *Token[T] {
	var t = &Token[T]{
		store:   store,
		prefix:  DefaultTokenKeyPrefix,
		length:  DefaultLength,
		tokExp:  DefaultTokExpired,
		refExp:  DefaultRefExpired,
		encoder: &DefaultEncoder[T]{},
	}

	for _, apply := range opt {
		apply(t)
	}
	return t
}

func (t *Token[T]) Generate(ctx context.Context, data T) (string, string, error) {
	var payload string
	var err error

	if payload, err = t.encoder.Marshal(data); err != nil {
		return "", "", err
	}

	var token = strings.ToUpper(String(t.length))
	var refresh = strings.ToUpper(String(t.length))
	ok, err := t.store.SetCtx(ctx, t.key(token, Tok), t.key(refresh, Ref), payload, t.tokExp, t.refExp)
	if err != nil {
		return "", "", err
	}
	if !ok {
		return "", "", GenerateTokenError
	}

	return token, refresh, nil
}

func (t *Token[T]) Get(ctx context.Context, key string, typ Type) (*T, error) {
	payload, err := t.store.GetCtx(ctx, t.key(key, typ))
	if err != nil {
		return nil, err
	}
	var v = new(T)
	if err = t.encoder.Unmarshal(v, payload); err != nil {
		return nil, err
	}
	return v, nil
}

func (t *Token[T]) Length() int64 {
	return t.length
}

func (t *Token[T]) Expired(typ Type) int64 {
	switch typ {
	case Tok:
		return t.tokExp
	case Ref:
		return t.refExp
	}
	return -1
}

func (t *Token[T]) key(k string, typ Type) string {
	return strings.Join([]string{t.prefix, string(typ), k}, ":")
}
