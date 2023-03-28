package main

import (
	"context"
	"github.com/tislib/data-handler/examples/extension/tax_calculator_simple/model"
	"github.com/tislib/data-handler/pkg/client"
	"math"
	"sort"
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
		func(entity *model.Income) (*model.Income, error) {
			rates, err := taxRateRepository.List(context.TODO())

			if err != nil {
				return nil, err
			}

			sort.Slice(rates, func(i, j int) bool {
				return rates[i].Order < rates[j].Order
			})

			var matchedRates []*model.TaxRate

			for _, rate := range rates {
				//if entity.
				if *rate.Country.Name == *entity.Country.Name {
					matchedRates = append(matchedRates, rate)
				}
			}

			var prevRate *model.TaxRate = nil
			entity.Tax = new(int32)
			entity.NetIncome = new(int32)

			for _, rate := range matchedRates {
				if prevRate == nil {
					*entity.Tax += int32(math.Min(float64(rate.Until), float64(entity.GrossIncome)) * float64(rate.Rate))
				} else {
					var taxable int32
					if entity.GrossIncome > rate.Until {
						taxable = rate.Until - prevRate.Until
					} else {
						taxable = entity.GrossIncome - prevRate.Until
					}
					*entity.Tax += int32(float64(taxable) * float64(rate.Rate))
				}
				prevRate = rate
			}

			*entity.NetIncome = entity.GrossIncome - *entity.Tax

			return entity, nil
		}))

	if err := extension.Run(context.TODO()); err != nil {
		panic(err)
	}
}
