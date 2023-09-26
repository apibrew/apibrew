package logging

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/model"
	"strings"
)

func ShortEventInfo(event *model.Event) string {
	var ids []string

	for _, id := range event.Ids {
		ids = append(ids, id)
	}

	for _, rec := range event.Records {
		if rec != nil {
			ids = append(ids, rec.Id)
		}
	}

	return fmt.Sprintf("[%s]%s/%s/%s - [%s]", event.Action, event.Resource.Namespace, event.Resource.Name, event.Id, strings.Join(ids, ","))
}
