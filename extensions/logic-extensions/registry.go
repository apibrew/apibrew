package main

import (
	"context"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/stub"
	log "github.com/sirupsen/logrus"
)

var functions []*model.Record
var functionTriggers []*model.Record
var resourceRules []*model.Record
var schedules []*model.Record

func reloadFunctions() {
	log.Println("Reloading functions")
	list, err := dhClient.GetRecordClient().List(context.TODO(), &stub.ListRecordRequest{
		Namespace: "extensions",
		Resource:  "Function",
	})

	if err != nil {
		log.Fatal(err)
	}

	functions = list.Content
	log.Println("Reloaded functions")
}

func reloadFunctionTriggers() {
	log.Println("Reloading function triggers")
	list, err := dhClient.GetRecordClient().List(context.TODO(), &stub.ListRecordRequest{
		Namespace: "extensions",
		Resource:  "FunctionTrigger",
	})

	if err != nil {
		log.Fatal(err)
	}

	functionTriggers = list.Content
	log.Println("Reloaded function triggers")
	reconfigureResourceExtensions()
}

func reloadResourceRules() {
	log.Println("Reloading resource rules")
	list, err := dhClient.GetRecordClient().List(context.TODO(), &stub.ListRecordRequest{
		Namespace: "extensions",
		Resource:  "ResourceRule",
	})

	if err != nil {
		log.Fatal(err)
	}

	resourceRules = list.Content
	log.Println("Reloaded resource rules")
	reconfigureResourceExtensions()
}

func reloadSchedules() {
	log.Println("Reloading schedules")
	list, err := dhClient.GetRecordClient().List(context.TODO(), &stub.ListRecordRequest{
		Namespace: "extensions",
		Resource:  "Schedule",
	})

	if err != nil {
		log.Fatal(err)
	}

	schedules = list.Content
	log.Println("Reloaded schedules")
}

func reloadAll() {
	reloadFunctions()
	reloadFunctionTriggers()
	reloadResourceRules()
	reloadSchedules()
}
