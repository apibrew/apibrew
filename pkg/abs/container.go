package abs

type Container interface {
	GetRecordService() RecordService
	GetAuthenticationService() AuthenticationService
	GetResourceService() ResourceService
	GetDataSourceService() DataSourceService
	GetWatchService() WatchService
	GetNamespaceService() NamespaceService
	GetUserService() UserService
	GetExtensionService() ExtensionService
	GetPluginService() PluginService
}
