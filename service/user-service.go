package service

import (
	"context"
	"data-handler/model"
	"data-handler/service/errors"
	"data-handler/service/mapping"
	"data-handler/service/params"
	"data-handler/service/security"
	"data-handler/service/system"
	log "github.com/sirupsen/logrus"
)

type UserService interface {
	InjectRecordService(service RecordService)
	InjectAuthenticationService(service AuthenticationService)
	InjectResourceService(service ResourceService)
	Init(data *model.InitData)
	Create(ctx context.Context, users []*model.User) ([]*model.User, errors.ServiceError)
	Update(ctx context.Context, users []*model.User) ([]*model.User, errors.ServiceError)
	Delete(ctx context.Context, ids []string) errors.ServiceError
	Get(ctx context.Context, id string) (*model.User, errors.ServiceError)
	List(ctx context.Context, query *model.BooleanExpression, limit uint32, offset uint64) ([]*model.User, errors.ServiceError)
}

type userService struct {
	recordService         RecordService
	authenticationService AuthenticationService
	serviceName           string
	resourceService       ResourceService
}

func (u *userService) InjectResourceService(service ResourceService) {
	u.resourceService = service
}

func (u *userService) InjectAuthenticationService(service AuthenticationService) {
	u.authenticationService = service
}

func (u *userService) InjectRecordService(service RecordService) {
	u.recordService = service
}

func (u *userService) Create(ctx context.Context, users []*model.User) ([]*model.User, errors.ServiceError) {
	u.encodePasswords(users)

	// insert records via resource service
	records := mapping.MapToRecord(users, mapping.UserToRecord)
	systemCtx := security.WithSystemContext(ctx)

	result, _, err := u.recordService.Create(systemCtx, params.RecordCreateParams{
		Workspace: system.UserResource.Workspace,
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

	systemCtx := security.WithSystemContext(ctx)

	for _, user := range users {
		if user.Password == "" {
			record, err := u.recordService.Get(systemCtx, params.RecordGetParams{
				Workspace: system.UserResource.Workspace,
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

	result, err := u.recordService.Update(systemCtx, params.RecordUpdateParams{
		Workspace: system.UserResource.Workspace,
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
	systemCtx := security.WithSystemContext(ctx)

	return u.recordService.Delete(systemCtx, params.RecordDeleteParams{
		Workspace: system.UserResource.Workspace,
		Resource:  system.UserResource.Name,
		Ids:       ids,
	})
}

func (u *userService) Get(ctx context.Context, id string) (*model.User, errors.ServiceError) {
	systemCtx := security.WithSystemContext(ctx)
	record, err := u.recordService.Get(systemCtx, params.RecordGetParams{
		Workspace: system.UserResource.Workspace,
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
	systemCtx := security.WithSystemContext(ctx)
	result, _, err := u.recordService.List(systemCtx, params.RecordListParams{
		Query:     query,
		Workspace: system.UserResource.Workspace,
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
	d.resourceService.InitResource(system.UserResource)

	if len(data.InitUsers) > 0 {
		d.encodePasswords(data.InitUsers)
		_, _, err := d.recordService.Create(security.SystemContext, params.RecordCreateParams{
			Workspace:      system.UserResource.Workspace,
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
