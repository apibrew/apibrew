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

type WorkspaceService interface {
	InjectRecordService(service RecordService)
	InjectAuthenticationService(service AuthenticationService)
	InjectResourceService(service ResourceService)
	Init(data *model.InitData)
	Create(ctx context.Context, workspaces []*model.Workspace) ([]*model.Workspace, errors.ServiceError)
	Update(ctx context.Context, workspaces []*model.Workspace) ([]*model.Workspace, errors.ServiceError)
	Delete(ctx context.Context, ids []string) errors.ServiceError
	Get(ctx context.Context, id string) (*model.Workspace, errors.ServiceError)
	List(ctx context.Context) ([]*model.Workspace, errors.ServiceError)
}

type workspaceService struct {
	recordService         RecordService
	authenticationService AuthenticationService
	serviceName           string
	resourceService       ResourceService
}

func (u *workspaceService) InjectResourceService(service ResourceService) {
	u.resourceService = service
}

func (u *workspaceService) InjectAuthenticationService(service AuthenticationService) {
	u.authenticationService = service
}

func (u *workspaceService) InjectRecordService(service RecordService) {
	u.recordService = service
}

func (u *workspaceService) Create(ctx context.Context, workspaces []*model.Workspace) ([]*model.Workspace, errors.ServiceError) {
	// insert records via resource service
	records := mapping.MapToRecord(workspaces, mapping.WorkspaceToRecord)
	systemCtx := security.WithSystemContext(ctx)

	result, _, err := u.recordService.Create(systemCtx, params.RecordCreateParams{
		Workspace: system.WorkspaceResource.Workspace,
		Resource:  system.WorkspaceResource.Name,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	return mapping.MapFromRecord(result, mapping.WorkspaceFromRecord), nil
}

func (u *workspaceService) Update(ctx context.Context, workspaces []*model.Workspace) ([]*model.Workspace, errors.ServiceError) {
	// insert records via resource service
	records := mapping.MapToRecord(workspaces, mapping.WorkspaceToRecord)
	systemCtx := security.WithSystemContext(ctx)
	result, err := u.recordService.Update(systemCtx, params.RecordUpdateParams{
		Workspace: system.WorkspaceResource.Workspace,
		Records:   records,
	})

	if err != nil {
		return nil, err
	}

	return mapping.MapFromRecord(result, mapping.WorkspaceFromRecord), nil
}

func (u *workspaceService) Delete(ctx context.Context, ids []string) errors.ServiceError {
	systemCtx := security.WithSystemContext(ctx)

	return u.recordService.Delete(systemCtx, params.RecordDeleteParams{
		Workspace: system.WorkspaceResource.Workspace,
		Resource:  system.WorkspaceResource.Name,
		Ids:       ids,
	})
}

func (u *workspaceService) Get(ctx context.Context, id string) (*model.Workspace, errors.ServiceError) {
	systemCtx := security.WithSystemContext(ctx)
	record, err := u.recordService.Get(systemCtx, params.RecordGetParams{
		Workspace: system.WorkspaceResource.Workspace,
		Resource:  system.WorkspaceResource.Name,
		Id:        id,
	})

	if err != nil {
		return nil, err
	}

	return mapping.WorkspaceFromRecord(record), nil
}

func (u *workspaceService) List(ctx context.Context) ([]*model.Workspace, errors.ServiceError) {
	systemCtx := security.WithSystemContext(ctx)
	result, _, err := u.recordService.List(systemCtx, params.RecordListParams{
		Workspace: system.WorkspaceResource.Workspace,
		Resource:  system.WorkspaceResource.Name,
	})

	if err != nil {
		return nil, err
	}

	return mapping.MapFromRecord(result, mapping.WorkspaceFromRecord), err
}

func (d *workspaceService) Init(data *model.InitData) {
	d.resourceService.InitResource(system.WorkspaceResource)

	if len(data.InitWorkspaces) > 0 {
		_, _, err := d.recordService.Create(security.SystemContext, params.RecordCreateParams{
			Workspace:      system.WorkspaceResource.Workspace,
			Resource:       system.WorkspaceResource.Name,
			Records:        mapping.MapToRecord(data.InitWorkspaces, mapping.WorkspaceToRecord),
			IgnoreIfExists: true,
		})

		if err != nil {
			log.Error(err)
		}
	}
}

func NewWorkspaceService() WorkspaceService {
	return &workspaceService{
		serviceName: "WorkspaceService",
	}
}
