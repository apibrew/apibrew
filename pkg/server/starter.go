package server

import (
	"flag"
	"fmt"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/server/grpc"
	"github.com/apibrew/apibrew/pkg/server/rest"
	"github.com/apibrew/apibrew/pkg/service/impl"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"net"
	"net/http"
	"strings"
)

import _ "net/http/pprof"

func Run() {
	init := flag.String("config", "", "Config file")
	logLevelStr := flag.String("log-level", "info", "Debug flag")
	flag.Parse()
	//grayLogAddr := flag.String("gray-log-addr", "", "Initial Data for configuring system")

	logLevel, err := log.ParseLevel(*logLevelStr)

	if err != nil {
		log.Fatal(err)
	}

	log.SetReportCaller(logLevel == log.TraceLevel)
	log.SetLevel(logLevel)

	flag.Parse()

	appConfig := &model.AppConfig{}

	if strings.HasSuffix(*init, "pb") {
		err = util.Read(*init, appConfig)
	} else if strings.HasSuffix(*init, "json") {
		err = util.ReadJson(*init, appConfig)
	} else {
		log.Fatal("config is not set")
	}

	if err != nil {
		log.Fatalf("failed to load init data: %v", err)
	}

	app := new(impl.App)

	app.SetConfig(appConfig)

	//if grayLogAddr != nil {
	//	app.SetGrayLogAddr(*grayLogAddr)
	//}

	app.Init()

	// Create the main listener.
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port))
	if err != nil {
		log.Fatal(err)
	}
	tcpm := cmux.New(l)

	// Declare the match for different services required.
	httpl := tcpm.Match(cmux.HTTP1Fast("PATCH"))
	grpcl := tcpm.MatchWithWriters(
		cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"),
		cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc+protohelper"),
	)
	http2 := tcpm.Match(cmux.HTTP2())
	http2Tls := tcpm.Match(cmux.TLS())

	grpcServer := grpc.NewGrpcServer(app)
	grpcServer.Init()
	restServer := rest.NewServer(app)
	restServer.Init()

	go grpcServer.Serve(grpcl)
	go restServer.ServeHttp(httpl)
	go restServer.ServeH2C(http2)
	go restServer.ServeHttp2Tls(http2Tls)

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	if err = tcpm.Serve(); err != nil {
		grpcServer.Stop()
		log.Fatal(err)
	}
}
