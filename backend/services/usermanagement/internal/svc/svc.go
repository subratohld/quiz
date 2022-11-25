package svc

import (
	"fmt"
	"log"
	"net"

	"github.com/subratohld/quiz/cmnlib/logger"
	"github.com/subratohld/quiz/usermanagement/internal/common/config"
	"github.com/subratohld/quiz/usermanagement/internal/controllers"
	"github.com/subratohld/quiz/usermanagement/internal/databases"
	"github.com/subratohld/quiz/usermanagement/internal/repositories"
	"github.com/subratohld/quiz/usermanagement/internal/services"
	umpb "github.com/subratohld/quiz/usermanagement/internal/umproto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Service struct {
	controllerManager *controllers.ControllerManager
}

func (s *Service) Init() {
	conf := config.InitAndGetConfig()
	if conf == nil {
		logger.Logger().Fatal("could not initialize config")
	}

	initialFields := map[string]interface{}{
		"service": conf.Service,
		"env":     conf.Env,
	}
	err := logger.InitLogger(
		logger.InitLoggerWithLevelOption(logger.LogLevel(conf.Logging.Level)),
		logger.InitLoggerWithInitialFieldsOption(initialFields),
	)
	if err != nil {
		log.Fatal("could not initialize logger. err: ", err)
	}

	dbManager := databases.NewDBManager(conf)
	if dbManager == nil {
		logger.Logger().Fatal("could not initialize db")
	}

	repoManager := repositories.NewRepoManager(dbManager)

	svcManager := services.NewServiceManager(repoManager)

	controllerManager := controllers.NewControllerManager(svcManager)

	s.controllerManager = controllerManager
}

func (s *Service) Start() {
	port := config.GetConfig().GetPort(config.GetConfig().Service)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Logger().Panic(fmt.Sprintf("could not start service at port '%d'", port), zap.Error(err))
	}

	server := grpc.NewServer()

	umpb.RegisterUsermanagementServer(server, s.controllerManager)

	logger.Logger().Sugar().Debugf("service started at port '%d'", port)

	if err = server.Serve(lis); err != nil {
		logger.Logger().Panic(fmt.Sprintf("could not start service at port '%d'", port), zap.Error(err))
	}
}
