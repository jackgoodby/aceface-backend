package main

import (
	configMod "github.com/jackgoodby/aceface-backend/internal/common/config"
	"go.uber.org/zap"
)

var config *configMod.ConfigStruct
var logger *zap.Logger

func init() {
	config = configMod.GetConfig()
	logger = zap.NewExample()
	defer logger.Sync()
}
