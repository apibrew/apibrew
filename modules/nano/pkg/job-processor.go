package nano

import (
	"context"
	model2 "github.com/apibrew/apibrew/modules/nano/pkg/model"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"time"
)

type jobProcessor struct {
	jobCtxMap    map[string]context.CancelFunc
	m            sync.Mutex
	api          api.Interface
	codeExecutor *codeExecutorService
}

func (f *jobProcessor) MapperTo(record *model.Record) *model2.Job {
	return model2.JobMapperInstance.FromRecord(record)
}

func (f *jobProcessor) Register(ctx context.Context, entity *model2.Job) error {
	f.m.Lock()
	defer f.m.Unlock()

	if f.jobCtxMap == nil {
		f.jobCtxMap = make(map[string]context.CancelFunc)
	}

	var jobCtx, cancelFn = context.WithCancel(context.Background())

	f.jobCtxMap[entity.Name] = cancelFn

	var timeDiff = time.Until(entity.NextExecutionTime)

	if timeDiff < 0 {
		entity.LastExecutionError = util.Pointer("Job is already expired")
		return nil
	}

	log.Printf("Job %s scheduled to be executed on %s in %d seconds \n", entity.Name, entity.NextExecutionTime, int(timeDiff.Seconds()))

	go func() {
		select {
		case <-jobCtx.Done():
			log.Debug("Job context is done before execution time arrived for job:", entity.Name)
			return
		case <-time.After(timeDiff):
		}

		delete(f.jobCtxMap, entity.Name)
		f.execute(util.WithSystemContext(jobCtx), 0, entity.Id.String())
	}()

	return nil
}

func (f *jobProcessor) execute(ctx context.Context, executionNumber int32, jobId string) {
	record, serr := f.api.Load(ctx, unstructured.Unstructured{
		"type": "nano/Job",
		"id":   jobId,
	}, api.LoadParams{})

	if serr != nil {
		log.Error(serr)
		return
	}
	log.Debug("Executing job:", record["name"], executionNumber)

	_, err := f.codeExecutor.RunInlineScript(ctx, record["name"].(string)+"-"+strconv.Itoa(int(executionNumber)), record["source"].(string))

	record = make(unstructured.Unstructured)
	record["id"] = jobId
	record["lastExecutionTime"] = time.Now().Format(time.RFC3339)

	if err != nil {
		log.Error(err)
		record["lastExecutionError"] = err.Error()
	} else {
		record["lastExecutionError"] = ""
	}

	record["type"] = "nano/Job"

	_, serr = f.api.Update(util.SystemContext, record)

	if serr != nil {
		log.Error(serr)
	}
}

func (f *jobProcessor) Update(ctx context.Context, entity *model2.Job) error {
	if err := f.UnRegister(ctx, entity); err != nil {
		return err
	}

	return f.Register(ctx, entity)
}

func (f *jobProcessor) UnRegister(ctx context.Context, entity *model2.Job) error {
	f.m.Lock()
	defer f.m.Unlock()

	if f.jobCtxMap[entity.Name] != nil {
		f.jobCtxMap[entity.Name]()
		delete(f.jobCtxMap, entity.Name)
	}

	return nil
}
