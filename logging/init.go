package logging

import (
	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func SetupGrayLog(gelfAddr string, appMode string) {
	hook := graylog.NewGraylogHook(gelfAddr, map[string]interface{}{
		"app-mode": appMode,
		"app-key":  RandStringRunes(6),
	})
	hook.Level = log.TraceLevel

	log.AddHook(hook)
	log.Info("aaaa")
}
