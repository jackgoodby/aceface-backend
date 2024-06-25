package main

import configMod "github.com/jackgoodby/aceface-backend/internal/common/config"

var config *configMod.ConfigStruct

func init() {
	config = configMod.GetConfig()
}
