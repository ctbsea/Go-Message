package gateway

import (
	"github.com/ctbsea/Go-Message/config"
	"github.com/iris-contrib/middleware/newrelic"
	"log"
)

func Newrelic(config config.Config) *newrelic.Newrelic {
	conf := newrelic.Config("APP_SERVER_NAME", "NEWRELIC_LICENSE_KEY")
	conf.Enabled = true
	m, err := newrelic.New(conf)
	if err != nil {
		log.Fatal(err)
	}
	return m
}
