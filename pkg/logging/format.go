package logging

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"strings"
)

func ShortEventInfo(event *model.Event) string {
	var ids []string

	if event.Records == nil {
		return fmt.Sprintf("%s [no records]", event.Id)
	}

	for _, rec := range event.Records {
		if rec != nil {
			ids = append(ids, util.GetRecordId(rec))
		}
	}

	return fmt.Sprintf("%s [%s]", event.Id, strings.Join(ids, ","))
}
