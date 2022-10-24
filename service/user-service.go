package service

import (
	"context"
	"data-handler/service/mapping"
	"data-handler/service/security"
	"data-handler/service/system"
	"data-handler/stub"
	"data-handler/stub/model"
	log "github.com/sirupsen/logrus"
)

type UserService interface {
	stub.UserServiceServer
	InjectRecordService(service RecordService)
	InjectAuthenticationService(service AuthenticationService)
	InjectResourceService(service ResourceService)
	Init(data *model.InitData)
}

type userService struct {
	stub.UserServiceServer
	recordService         RecordServiceInternal
	authenticationService AuthenticationServiceInternal
	serviceName           string
	resourceService       ResourceServiceInternal
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

func (u *userService) Create(ctx context.Context, request *stub.CreateUserRequest) (*stub.CreateUserResponse, error) {
	u.encodePasswords(request.Users)

	// insert records via resource service
	records := mapping.MapToRecord(request.Users, mapping.UserToRecord)
	systemCtx := security.WithSystemContext(ctx)

	result, err := u.recordService.Create(systemCtx, &stub.CreateRecordRequest{
		Token:   request.Token,
		Records: records,
	})

	if err != nil {
		return nil, err
	}

	return &stub.CreateUserResponse{
		Users: mapping.MapFromRecord(result.Records, mapping.UserFromRecord),
		Error: result.Error,
	}, err
}

func (u *userService) Update(ctx context.Context, request *stub.UpdateUserRequest) (*stub.UpdateUserResponse, error) {
	// insert records via resource service
	records := mapping.MapToRecord(request.Users, mapping.UserToRecord)
	systemCtx := security.WithSystemContext(ctx)
	result, err := u.recordService.Update(systemCtx, &stub.UpdateRecordRequest{
		Token:   request.Token,
		Records: records,
	})

	if err != nil {
		return nil, err
	}

	return &stub.UpdateUserResponse{
		Users: mapping.MapFromRecord(result.Records, mapping.UserFromRecord),
		Error: result.Error,
	}, err
}

func (u *userService) Delete(ctx context.Context, request *stub.DeleteUserRequest) (*stub.DeleteUserResponse, error) {
	systemCtx := security.WithSystemContext(ctx)

	record, err := u.recordService.Delete(systemCtx, &stub.DeleteRecordRequest{
		Token:    request.Token,
		Resource: system.UserResource.Name,
		Ids:      request.Ids,
	})

	if err != nil {
		return nil, err
	}

	return &stub.DeleteUserResponse{
		Error: record.Error,
	}, nil
}

func (u *userService) Get(ctx context.Context, request *stub.GetUserRequest) (*stub.GetUserResponse, error) {
	systemCtx := security.WithSystemContext(ctx)
	record, err := u.recordService.Get(systemCtx, &stub.GetRecordRequest{
		Token:    request.Token,
		Resource: system.UserResource.Name,
		Id:       request.Id,
	})

	if err != nil {
		return nil, err
	}

	return &stub.GetUserResponse{
		User:  mapping.UserFromRecord(record.Record),
		Error: record.Error,
	}, nil
}

func (u *userService) List(ctx context.Context, request *stub.ListUserRequest) (*stub.ListUserResponse, error) {
	systemCtx := security.WithSystemContext(ctx)
	result, err := u.recordService.List(systemCtx, &stub.ListRecordRequest{
		Resource: system.UserResource.Name,
		Token:    request.Token,
	})

	if err != nil {
		return nil, err
	}

	return &stub.ListUserResponse{
		Content: mapping.MapFromRecord(result.Content, mapping.UserFromRecord),
		Error:   result.Error,
	}, err
}

func (d *userService) Init(data *model.InitData) {
	d.resourceService.InitResource(system.UserResource)

	if len(data.InitUsers) > 0 {
		d.encodePasswords(data.InitUsers)
		res, err := d.recordService.Create(security.SystemContext, &stub.CreateRecordRequest{
			Records:        mapping.MapToRecord(data.InitUsers, mapping.UserToRecord),
			IgnoreIfExists: true,
		})

		if err != nil || res.Error != nil {
			log.Error(err, res.Error)
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

func NewUserService() UserService {
	return &userService{
		serviceName: "UserService",
	}
}
