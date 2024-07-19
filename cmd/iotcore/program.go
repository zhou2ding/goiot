package main

import (
	_ "goiot/internal/pkg/database"
	"goiot/internal/pkg/logger"
	"goiot/internal/pkg/mq"
	"goiot/internal/pkg/service"
)

type Program struct {
}

func (p *Program) Start(s service.Service) error {
	logger.Logger.Warnf("starting programme...")

	if err := mq.InitSqsClient(); err != nil {
		logger.Logger.Warnf("init sqs client error: %v", err)
		_ = p.Stop(s)
	}

	return nil
}

func (p *Program) Stop(s service.Service) (e error) {
	logger.Logger.Warnf("stoping programme...")
	return nil
}
