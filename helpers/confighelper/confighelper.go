package confighelper

import (
	"afford_meds_interview/models"

	"github.com/BurntSushi/toml"
)

// InitConfig - initialise app configuration
func InitConfig() error {
	_, err := toml.DecodeFile(models.AppConfig.GetConfigFilePath(), &models.AppConfig)
	if err != nil {
		return err
	}
	return nil
}
