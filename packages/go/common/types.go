package common

type KeyValue struct {
	key   string
	value string
}

func NewKeyValue(key string, value string) KeyValue {
	return KeyValue{key: key, value: value}
}

func Key(key string) func(string) KeyValue {
	return func(value string) KeyValue {
		return NewKeyValue(key, value)
	}
}