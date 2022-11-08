package service

import (
	"github.com/sirupsen/logrus"

	"github.com/truekupo/cluster/account-srv/account"
	"github.com/truekupo/cluster/account-srv/lib/config"
	"github.com/truekupo/cluster/account-srv/server"
	"github.com/truekupo/cluster/common/execute"

	"github.com/truekupo/cluster/lib/logger"
)

type service struct {
	cfg *config.Config
	gs  *server.GrpcServer
}

var (
	log *logrus.Entry = nil
)

var (
	Name    string = "account-srv"
	Version string = ""
)

func NewService(conf *config.Config, stderr bool) execute.Service {
	// Init logger
	logger.InitLogger(conf.System.LogLevel, conf.System.LogFormat, conf.System.LogFile, stderr, logrus.Fields{})

	s := service{
		cfg: conf,
	}

	log = logger.LogModule(s.Name())

	log.Debug("service:New")

	return &s
}

func (s *service) Name() string {
	return Name
}

func (s *service) Version() string {
	return Version
}

func (s *service) Init() error {
	log.WithField("version", Version).Info("service:Init")

	wh, err := account.NewAccountHandler(s.cfg)
	if err != nil {
		log.Error("service:Init", err.Error())
		return err
	}

	gs, err := server.NewGrpcServer(s.cfg, wh)
	if err != nil {
		log.Error("service:Init", err.Error())
		return err
	}

	s.gs = gs

	return nil
}

func (s *service) Start() error {
	log.Debug("service:Start")

	s.gs.Start()

	return nil
}

func (s *service) Stop() error {
	log.Debug("service:Stop")

	s.gs.Stop()

	return nil
}
