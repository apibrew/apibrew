package abs

type contextKey struct{}

var SystemContextKey = contextKey{}
var UserContextKey = contextKey{}
var TransactionContextKey = contextKey{}
var ClientTrackIdContextKey = contextKey{}
