package strategies

type Context interface {
	String(key string) string
	Bool(key string) bool
}
