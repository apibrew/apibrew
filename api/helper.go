package api

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"net/http"
)

var mo = protojson.MarshalOptions{
	Multiline:       true,
	EmitUnpopulated: true,
}

var umo = protojson.UnmarshalOptions{}

func respondMessage(writer http.ResponseWriter) func(proto.Message, error) {
	return func(msg proto.Message, err error) {
		if err != nil {
			log.Error(err)
		}

		body, err := mo.Marshal(msg)

		if err != nil {
			log.Error(err)
		}

		writer.Header().Set("Content-Type", "application/json")

		writer.Write(body)
	}
}

func getToken(request *http.Request) string {
	return request.Header.Get("Authorization")
}
