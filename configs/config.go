package configs

var (
	OpenaiServerConfigInfo *OpenaiServerConfig = &OpenaiServerConfig{}
)

type OpenaiServerConfig struct {
	DbConfigInfo DbConfig `mapstructure:"db" json:"db"`
}

type DbConfig struct {
	Host                  string `mapstructure:"host" json:"host"`
	Username              string `mapstructure:"username" json:"username"`
	Password              string `mapstructure:"password" json:"password"`
	Datebase              string `mapstructure:"datebase" json:"datebase"`
	MaxIdleConnections    string `mapstructure:"max-idle-connections" json:"max-idle-connections"`
	MaxOpenConnections    string `mapstructure:"max-open-connections" json:"max-open-connections"`
	MaxConnectionLifeTime string `mapstructure:"max-connection-life-time" json:"max-connection-life-time"`
	LogLevel              string `mapstructure:"log-level" json:"log-level"`
}

// 	 host: 127.0.0.1  # MySQL 机器 IP 和端口，默认 127.0.0.1:3306
//   username: openai # MySQL 用户名(建议授权最小权限集)
//   password: openai1234 # MySQL 用户密码
//   database: openai # miniblog 系统所用的数据库名
//   max-idle-connections: 100 # MySQL 最大空闲连接数，默认 100
//   max-open-connections: 100 # MySQL 最大打开的连接数，默认 100
//   max-connection-life-time: 10s # 空闲连接最大存活时间，默认 10s
//   log-level: 4 # GORM log level, 1: silent, 2:error, 3:warn, 4:info
