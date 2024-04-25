package service

import "github.com/apibrew/apibrew/pkg/model"

type Container interface {
	GetRecordService() RecordService
	GetEventChannelService() EventChannelService
	GetAuthenticationService() AuthenticationService
	GetAuthorizationService() AuthorizationService
	GetResourceService() ResourceService
	GetResourceMigrationService() ResourceMigrationService
	GetDataSourceService() DataSourceService
	GetWatchService() WatchService
	GetExtensionService() ExtensionService
	GetBackendProviderService() BackendProviderService
	GetBackendEventHandler() interface{}
	GetAppConfig() *model.AppConfig
}
