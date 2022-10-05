package app

import "os"

type Config struct {
	AppEnv   string
	AppDebug string
	AppName  string
}

func (c *Config) IsStaging() bool { return c.AppEnv == "development" }
func (c *Config) IsProd() bool    { return c.AppEnv == "production" }
func (c *Config) IsDebug() bool   { return c.AppDebug == "True" }

func InitConfig() *Config {
	c := Config{}
	c.AppEnv = os.Getenv("APP_ENV")
	c.AppDebug = os.Getenv("APP_DEBUG")
	c.AppName = os.Getenv("APP_NAME")

	return &c
}
