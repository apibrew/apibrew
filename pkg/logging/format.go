package logging

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/model"
	"strings"
)

func ShortEventInfo(event *model.Event) string {
	var ids = event.Ids

	for _, rec := range event.Records {
		if rec != nil {
			ids = append(ids, rec.Id)
		}
	}

	return fmt.Sprintf("%s [%s]", event.Id, strings.Join(ids, ","))
}
