package abs

type contextKey string

var SystemContextKey = contextKey("SystemContextKey")
var UserContextKey = contextKey("UserContextKey")
var ClientTrackIdContextKey = contextKey("ClientTrackIdContextKey")
var LogFieldsContextKey = contextKey("LogFieldsContextKey")
