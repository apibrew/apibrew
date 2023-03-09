package abs

type contextKey string

var SystemContextKey = contextKey("SystemContextKey")
var UserContextKey = contextKey("UserContextKey")
var TransactionContextKey = contextKey("TransactionContextKey")
var ClientTrackIdContextKey = contextKey("ClientTrackIdContextKey")
var LogFieldsContextKey = contextKey("LogFieldsContextKey")
