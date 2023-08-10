package service

type Container interface {
	GetRecordService() RecordService
	GetAuthenticationService() AuthenticationService
	GetResourceService() ResourceService
	GetMetricsService() MetricsService
	GetResourceMigrationService() ResourceMigrationService
	GetDataSourceService() DataSourceService
	GetWatchService() WatchService
	GetExtensionService() ExtensionService
}
