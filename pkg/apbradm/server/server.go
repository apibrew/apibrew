package server

import (
	"context"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	stub2 "github.com/tislib/apibrew/pkg/apbradm/stub"
	"github.com/tislib/apibrew/pkg/model"
	grpc2 "github.com/tislib/apibrew/pkg/server/grpc"
	"github.com/tislib/apibrew/pkg/service"
	"github.com/tislib/apibrew/pkg/stub"
	"github.com/tislib/apibrew/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"strings"
)

type Server struct {
	Init     string
	LogLevel string
}

func (s Server) Run() error {
	logLevel, err := log.ParseLevel(s.LogLevel)

	if err != nil {
		log.Fatal(err)
	}

	log.SetReportCaller(logLevel == log.TraceLevel)
	log.SetLevel(logLevel)

	flag.Parse()

	initData := &model.InitData{}

	if strings.HasSuffix(s.Init, "pb") {
		err = util.Read(s.Init, initData)
	} else if strings.HasSuffix(s.Init, "json") {
		err = util.ReadJson(s.Init, initData)
	} else {
		log.Fatal("init config is not set")
	}

	if err != nil {
		log.Fatalf("failed to load init data: %v", err)
	}

	app := new(service.App)

	app.SetInitData(initData)

	//if grayLogAddr != nil {
	//	app.SetGrayLogAddr(*grayLogAddr)
	//}

	app.Init()

	var opts = []grpc.ServerOption{}

	grpcServer := grpc.NewServer(opts...)

	reflection.Register(grpcServer)

	stub2.RegisterNodeServer(grpcServer, &nodeService{container: app})
	stub.RegisterAuthenticationServer(grpcServer, grpc2.NewAuthenticationServer(app.GetAuthenticationService()))

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", initData.Config.Host, initData.Config.Port))
	if err != nil {
		log.Fatal(err)
	}

	s.PostRun(app)

	return grpcServer.Serve(l)
}

func (s Server) PostRun(app *service.App) {
	if err := app.GetResourceService().Apply(context.Background(), NodeResource, true, true); err != nil {
		log.Fatalf("failed to apply node resource: %v", err)
	}
}
