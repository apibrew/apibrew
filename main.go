package main

import (
	"data-handler/app"
	"data-handler/model"
	"data-handler/util"
	"flag"
	log "github.com/sirupsen/logrus"
	"strings"
)

func main() {
	init := flag.String("init", "", "Initial Data for configuring system")

	flag.Parse()

	initData := &model.InitData{}

	var err error
	if strings.HasSuffix(*init, "pb") {
		err = util.Read(*init, initData)
	} else if strings.HasSuffix(*init, "json") {
		err = util.ReadJson(*init, initData)
	} else {
		log.Fatal("init config is not set")
	}

	if err != nil {
		log.Fatalf("failed to load init data: %v", err)
	}

	application := new(app.App)

	application.SetInitData(initData)

	application.Init()

	application.Serve()
}
