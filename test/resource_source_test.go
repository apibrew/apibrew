package test

import (
	"data-handler/model"
	"testing"
)

func TestCreateResource(t *testing.T) {
	ctx := prepareTextContext()

	withDataSource(ctx, t, container, dataSource1, func(createdDataSource *model.DataSource) {
		// testing is done

	})
}
