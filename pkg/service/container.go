package service

type Container interface {
	GetRecordService() RecordService
	GetAuthenticationService() AuthenticationService
	GetResourceService() ResourceService
	GetResourceMigrationService() ResourceMigrationService
	GetDataSourceService() DataSourceService
	GetWatchService() WatchService
	GetExtensionService() ExtensionService
}
