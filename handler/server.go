package handler

import (
	"live-code-farmer-ii/manager"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	usecaseManager manager.UsecaseManager
	srv            *gin.Engine
}

func (s *server) Run() {
	s.srv.Use(LoggerMiddleware())
	NewFarmersHandler(s.srv, s.usecaseManager.GetFarmersUsecase())
	NewPlantsHandler(s.srv, s.usecaseManager.GetPlantsUsecase())
	NewFertilizerHandler(s.srv, s.usecaseManager.GetFertilizerUsecase())
	NewFertilizerPricesHandler(s.srv, s.usecaseManager.GetFertilizerPricesUsecase())
	NewTransactionHandler(s.srv, s.usecaseManager.GetTransactionUsecase())

	s.srv.Run()
}

func NewServer() Server {
	infra := manager.NewInfraManager()
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)

	srv := gin.Default()

	return &server{
		usecaseManager: usecase,
		srv:            srv,
	}

}
