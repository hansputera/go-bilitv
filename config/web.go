package config

import (
	"strconv"
	"strings"
)

type WebConfig struct {
	ApiGateway string
	WebUrl     string
}

func (config *WebConfig) ApiEndpoint(version int, uri string) string {
	return config.ApiGateway + "/v" + strconv.Itoa(version) + "/" + uri
}

func (config *WebConfig) WebEndpoint(lang string, uri string) string {
	return config.WebUrl + "/" + strings.ToLower(lang) + "/" + uri
}

func GetWebConfig() *WebConfig {
	return &WebConfig{
		ApiGateway: "https://api.bilibili.tv/intl/gateway/web",
		WebUrl:     "https://bilibili.tv",
	}
}
