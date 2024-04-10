package main

type Config struct {
	Webhooks []WebhookConfig `yaml:"webhooks"`
}

type WebhookConfig struct {
	Route string `yaml:"route"`
	// Allow multiple routes with the same name but different types?
	Input  WebhookInputConfig  `yaml:"input"`
	Output WebhookOutputConfig `yaml:"output"`
	// Multiple input/outputs?
}

type WebhookInputConfig struct {
	Type       string `yaml:"type"`
	Method     string `yaml:"method,omitempty"`
	ReturnCode int    `yaml:"return-code,omitempty"`
	ReturnBody string `yaml:"return-body,omitempty"`
	// SMTP/other non-http protocols?
}

type WebhookOutputConfig struct {
	Type   string `yaml:"type"`
	URL    string `yaml:"url"`
	Method string `yaml:"method"`
	Body   string `yaml:"body,omitempty"`
	// success criteria?
	// backoff/retry?
}
