package impl

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"time"
)

type statsService struct {
	backendEventHandler backend_event_handler.BackendEventHandler
	redisClient         *redis.Client
	limitations         *model.Limitations
	redisPrefix         string
}

func (s *statsService) prepareRedis(redisConfig *model.RedisConfig) {
	s.redisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       int(redisConfig.Db),
	})

	_, err := s.redisClient.Ping().Result()

	if err != nil {
		log.Error(err)
	}

	s.redisPrefix = redisConfig.Prefix
}

func (s *statsService) Init(config *model.AppConfig) {
	if config.Stats == nil || !config.Stats.Enabled {
		return
	}

	s.prepareRedis(config.Stats.Redis)
	s.backendEventHandler.RegisterHandler(s.prepareHandler())

	if config.Limitations != nil && config.Limitations.Enabled {
		s.limitations = config.Limitations
	}
}

func (a *statsService) prepareHandler() backend_event_handler.Handler {
	return backend_event_handler.Handler{
		Id:       "stats-handler",
		Name:     "stats-handler",
		Fn:       a.handle,
		Order:    1,
		Sync:     true,
		Internal: true,
	}
}

func (s *statsService) handle(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
	if util.IsSystemContext(ctx) {
		return event, nil
	}

	resourceLevelKey := fmt.Sprintf("%s:%s:%s", s.redisPrefix, event.Resource.Namespace, event.Resource.Name)
	namespaceLevelKey := fmt.Sprintf("%s:%s", s.redisPrefix, event.Resource.Namespace)
	globalLevelKey := fmt.Sprintf("%s", s.redisPrefix)

	var listOfKeys = []string{
		resourceLevelKey,
		namespaceLevelKey,
		globalLevelKey,
	}

	var now = time.Now()

	var nowMinute = now.Round(time.Minute).Format("2006-01-02T15:04")

	// checking limitations

	if s.limitations != nil && s.limitations.Enabled {
		if event.Action == model.Event_CREATE {
			var globalCount = handleError(s.redisClient.Get(globalLevelKey + ":count").Int64())
			if globalCount > 0 && globalCount >= int64(s.limitations.MaxRecordCount) {
				return nil, errors.RateLimitError.WithDetails(fmt.Sprintf("MaxRecordCount exceeded: %d", s.limitations.MaxRecordCount))
			}

			if int64(s.limitations.MaxNamespaceCount) > 0 && event.Resource.Namespace == resources.NamespaceResource.Namespace && event.Resource.Name == resources.NamespaceResource.Name {
				var namespaceCount = handleError(s.redisClient.Get(resourceLevelKey + ":count").Int64())
				if namespaceCount > 0 && namespaceCount >= int64(s.limitations.MaxNamespaceCount) {
					return nil, errors.RateLimitError.WithDetails(fmt.Sprintf("MaxNamespaceCount exceeded: %d", s.limitations.MaxNamespaceCount))
				}
			} else if int64(s.limitations.MaxResourceCount) > 0 && event.Resource.Namespace == resources.ResourceResource.Namespace && event.Resource.Name == resources.ResourceResource.Name {
				var resourceCount = handleError(s.redisClient.Get(resourceLevelKey + ":count").Int64())
				if resourceCount > 0 && resourceCount >= int64(s.limitations.MaxResourceCount) {
					return nil, errors.RateLimitError.WithDetails(fmt.Sprintf("MaxResourceCount exceeded: %d", s.limitations.MaxResourceCount))
				}
			}
		}

		if s.limitations.RequestPerMinute > 0 {
			var globalMinuteRequestCount = handleError(s.redisClient.Get(globalLevelKey + ":rate:" + nowMinute).Int64())
			if globalMinuteRequestCount > 0 && globalMinuteRequestCount > int64(s.limitations.RequestPerMinute) {
				return nil, errors.RateLimitError.WithDetails(fmt.Sprintf("RequestPerMinute exceeded: %d", s.limitations.RequestPerMinute))
			}
		}
	}

	for _, key := range listOfKeys {
		s.redisClient.Incr(key + ":total")

		s.redisClient.Incr(key + ":rate:" + nowMinute)
		s.redisClient.Expire(key+":rate:"+nowMinute, time.Minute)
	}

	if event.Action == model.Event_CREATE {
		s.redisClient.IncrBy(resourceLevelKey+":count", int64(len(event.Records)))
		s.redisClient.IncrBy(namespaceLevelKey+":count", int64(len(event.Records)))
		s.redisClient.IncrBy(globalLevelKey+":count", int64(len(event.Records)))
	} else if event.Action == model.Event_DELETE {
		s.redisClient.DecrBy(resourceLevelKey+":count", int64(len(event.Records)))
		s.redisClient.DecrBy(namespaceLevelKey+":count", int64(len(event.Records)))
		s.redisClient.DecrBy(globalLevelKey+":count", int64(len(event.Records)))
	}

	return event, nil
}

func handleError(res int64, err error) int64 {
	if err != nil {
		return 0
	}

	return res
}

func NewStatsService(backendEventHandler backend_event_handler.BackendEventHandler) service.StatsService {
	return &statsService{
		backendEventHandler: backendEventHandler,
	}
}
