package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	mapping2 "github.com/tislib/data-handler/pkg/resources/mapping"
	"github.com/tislib/data-handler/pkg/service/security"
)

type userService struct {
	recordService          abs.RecordService
	authenticationService  abs.AuthenticationService
	serviceName            string
	resourceService        abs.ResourceService
	backendProviderService abs.BackendProviderService
}

func (u *userService) InjectBackendProviderService(backendProviderService abs.BackendProviderService) {
	u.backendProviderService = backendProviderService
}
func (u *userService) Create(ctx context.Context, users []*model.User) ([]*model.User, errors.ServiceError) {
	// insert records via resource service
	records := mapping2.MapToRecord(users, mapping2.UserToRecord)

	result, _, err := u.recordService.Create(ctx, abs.RecordCreateParams{
		Namespace: resources.UserResource.Namespace,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	response := mapping2.MapFromRecord(result, mapping2.UserFromRecord)

	return response, nil
}

func (u *userService) Update(ctx context.Context, users []*model.User) ([]*model.User, errors.ServiceError) {
	// update records via resource service
	records := mapping2.MapToRecord(users, mapping2.UserToRecord)

	result, err := u.recordService.Update(ctx, abs.RecordUpdateParams{
		Namespace: resources.UserResource.Namespace,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	response := mapping2.MapFromRecord(result, mapping2.UserFromRecord)

	return response, nil
}

func (u *userService) Delete(ctx context.Context, ids []string) errors.ServiceError {
	return u.recordService.Delete(ctx, abs.RecordDeleteParams{
		Namespace: resources.UserResource.Namespace,
		Resource:  resources.UserResource.Name,
		Ids:       ids,
	})
}

func (u *userService) Get(ctx context.Context, id string) (*model.User, errors.ServiceError) {
	record, err := u.recordService.Get(ctx, abs.RecordGetParams{
		Namespace: resources.UserResource.Namespace,
		Resource:  resources.UserResource.Name,
		Id:        id,
	})

	if err != nil {
		return nil, err
	}

	response := mapping2.UserFromRecord(record)

	return response, nil
}

func (u *userService) List(ctx context.Context, query *model.BooleanExpression, limit uint32, offset uint64) ([]*model.User, errors.ServiceError) {
	result, _, err := u.recordService.List(ctx, abs.RecordListParams{
		Query:     query,
		Namespace: resources.UserResource.Namespace,
		Resource:  resources.UserResource.Name,
		Limit:     limit,
		Offset:    offset,
	})

	if err != nil {
		return nil, err
	}

	response := mapping2.MapFromRecord(result, mapping2.UserFromRecord)

	return response, nil
}

func (d *userService) Init(data *model.InitData) {
	if len(data.InitUsers) > 0 {
		for _, user := range data.InitUsers {
			hashStr, err := security.EncodeKey(user.Password)

			if err != nil {
				panic(err)
			}

			user.Password = hashStr
		}
		_, _, err := d.recordService.Create(security.SystemContext, abs.RecordCreateParams{
			Namespace:      resources.UserResource.Namespace,
			Records:        mapping2.MapToRecord(data.InitUsers, mapping2.UserToRecord),
			IgnoreIfExists: true,
		})

		if err != nil {
			log.Error(err)
		}
	}
}

func NewUserService(resourceService abs.ResourceService, recordService abs.RecordService, backendProviderService abs.BackendProviderService) abs.UserService {
	return &userService{
		serviceName:            "UserService",
		recordService:          recordService,
		resourceService:        resourceService,
		backendProviderService: backendProviderService,
	}
}
