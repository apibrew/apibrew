package main

import (
	"data-handler/app"
	"data-handler/stub/model"
	"data-handler/util"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func main() {
	init := flag.String("init", "", "Initial Data for configuring system")

	flag.Parse()

	initData := &model.InitData{}

	err := util.Read(*init, initData)

	if err != nil {
		log.Fatalf("failed to load init data: %v", err)
	}

	application := new(app.App)

	application.GrpcAddr = fmt.Sprintf("localhost:%d", 9009)
	application.HttpAddr = fmt.Sprintf("localhost:%d", 8008)

	application.SetInitData(initData)

	application.Init()

	application.Serve()
}
