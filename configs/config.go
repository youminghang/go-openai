// Copyright 2022 Innkeeper youminghang(游铭杭) <youminghang123@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/youminghang/go-openai.

package configs

var (
	OpenaiServerConfigInfo *OpenaiServerConfig = &OpenaiServerConfig{}
)

type OpenaiServerConfig struct {
	DbConfigInfo  DbConfig  `mapstructure:"db" json:"db"`
	LogConfigInfo LogConfig `mapstructure:"log" json:"log"`
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

type LogConfig struct {
	DisableCaller     bool     `mapstructure:"disable_caller" json:"disable_caller"`
	DisableStacktrace bool     `mapstructure:"disable_stacktrace" json:"disable_stacktrace"`
	Level             string   `mapstructure:"level" json:"level"`
	Format            string   `mapstructure:"format" json:"format"`
	OutputPaths       []string `mapstructure:"output_paths" json:"output_paths"`
}
