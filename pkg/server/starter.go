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
	flag.Parse()
	//grayLogAddr := flag.String("gray-log-addr", "", "Initial Data for configuring system")

	flag.Parse()

	var err error
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

	if appConfig.LoggingConfig != nil {
		log.SetLevel(log.Level(appConfig.LoggingConfig.Level))
		log.SetReportCaller(appConfig.LoggingConfig.ReportCaller)

		switch appConfig.LoggingConfig.Format {
		case model.LogFormat_JSON:
			log.SetFormatter(&log.JSONFormatter{})
		case model.LogFormat_TEXT:
			log.SetFormatter(&log.TextFormatter{})
		}
	} else {
		log.SetLevel(log.InfoLevel)
		log.SetReportCaller(false)
		log.SetFormatter(&log.TextFormatter{})
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

	unMatchedLis := tcpm.Match(cmux.Any())

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

	go handleUnmatchedListener(unMatchedLis)

	if err = tcpm.Serve(); err != nil {
		grpcServer.Stop()
		log.Fatal(err)
	}
}

func handleUnmatchedListener(lis net.Listener) {
	srv := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Warnf("unmatched request: %s", r.URL.String())
			log.Warn(r.Header)
			w.WriteHeader(http.StatusNotFound)
		}),
	}

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
