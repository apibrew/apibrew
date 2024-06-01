package backend_event_handler

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/core"
	"github.com/apibrew/apibrew/pkg/util"
	"strings"
)

func ShortEventInfo(event *core.Event) string {
	var ids []string

	if event == nil {
		return "[removed event]"
	}

	if event.Records == nil {
		return fmt.Sprintf("%s [no records]", event.Id)
	}

	for index, rec := range event.Records {
		if rec != nil {
			ids = append(ids, util.GetRecordId(rec))
		}

		if index > 5 {
			ids = append(ids, "...")
			break
		}
	}

	return fmt.Sprintf("%s [%s]", event.Id, strings.Join(ids, ","))
}
