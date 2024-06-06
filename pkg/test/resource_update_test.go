package test

//func _ignord_TestPrepareResourceMigrationPlan(t *testing.T) {
//	resource1 := &model.Resource{
//		Name:      "test-resource-for-update-1",
//		Namespace: "default",
//		SourceConfig: &model.ResourceSourceConfig{
//			DataSource: setup.DhTest.Name,
//			Entity:     "test-resource-for-update-1",
//		},
//		Properties: []*model.ResourceProperty{
//			{
//				Name:     "prop-1",
//				Type:     model.ResourceProperty_STRING,
//				Length:   128,
//				Required: true,
//			}, {
//				Name:     "prop-2",
//				Type:     model.ResourceProperty_STRING,
//				Length:   128,
//				Required: true,
//				Annotations: map[string]string{
//					annotations.SourceMatchKey: "prop-2",
//				},
//			}, {
//				Name:     "prop-3",
//				Type:     model.ResourceProperty_STRING,
//				Length:   128,
//				Required: true,
//			},
//		},
//	}
//
//	resourceCreateRes, err := resourceClient.Create(setup.Ctx, &stub.CreateResourceRequest{Resources: []*model.Resource{resource1}, DoMigration: true})
//
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	defer func() {
//		if resourceCreateRes != nil {
//			_, _ = resourceClient.Delete(setup.Ctx, &stub.DeleteResourceRequest{
//				Ids:            []string{resourceCreateRes.Resources[0].Id},
//				DoMigration:    true,
//				ForceMigration: true,
//			})
//		}
//	}()
//
//	resource1 = &model.Resource{
//		Name:      "test-resource-for-update-1",
//		Namespace: "default",
//		SourceConfig: &model.ResourceSourceConfig{
//			DataSource: setup.DhTest.Name,
//			Entity:     "test-resource-for-update-1",
//		},
//		Properties: []*model.ResourceProperty{
//			{
//				Name:     "prop-1",
//				Type:     model.ResourceProperty_STRING,
//				Length:   128,
//				Required: true,
//			},
//			{
//				Name:     "prop-2a",
//				Type:     model.ResourceProperty_FLOAT32,
//				Length:   128,
//				Required: false,
//				Annotations: map[string]string{
//					annotations.SourceMatchKey: "prop-2",
//				},
//			},
//			{
//				Name:     "prop-5",
//				Type:     model.ResourceProperty_STRING,
//				Length:   127,
//				Required: false,
//			},
//		},
//	}
//
//	util.NormalizeResource(resource1)
//
//	resource1.Id = resourceCreateRes.Resources[0].Id
//
//	res, err := resourceClient.PrepareResourceMigrationPlan(setup.Ctx, &stub.PrepareResourceMigrationPlanRequest{
//		Resources: []*model.Resource{resource1},
//	})
//
//	if err != nil {
//		assert.Error(t, err)
//	}
//
//	assert.Len(t, res.Plans, 1)
//	assert.Len(t, res.Plans[0].Steps, 4)
//
//	if t.Failed() {
//		return
//	}
//
//	steps := res.Plans[0].Steps
//
//	assert.IsType(t, steps[0].Kind, &model.ResourceMigrationStep_UpdateResource{})
//	assert.IsType(t, steps[1].Kind, &model.ResourceMigrationStep_UpdateProperty{})
//	assert.IsType(t, steps[2].Kind, &model.ResourceMigrationStep_DeleteProperty{})
//	assert.IsType(t, steps[3].Kind, &model.ResourceMigrationStep_CreateProperty{})
//
//	if t.Failed() {
//		return
//	}
//
//	assert.Equal(t, steps[1].Kind.(*model.ResourceMigrationStep_UpdateProperty).UpdateProperty.ChangedFields, []string{"name", "type", "required"})
//}

//func TestResourceUpdateCreateNewPropertyAndMarkAsRequired(t *testing.T) {
//	resource1 := &model.Resource{
//		Name:      "test-resource-for-update-1",
//		Namespace: "default",
//		SourceConfig: &model.ResourceSourceConfig{
//			DataSource: setup.DhTest.Name,
//			Entity:     "test-resource-for-update-1",
//		},
//		Properties: []*model.ResourceProperty{
//			{
//				Name:     "prop-1",
//				Type:     model.ResourceProperty_STRING,
//				Length:   128,
//				Required: true,
//			},
//		},
//	}
//
//	resourceCreateRes, err := resourceClient.Create(setup.Ctx, &stub.CreateResourceRequest{Resources: []*model.Resource{resource1}, DoMigration: true})
//
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	defer func() {
//		if resourceCreateRes != nil {
//			_, _ = resourceClient.Delete(setup.Ctx, &stub.DeleteResourceRequest{
//				Ids:            []string{resourceCreateRes.Resources[0].Id},
//				DoMigration:    true,
//				ForceMigration: true,
//			})
//		}
//	}()
//
//	recordCreateResult1, err := recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{Resource: resource1.Name, Records: []abs.RecordLike{
//		{
//			Properties: map[string]interface{}{
//				"prop-1": structpb.NewStringValue("test-123321"),
//			},
//		},
//	}})
//
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	resource1 = &model.Resource{
//		Id:        resourceCreateRes.Resources[0].Id,
//		Name:      "test-resource-for-update-1",
//		Namespace: "default",
//		SourceConfig: &model.ResourceSourceConfig{
//			DataSource: setup.DhTest.Name,
//			Entity:     "test-resource-for-update-1",
//		},
//		Properties: []*model.ResourceProperty{
//			{
//				Name:     "prop-1",
//				Type:     model.ResourceProperty_STRING,
//				Length:   128,
//				Required: true,
//			},
//			{
//				Name:     "prop-2",
//				Type:     model.ResourceProperty_STRING,
//				Length:   128,
//				Required: false,
//			},
//		},
//	}
//
//	util.NormalizeResource(resource1)
//
//	_, err = resourceClient.Update(setup.Ctx, &stub.UpdateResourceRequest{Resources: []*model.Resource{resource1}, DoMigration: true})
//
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	_, err = recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{Resource: resource1.Name, Records: []abs.RecordLike{
//		{
//			Properties: map[string]interface{}{
//				"prop-1": structpb.NewStringValue("test-123321"),
//				"prop-2": structpb.NewStringValue("test-12332144"),
//			},
//		},
//	}})
//
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	resource1 = &model.Resource{
//		Id:        resourceCreateRes.Resources[0].Id,
//		Name:      "test-resource-for-update-1",
//		Namespace: "default",
//		SourceConfig: &model.ResourceSourceConfig{
//			DataSource: setup.DhTest.Name,
//			Entity:     "test-resource-for-update-1",
//		},
//		Properties: []*model.ResourceProperty{
//			{
//				Name:     "prop-1",
//				Type:     model.ResourceProperty_STRING,
//				Length:   128,
//				Required: true,
//			},
//			{
//				Name:     "prop-2",
//				Type:     model.ResourceProperty_STRING,
//				Length:   128,
//				Required: true,
//			},
//		},
//	}
//
//	util.NormalizeResource(resource1)
//
//	_, err = resourceClient.Update(setup.Ctx, &stub.UpdateResourceRequest{Resources: []*model.Resource{resource1}, DoMigration: true})
//
//	if err == nil {
//		t.Error("marking property prop-2 should be failed, because it is containing null values")
//		return
//	}
//
//	_, err = recordClient.Update(setup.Ctx, &stub.UpdateRecordRequest{Resource: resource1.Name, Records: []abs.RecordLike{
//		{
//			Properties: map[string]interface{}{
//				"id":     recordCreateResult1.Records[0].Properties["id"],
//				"prop-1": structpb.NewStringValue("test-123321"),
//				"prop-2": structpb.NewStringValue("test-12332144"),
//			},
//		},
//	}})
//
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	util.NormalizeResource(resource1)
//
//	_, err = resourceClient.Update(setup.Ctx, &stub.UpdateResourceRequest{Resources: []*model.Resource{resource1}})
//
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	resource1 = &model.Resource{
//		Id:        resourceCreateRes.Resources[0].Id,
//		Name:      "test-resource-for-update-1",
//		Namespace: "default",
//		SourceConfig: &model.ResourceSourceConfig{
//			DataSource: setup.DhTest.Name,
//			Entity:     "test-resource-for-update-1",
//		},
//		Properties: []*model.ResourceProperty{
//			{
//				Name:     "prop-1",
//				Type:     model.ResourceProperty_STRING,
//				Length:   128,
//				Required: true,
//			},
//		},
//	}
//
//	util.NormalizeResource(resource1)
//
//	_, err = resourceClient.Update(setup.Ctx, &stub.UpdateResourceRequest{Resources: []*model.Resource{resource1}, DoMigration: true})
//
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	_, err = recordClient.Create(setup.Ctx, &stub.CreateRecordRequest{Resource: resource1.Name, Records: []abs.RecordLike{
//		{
//			Properties: map[string]interface{}{
//				"prop-1": structpb.NewStringValue("test-123321"),
//				"prop-2": structpb.NewStringValue("test-12332144"),
//			},
//		},
//	}})
//
//	if err == nil {
//		t.Error("prop-2 should complaint about property not exists")
//		return
//	}
//}
