package main

import (
	"flag"
	"net/http"

	"github.com/MonkeyBuisness/family-api/controller"
	"github.com/MonkeyBuisness/family-api/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	addr := flag.String("addr", "127.0.0.1", "network address")
	flag.Parse()
	service := controller.NewService()
	h := handler.New(service)
	logrus.Infof("service started at: %s", *addr)
	logrus.Panic(http.ListenAndServe(*addr, h))
}
