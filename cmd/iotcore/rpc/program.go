package main

import (
	_ "goiot/internal/pkg/cache"
	_ "goiot/internal/pkg/database"
	"goiot/internal/pkg/logger"
	"goiot/internal/pkg/mq"
	"goiot/internal/pkg/oss"
	"goiot/internal/pkg/push"
	"goiot/internal/pkg/service"
	"goiot/internal/pkg/trans"
)

type Program struct {
}

func (p *Program) Start(s service.Service) error {
	logger.Logger.Warnf("starting programme...")

	if err := oss.InitS3Client(); err != nil {
		logger.Logger.Warnf("init s3 client error: %v", err)
		_ = p.Stop(s)
		return err
	}

	if err := push.InitSnsClient(); err != nil {
		logger.Logger.Warnf("init sns client error: %v", err)
		_ = p.Stop(s)
		return err
	}

	if err := push.InitSESClient(); err != nil {
		logger.Logger.Warnf("init ses client error: %v", err)
		_ = p.Stop(s)
		return err
	}

	if err := mq.InitSqsClient(); err != nil {
		logger.Logger.Warnf("init sqs client error: %v", err)
		_ = p.Stop(s)
		return err
	}

	if err := trans.InitTranslatorOfValidator("en"); err != nil {
		logger.Logger.Warnf("init translator of validator error: %v", err)
		_ = p.Stop(s)
		return err
	}

	return nil
}

func (p *Program) Stop(s service.Service) (e error) {
	logger.Logger.Warnf("stoping programme...")
	return nil
}
