package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/model"
	"github.com/tislib/data-handler/service/errors"
	"github.com/tislib/data-handler/service/mapping"
	"github.com/tislib/data-handler/service/params"
	"github.com/tislib/data-handler/service/security"
	"github.com/tislib/data-handler/service/system"
)

type UserService interface {
	InjectRecordService(service RecordService)
	InjectResourceService(service ResourceService)
	Init(data *model.InitData)
	Create(ctx context.Context, users []*model.User) ([]*model.User, errors.ServiceError)
	Update(ctx context.Context, users []*model.User) ([]*model.User, errors.ServiceError)
	Delete(ctx context.Context, ids []string) errors.ServiceError
	Get(ctx context.Context, id string) (*model.User, errors.ServiceError)
	List(ctx context.Context, query *model.BooleanExpression, limit uint32, offset uint64) ([]*model.User, errors.ServiceError)
	InjectBackendProviderService(service BackendProviderService)
}

type userService struct {
	recordService          RecordService
	authenticationService  AuthenticationService
	serviceName            string
	resourceService        ResourceService
	backendProviderService BackendProviderService
}

func (u *userService) InjectBackendProviderService(backendProviderService BackendProviderService) {
	u.backendProviderService = backendProviderService
}

func (u *userService) InjectResourceService(service ResourceService) {
	u.resourceService = service
}

func (u *userService) InjectRecordService(service RecordService) {
	u.recordService = service
}

func (u *userService) Create(ctx context.Context, users []*model.User) ([]*model.User, errors.ServiceError) {
	u.encodePasswords(users)

	// insert records via resource service
	records := mapping.MapToRecord(users, mapping.UserToRecord)

	result, _, err := u.recordService.Create(ctx, params.RecordCreateParams{
		Namespace: system.UserResource.Namespace,
		Resource:  system.UserResource.Name,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	response := mapping.MapFromRecord(result, mapping.UserFromRecord)

	u.cleanPasswords(response)

	return response, nil
}

func (u *userService) Update(ctx context.Context, users []*model.User) ([]*model.User, errors.ServiceError) {
	u.encodePasswords(users)

	for _, user := range users {
		if user.Password == "" {
			record, err := u.recordService.Get(ctx, params.RecordGetParams{
				Namespace: system.UserResource.Namespace,
				Resource:  system.UserResource.Name,
				Id:        user.Id,
			})

			if err != nil {
				return nil, err
			}

			existingUser := mapping.UserFromRecord(record)
			user.Password = existingUser.Password
		}
	}

	// insert records via resource service
	records := mapping.MapToRecord(users, mapping.UserToRecord)

	result, err := u.recordService.Update(ctx, params.RecordUpdateParams{
		Namespace: system.UserResource.Namespace,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	response := mapping.MapFromRecord(result, mapping.UserFromRecord)

	u.cleanPasswords(response)

	return response, nil
}

func (u *userService) Delete(ctx context.Context, ids []string) errors.ServiceError {
	return u.recordService.Delete(ctx, params.RecordDeleteParams{
		Namespace: system.UserResource.Namespace,
		Resource:  system.UserResource.Name,
		Ids:       ids,
	})
}

func (u *userService) Get(ctx context.Context, id string) (*model.User, errors.ServiceError) {
	record, err := u.recordService.Get(ctx, params.RecordGetParams{
		Namespace: system.UserResource.Namespace,
		Resource:  system.UserResource.Name,
		Id:        id,
	})

	if err != nil {
		return nil, err
	}

	response := mapping.UserFromRecord(record)

	u.cleanPasswords([]*model.User{response})

	return response, nil
}

func (u *userService) List(ctx context.Context, query *model.BooleanExpression, limit uint32, offset uint64) ([]*model.User, errors.ServiceError) {
	result, _, err := u.recordService.List(ctx, params.RecordListParams{
		Query:     query,
		Namespace: system.UserResource.Namespace,
		Resource:  system.UserResource.Name,
		Limit:     limit,
		Offset:    offset,
	})

	if err != nil {
		return nil, err
	}

	response := mapping.MapFromRecord(result, mapping.UserFromRecord)

	u.cleanPasswords(response)

	return response, nil
}

func (d *userService) Init(data *model.InitData) {
	d.backendProviderService.MigrateResource(system.UserResource, nil)

	if len(data.InitUsers) > 0 {
		d.encodePasswords(data.InitUsers)
		_, _, err := d.recordService.Create(security.SystemContext, params.RecordCreateParams{
			Namespace:      system.UserResource.Namespace,
			Resource:       system.UserResource.Name,
			Records:        mapping.MapToRecord(data.InitUsers, mapping.UserToRecord),
			IgnoreIfExists: true,
		})

		if err != nil {
			log.Error(err)
		}
	}
}

func (d *userService) encodePasswords(users []*model.User) {
	for _, user := range users {
		if user.Password != "" {
			hashStr, err := security.EncodeKey(user.Password)

			if err != nil {
				panic(err)
			}

			user.Password = hashStr
		}
	}
}

func (u *userService) cleanPasswords(users []*model.User) {
	for _, user := range users {
		user.Password = ""
	}
}

func NewUserService() UserService {
	return &userService{
		serviceName: "UserService",
	}
}
