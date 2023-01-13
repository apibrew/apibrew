package output

import log "github.com/sirupsen/logrus"

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
