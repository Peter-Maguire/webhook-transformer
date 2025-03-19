package config

type Config struct {
	Webhooks []WebhookConfig `yaml:"webhooks"`
}

type WebhookConfig struct {
	Inputs  []WebhookIOConfig `yaml:"input"`
	Outputs []WebhookIOConfig `yaml:"output"`
}

type WebhookIOConfig struct {
	Type string            `yaml:"type"`
	Data map[string]string `yaml:"data,omitempty"`
}
