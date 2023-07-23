package constants

type ContextKey struct{ Name string }

var QueryContextKey = &ContextKey{"query"}
var RedisContextKey = &ContextKey{"redis"}
