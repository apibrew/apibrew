package mongo

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoResourceServiceBackend(dataSource *model.DataSource) abs.Backend {
	mongoOptions := dataSource.Params.(*model.DataSource_MongoParams)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoOptions.MongoParams.Uri))
	if err != nil {
		panic(err)
	}

	bck := mongoBackend{
		dataSource: dataSource,
		client:     client,
	}

	return bck
}
