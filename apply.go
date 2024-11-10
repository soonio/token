package token

type Apply[T any] func(c *Token[T]) *Token[T]

func WithPrefix[T any](v string) Apply[T] {
	return func(c *Token[T]) *Token[T] {
		c.prefix = v
		return c
	}
}

func WithLength[T any](v int64) Apply[T] {
	return func(c *Token[T]) *Token[T] {
		c.length = v
		return c
	}
}

func WithTokExpired[T any](v int64) Apply[T] {
	return func(c *Token[T]) *Token[T] {
		c.tokExp = v
		return c
	}
}
func WithRefExpired[T any](v int64) Apply[T] {
	return func(c *Token[T]) *Token[T] {
		c.refExp = v
		return c
	}
}

func WithEncoder[T any](v Encoder[T]) Apply[T] {
	return func(c *Token[T]) *Token[T] {
		c.encoder = v
		return c
	}
}
