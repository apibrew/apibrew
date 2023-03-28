package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/examples/extension/tax_calculator_simple/model"
	"github.com/tislib/data-handler/pkg/client"
)

func main() {
	dhClient, err := client.NewDhClient(client.DhClientParams{
		Addr:     "127.0.0.1:9009",
		Insecure: true,
	})

	if err != nil {
		panic(err)
	}

	err = dhClient.AuthenticateWithUsernameAndPassword("admin", "admin")

	taxRateRepository := model.NewTaxRateRepository(dhClient)

	if err != nil {
		panic(err)
	}

	extension := dhClient.NewExtension("127.0.0.1:37612")

	extension.RegisterFunction("income_calculator_calculate", client.CreateRecordTypedFunction[*model.Income](&model.Income{},
		func(entity *model.Income) *model.Income {
			rates, err := taxRateRepository.List(context.TODO())

			if err != nil {
				panic(err)
			}

			log.Print(rates)

			entity.Tax = new(int32)

			*entity.Tax = int32(0.05 * float64(entity.GrossIncome))

			return entity
		}))

	if err := extension.Run(context.TODO()); err != nil {
		panic(err)
	}
}
