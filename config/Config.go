package config

import "strconv"

type Config struct {
	Webhooks []WebhookConfig `yaml:"webhooks"`
}

type WebhookConfig struct {
	Inputs  []WebhookIOConfig `yaml:"input"`
	Outputs []WebhookIOConfig `yaml:"output"`
}

type WebhookIOConfig struct {
	Type string    `yaml:"type"`
	Data ConfigMap `yaml:"data,omitempty"`
}

type ConfigMap map[string]any

func (cf ConfigMap) GetString(key string) string {
	str, ok := cf[key].(string)
	if !ok {
		return ""
	}
	return str
}

func (cf ConfigMap) GetInt(key string) int {
	intValue, ok := cf[key].(int)
	if !ok {
		int64Value, err := strconv.ParseInt(cf[key].(string), 10, 64)
		if err != nil {
			panic(err)
		}
		intValue = int(int64Value)
	}
	return intValue
}

func (cf ConfigMap) GetBool(key string) bool {
	boolean, ok := cf[key].(bool)
	if !ok {
		return false
	}
	return boolean
}
