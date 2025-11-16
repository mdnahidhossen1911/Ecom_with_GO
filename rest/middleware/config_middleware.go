package middleware

import "ecom_project/config"

type ConfigMiddleware struct {
	config *config.Config
}

func NewConfigMiddleware(config *config.Config) *ConfigMiddleware {
	return &ConfigMiddleware{
		config: config,
	}
}

func (cm *ConfigMiddleware) GetJwtSecret() string {
	if cm == nil || cm.config == nil {
		return ""
	}
	return cm.config.JwtSecureKey
}