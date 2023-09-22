package service

type Container interface {
	GetRecordService() RecordService
	GetEventChannelService() EventChannelService
	GetAuthenticationService() AuthenticationService
	GetAuthorizationService() AuthorizationService
	GetResourceService() ResourceService
	GetMetricsService() MetricsService
	GetResourceMigrationService() ResourceMigrationService
	GetDataSourceService() DataSourceService
	GetWatchService() WatchService
	GetExtensionService() ExtensionService
}
