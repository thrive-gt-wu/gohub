package config

import (
	"gohub/pkg/config"

)
func init() {
	config.Add("database", func() map[string]interface{} {
		// 默认数据库
		return map[string]interface{} {
			"connection": config.Env("DB_CONNECTION", "mysql"),

			"mysql": map[string]interface{} {
				// 连接数据库
				"host": config.Env("DB_HOST", "127.0.0.1"),
				"port": config.Env("DB_PORT", "3306"),
				"database": config.Env("DB_DATABASE", "gohub"),
				"username": config.Env("DB_USERNAME", ""),
				"password": config.Env("DB_PASSWORD", ""),
				"charset": "utf8mb4",

				// 连接池配置
				"max_idle_connection": config.Env("DB_MAX_IDLE_CONNECTION", 100),
				"max_open_connection": config.Env("DB_MAX_OPEN_CONNECTION", 25),
				"max_life_second": config.Env("DB_MAX_LIFE_SECOND", 5*60),
			},

			"sqlie": map[string]interface{}{
				"database":config.Env("DB_SQL_FILE", "database/database.db"),
			},

		}
	})
}