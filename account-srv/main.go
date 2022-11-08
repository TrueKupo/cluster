package main

import (
	"flag"
	"fmt"

	"github.com/truekupo/cluster/account-srv/lib/config"
	"github.com/truekupo/cluster/account-srv/service"
	"github.com/truekupo/cluster/common/execute"
)

var (
	configFile = flag.String("config", "config.yaml", "/path/to/config.yaml")
	stderr     = flag.Bool("stderr", false, "[true|false] - enabled debug to console")
)

func main() {
	flag.Parse()

	// Load config
	conf, err := config.Parse(*configFile)
	if err != nil {
		panic(err)
	}

	fmt.Println(conf)

	// Start service
	s := service.NewService(conf, *stderr)
	err = execute.StartService(s)
	if err != nil {
		panic(err)
	}

	<-make(chan int)
}
