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

