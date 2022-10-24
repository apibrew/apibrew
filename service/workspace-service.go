package service

import (
	"context"
	mapping "data-handler/service/mapping"
	"data-handler/service/security"
	"data-handler/service/system"
	"data-handler/stub"
	"data-handler/stub/model"
	log "github.com/sirupsen/logrus"
)

type WorkspaceService interface {
	stub.WorkspaceServiceServer
	InjectRecordService(service RecordService)
	InjectAuthenticationService(service AuthenticationService)
	Init(data *model.InitData)
	InjectResourceService(service ResourceService)
}

type workspaceService struct {
	stub.WorkspaceServiceServer
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

func (u *workspaceService) Create(ctx context.Context, request *stub.CreateWorkspaceRequest) (*stub.CreateWorkspaceResponse, error) {
	// insert records via resource service
	records := mapping.MapToRecord(request.Workspaces, mapping.WorkspaceToRecord)
	systemCtx := security.WithSystemContext(ctx)

	result, err := u.recordService.Create(systemCtx, &stub.CreateRecordRequest{
		Token:   request.Token,
		Records: records,
	})

	if err != nil {
		return nil, err
	}

	return &stub.CreateWorkspaceResponse{
		Workspaces: mapping.MapFromRecord(result.Records, mapping.WorkspaceFromRecord),
		Error:      result.Error,
	}, err
}

func (u *workspaceService) Update(ctx context.Context, request *stub.UpdateWorkspaceRequest) (*stub.UpdateWorkspaceResponse, error) {
	// insert records via resource service
	records := mapping.MapToRecord(request.Workspaces, mapping.WorkspaceToRecord)
	systemCtx := security.WithSystemContext(ctx)
	result, err := u.recordService.Update(systemCtx, &stub.UpdateRecordRequest{
		Token:   request.Token,
		Records: records,
	})

	if err != nil {
		return nil, err
	}

	return &stub.UpdateWorkspaceResponse{
		Workspaces: mapping.MapFromRecord(result.Records, mapping.WorkspaceFromRecord),
		Error:      result.Error,
	}, err
}

func (u *workspaceService) Delete(ctx context.Context, request *stub.DeleteWorkspaceRequest) (*stub.DeleteWorkspaceResponse, error) {
	systemCtx := security.WithSystemContext(ctx)

	record, err := u.recordService.Delete(systemCtx, &stub.DeleteRecordRequest{
		Token:    request.Token,
		Resource: system.WorkspaceResource.Name,
		Ids:      request.Ids,
	})

	if err != nil {
		return nil, err
	}

	return &stub.DeleteWorkspaceResponse{
		Error: record.Error,
	}, nil
}

func (u *workspaceService) Get(ctx context.Context, request *stub.GetWorkspaceRequest) (*stub.GetWorkspaceResponse, error) {
	systemCtx := security.WithSystemContext(ctx)
	record, err := u.recordService.Get(systemCtx, &stub.GetRecordRequest{
		Token:    request.Token,
		Resource: system.WorkspaceResource.Name,
		Id:       request.Id,
	})

	if err != nil {
		return nil, err
	}

	return &stub.GetWorkspaceResponse{
		Workspace: mapping.WorkspaceFromRecord(record.Record),
		Error:     record.Error,
	}, nil
}

func (u *workspaceService) List(ctx context.Context, request *stub.ListWorkspaceRequest) (*stub.ListWorkspaceResponse, error) {
	systemCtx := security.WithSystemContext(ctx)
	result, err := u.recordService.List(systemCtx, &stub.ListRecordRequest{
		Resource: system.WorkspaceResource.Name,
		Token:    request.Token,
	})

	if err != nil {
		return nil, err
	}

	return &stub.ListWorkspaceResponse{
		Content: mapping.MapFromRecord(result.Content, mapping.WorkspaceFromRecord),
		Error:   result.Error,
	}, err
}

func (d *workspaceService) Init(data *model.InitData) {
	d.resourceService.InitResource(system.WorkspaceResource)

	if len(data.InitWorkspaces) > 0 {
		res, err := d.recordService.Create(security.SystemContext, &stub.CreateRecordRequest{
			Records:        mapping.MapToRecord(data.InitWorkspaces, mapping.WorkspaceToRecord),
			IgnoreIfExists: true,
		})

		if err != nil || res.Error != nil {
			log.Error(err, res.Error)
		}
	}
}

func NewWorkspaceService() WorkspaceService {
	return &workspaceService{
		serviceName: "WorkspaceService",
	}
}
