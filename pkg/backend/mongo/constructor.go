package mongo

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/resource_model"
	_ "github.com/lib/pq"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoResourceServiceBackend(dataSource *resource_model.DataSource) abs.Backend {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dataSource.Options["uri"]))
	if err != nil {
		panic(err)
	}

	bck := &mongoBackend{
		dataSource: dataSource,
		client:     client,
		dbName:     dataSource.Options["db_name"],
	}

	return bck
}
