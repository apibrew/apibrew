package flags

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

func check(err error) {
	if err != nil {
		st, isStatus := status.FromError(err)

		if isStatus {
			log.Fatalf(st.Message())
		} else {
			log.Fatal(err)
		}
	}
}
