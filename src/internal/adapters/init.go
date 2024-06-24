package adapters

import (
	"go.uber.org/zap"

	configMod "github.com/jackgoodby/aceface-backend/internal/common/config"
)

var logger *zap.Logger
var config *configMod.ConfigStruct

func init() {
	logger = zap.NewExample()
	defer logger.Sync()

	config = configMod.GetConfig()
}
