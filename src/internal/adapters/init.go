package adapters

import (
	configMod "github.com/jackgoodby/aceface-backend/internal/common/config"
	"go.uber.org/zap"
)

var logger *zap.Logger
var config *configMod.ConfigStruct

func init() {
	logger = zap.NewExample()
	defer logger.Sync()

	config = configMod.GetConfig()
}
