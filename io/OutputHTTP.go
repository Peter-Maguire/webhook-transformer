package io

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"webhook-transformer/config"
	"webhook-transformer/helper"
)

type OutputHTTP struct {
}

func (o *OutputHTTP) Initialise() {
}

func (o *OutputHTTP) SetupOutput(output config.WebhookIOConfig) OutputFunc {
	return func(input config.WebhookIOConfig, data map[string]interface{}) {
		url, _ := helper.Template(output.Data.GetString("url"), data)
		method, _ := helper.Template(output.Data.GetString("method"), data)

		bodyRaw := output.Data.GetString("body")
		forwardBody := output.Data.GetBool("forward_body")

		var bodyData io.Reader = nil

		if forwardBody {
			bodyData = strings.NewReader(data["body_raw"].(string))
		} else if bodyRaw != "" {
			body, err := helper.Template(bodyRaw, data)
			if err != nil {
				fmt.Println(err)
			}
			bodyData = strings.NewReader(body)
		}

		req, err := http.NewRequest(method, url, bodyData)
		if err != nil {
			fmt.Println(err)
			return
		}

		if forwardBody {
			dataType := data["headers"].(http.Header).Get("Content-Type")
			req.Header.Set("content-type", dataType)
		}

		_, err = http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
		}
	}
}
