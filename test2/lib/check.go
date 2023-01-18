package lib

import "data-handler/model"

type HasError interface {
	GetError() *model.Error
}

func checkResp(resp HasError, err error) {
	if err != nil {
		panic(err)
	}

	if resp.GetError() != nil {
		panic(resp.GetError())
	}
}
