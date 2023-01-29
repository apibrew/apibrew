package server

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"github.com/tislib/data-handler/pkg/app"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/server/grpc"
	"github.com/tislib/data-handler/pkg/server/rest"
	"github.com/tislib/data-handler/pkg/util"
	"net"
	"strings"
)

func Run() {
	log.SetLevel(log.TraceLevel)
	log.SetReportCaller(true)
	init := flag.String("init", "", "Initial Data for configuring system")
	grayLogAddr := flag.String("gray-log-addr", "", "Initial Data for configuring system")

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

	app := new(app.App)

	app.SetInitData(initData)

	if grayLogAddr != nil {
		app.SetGrayLogAddr(*grayLogAddr)
	}

	app.Init()

	// Create the main listener.
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", initData.Config.Host, initData.Config.Port))
	if err != nil {
		log.Fatal(err)
	}
	tcpm := cmux.New(l)

	// Declare the match for different services required.
	httpl := tcpm.Match(cmux.HTTP1Fast())
	grpcl := tcpm.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	http2 := tcpm.Match(cmux.HTTP2())

	grpcServer := grpc.NewGrpcServer(app)
	grpcServer.Init(initData)
	restServer := rest.NewServer(app)

	go grpcServer.Serve(grpcl)
	go restServer.Serve(httpl)
	go restServer.Serve(http2)

	if err = tcpm.Serve(); err != nil {
		grpcServer.Stop()
		log.Fatal(err)
	}
}
