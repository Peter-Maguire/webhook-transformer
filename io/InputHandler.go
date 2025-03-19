package io

import (
	"webhook-transformer/config"
)

type InputHandler interface {
	Initialise()
	SetupInput(input config.WebhookIOConfig, outputs []OutputFunc)
}
