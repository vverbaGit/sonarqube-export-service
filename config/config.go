package config

import "os"

type SonarConfig struct {
	BaseUrl string
	Token   string
}

func Load() *SonarConfig {
	c := SonarConfig{}
	c.BaseUrl = os.Getenv("SONAR_URL")
	c.Token = os.Getenv("SONAR_TOKEN")
	return &c
}
