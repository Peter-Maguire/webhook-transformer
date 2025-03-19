package io

import "webhook-transformer/config"

type OutputFunc func(input config.WebhookIOConfig, data map[string]interface{})

type OutputHandler interface {
	Initialise()
	SetupOutput(output config.WebhookIOConfig) OutputFunc
}
