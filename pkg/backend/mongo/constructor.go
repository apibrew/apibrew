package mongo

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	_ "github.com/lib/pq"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoResourceServiceBackend(dataSource *model.DataSource) abs.Backend {
	mongoOptions := dataSource.Params.(*model.DataSource_MongoParams)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoOptions.MongoParams.Uri))
	if err != nil {
		panic(err)
	}

	bck := &mongoBackend{
		dataSource: dataSource,
		client:     client,
		dbName:     mongoOptions.MongoParams.DbName,
	}

	return bck
}
