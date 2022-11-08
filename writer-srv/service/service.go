package service

import (
	"github.com/sirupsen/logrus"

	"github.com/truekupo/cluster/common/execute"
	"github.com/truekupo/cluster/writer-srv/lib/config"
	"github.com/truekupo/cluster/writer-srv/server"
	"github.com/truekupo/cluster/writer-srv/writer"

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
	Name    string = ""
	Version string = ""
)

func NewService(conf *config.Config, stderr bool) execute.Service {
	s := service{
		cfg: conf,
	}

	// Init logger
	logger.InitLogger(conf.System.LogLevel, conf.System.LogFormat, conf.System.LogFile, stderr, map[string]interface{}{"srv": s.Name()})

	log = logger.LogModule("service")

	log.Debug("new")

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

	wh, err := writer.NewWriterHandler(s.cfg)
	if err != nil {
		log.WithField("error", err).Error("init")
		return err
	}

	gs, err := server.NewGrpcServer(s.cfg, wh)
	if err != nil {
		log.Error("Init", err.Error())
		return err
	}

	s.gs = gs

	return nil
}

func (s *service) Start() error {
	log.Debug("start")

	s.gs.Start()

	return nil
}

func (s *service) Stop() error {
	log.Debug("stop")

	s.gs.Stop()

	return nil
}
