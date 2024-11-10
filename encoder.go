package token

import "encoding/json"

type Encoder[T any] interface {
	Marshal(any) (string, error) //
	Unmarshal(*T, string) error  //
}

type DefaultEncoder[T any] struct {
}

func (e *DefaultEncoder[T]) Marshal(data any) (string, error) {
	bs, err := json.Marshal(data)
	return string(bs), err
}

func (e *DefaultEncoder[T]) Unmarshal(data *T, raw string) error {
	return json.Unmarshal([]byte(raw), data)
}
