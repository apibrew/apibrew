package abs

type Container interface {
	GetRecordService() RecordService
	GetAuthenticationService() AuthenticationService
	GetResourceService() ResourceService
	GetResourceMigrationService() ResourceMigrationService
	GetDataSourceService() DataSourceService
	GetWatchService() WatchService
	GetNamespaceService() NamespaceService
	GetUserService() UserService
	GetExtensionService() ExtensionService
}
