package abs

import "github.com/apibrew/apibrew/pkg/model"

type Container interface {
	GetRecordService() RecordService
	GetAuthenticationService() AuthenticationService
	GetResourceService() ResourceService
	GetResourceMigrationService() ResourceMigrationService
	GetDataSourceService() DataSourceService
	GetWatchService() WatchService
	GetNamespaceService() NamespaceService
	GetUserService() GenericRecordService[*model.User]
	GetExtensionService() ExtensionService
}
