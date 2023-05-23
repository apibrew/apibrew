package main

import (
	"context"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/apibrew/apibrew/pkg/model"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := NewConfig()

	dhClient, err := client.NewDhClient(client.DhClientParams{
		Addr:     config.ApbrAddr,
		Insecure: true,
		Token:    config.Token,
	})

	if err != nil {
		log.Fatal(err)
	}

	ext := dhClient.NewExtension(config.EngineAddr)

	ext.RegisterFunction("before", func(ctx context.Context, entity *model.Event) (*model.Event, error) {
		return entity, nil
	})

	ext.RegisterFunction("after", func(ctx context.Context, entity *model.Event) (*model.Event, error) {
		return entity, nil
	})

	if err != nil {
		log.Fatal(err)
	}

	applyExtensions(dhClient, config)

	log.Print("Started")
	err = ext.Run(context.TODO())
}

func applyExtensions(dhClient client.DhClient, config *Config) {
	var err error

	err = dhClient.ApplyExtension(context.TODO(), &model.Extension{
		Name: config.ExtensionName + "-before",
		Selector: &model.EventSelector{
			Namespaces: []string{"extensions"},
			Resources:  []string{"Function", "FunctionExecutionEngine", "FunctionExecution", "FunctionTrigger", "ResourceRule", "Schedule"},
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
