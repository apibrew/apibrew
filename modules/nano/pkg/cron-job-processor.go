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
	"sync/atomic"
	"time"
)
import "github.com/robfig/cron/v3"

type cronJobProcessor struct {
	cronMap      map[string]*cron.Cron
	m            sync.Mutex
	api          api.Interface
	codeExecutor *codeExecutorService
}

func (f *cronJobProcessor) MapperTo(record *model.Record) *model2.CronJob {
	return model2.CronJobMapperInstance.FromRecord(record)
}

func (f *cronJobProcessor) Register(ctx context.Context, entity *model2.CronJob) error {
	f.m.Lock()
	defer f.m.Unlock()

	if f.cronMap == nil {
		f.cronMap = make(map[string]*cron.Cron)
	}

	var c = cron.New(cron.WithSeconds())
	f.cronMap[entity.Name] = c

	var executionNumber int32 = 0

	_, err := c.AddFunc(entity.Expression, func() {
		atomic.AddInt32(&executionNumber, 1)

		f.execute(util.WithSystemContext(context.Background()), executionNumber, entity.Id.String())
	})

	if err != nil {
		return err
	}

	c.Start()

	return nil
}

func (f *cronJobProcessor) execute(ctx context.Context, executionNumber int32, cronId string) {
	record, serr := f.api.Load(ctx, unstructured.Unstructured{
		"type": "nano/CronJob",
		"id":   cronId,
	}, api.LoadParams{})

	if serr != nil {
		log.Error(serr)
		return
	}
	log.Debug("Executing CronJob:", record["name"], executionNumber)

	_, err := f.codeExecutor.RunInlineScript(ctx, record["name"].(string)+"-"+strconv.Itoa(int(executionNumber)), record["source"].(string))

	record = make(unstructured.Unstructured)
	record["id"] = cronId
	record["lastExecutionTime"] = time.Now().Format(time.RFC3339)

	if err != nil {
		log.Error(err)
		record["lastExecutionError"] = err.Error()
	} else {
		record["lastExecutionError"] = ""
	}

	record["type"] = "nano/CronJob"

	_, serr = f.api.Update(util.SystemContext, record)

	if serr != nil {
		log.Error(serr)
	}
}

func (f *cronJobProcessor) Update(ctx context.Context, entity *model2.CronJob) error {
	if entity.Expression == "" && entity.Source == "" {
		return nil
	}

	if err := f.UnRegister(ctx, entity); err != nil {
		return err
	}

	return f.Register(ctx, entity)
}

func (f *cronJobProcessor) UnRegister(ctx context.Context, entity *model2.CronJob) error {
	f.m.Lock()
	defer f.m.Unlock()

	c, ok := f.cronMap[entity.Name]
	if !ok {
		log.Warnf("cron job not found")
		return nil
	}

	c.Stop()

	delete(f.cronMap, entity.Name)

	return nil
}
