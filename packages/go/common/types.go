package common

type KeyValue struct {
	key string
	value string
}

func NewKeyValue(key string, value string) KeyValue {
	return KeyValue{key: key, value: value}
}

type Key string

func (k Key) Value(val string) KeyValue {
	return NewKeyValue(string(k), val)
}
