package logging

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"strings"
)

func ShortEventInfo(event *model.Event) string {
	var ids []string

	if event == nil {
		return "[removed event]"
	}

	if event.Records == nil {
		return fmt.Sprintf("%s [no records]", event.Id)
	}

	for index, rec := range event.Records {
		if rec != nil {
			uns, err := unstructured.FromRecord(rec)

			if err != nil {
				panic(err)
			}

			ids = append(ids, util.GetRecordId(uns))
		}

		if index > 5 {
			ids = append(ids, "...")
			break
		}
	}

	return fmt.Sprintf("%s [%s]", event.Id, strings.Join(ids, ","))
}
