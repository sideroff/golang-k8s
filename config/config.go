package config

import "google.com/sideroff/golang-k8s/env"

type Config struct {
	Env string
	WebServerConfig *WebServerConfig
}

type WebServerConfig struct {
	Host string
	Port string
	PublicFolderPath string
}

// func Get() *Config {
// 	return &Config {
// 		Env: env.Require("ENV"),
// 		WebServerConfig: &WebServerConfig {
// 			Host: env.Require("WEB_SERVER_HOST"),
// 			Port: env.WithDefault("WEB_SERVR_PORT", "3000"),
// 			PublicFolderPath: env.WithDefault("WEB_SERVER_PUBLIC_FOLDER_PATH", "../public"),
// 		},
// 	}
// }

// Used for development
func Get() *Config {
	return &Config {
		Env: env.WithDefault("ENV", "development"),
		WebServerConfig: &WebServerConfig {
			Host: env.WithDefault("WEB_SERVER_HOST", ""),
			Port: env.WithDefault("WEB_SERVR_PORT", "3000"),
			PublicFolderPath: env.WithDefault("WEB_SERVER_PUBLIC_FOLDER_PATH", "/public"),
		},
	}
}