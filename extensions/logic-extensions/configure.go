package main

import (
	"context"
	"github.com/apibrew/apibrew/pkg/model"
	log "github.com/sirupsen/logrus"
)

func reconfigureResourceExtensions() {
	var err error

	err = dhClient.ApplyExtension(context.TODO(), &model.Extension{
		Name: config.ExtensionName + "-before",
		Selector: &model.EventSelector{
			Namespaces: []string{"extensions"},
			Resources:  []string{"ForceCheck"},
		},
		Order:     90,
		Finalizes: false,
		Sync:      true,
		Responds:  true,
		Call: &model.ExternalCall{
			FunctionCall: &model.FunctionCall{
				Host:         config.RemoteEngineAddr,
				FunctionName: "before",
			},
		},
		AuditData:   nil,
		Version:     0,
		Annotations: nil,
	})

	if err != nil {
		log.Fatal(err)
	}

	err = dhClient.ApplyExtension(context.TODO(), &model.Extension{
		Name: config.ExtensionName + "-after",
		Selector: &model.EventSelector{
			Namespaces: []string{"extensions"},
			Resources:  []string{"Function", "FunctionExecutionEngine", "FunctionExecution", "FunctionTrigger", "ResourceRule", "Schedule"},
		},
		Order:     110,
		Finalizes: true,
		Sync:      true,
		Responds:  true,
		Call: &model.ExternalCall{
			FunctionCall: &model.FunctionCall{
				Host:         config.RemoteEngineAddr,
				FunctionName: "after",
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
