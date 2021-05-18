package config

import "os"

type Config struct {
	GcpProjectId string
}

func NewConfig() Config {
	var conf Config

	conf.GcpProjectId = os.Getenv("GCP_PROJECT")

	return conf
}
