package io

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"strconv"
	"webhook-transformer/config"
)

type InputHTTP struct {
	e *echo.Echo
}

func (i *InputHTTP) SetupInput(input config.WebhookIOConfig, outputs []OutputFunc) {
	method := input.Data["method"]
	path := input.Data["path"]
	returnCode, err := strconv.ParseInt(input.Data["return_code"], 10, 32)

	bodyType := input.Data["body_type"]

	if err != nil {
		returnCode = http.StatusNoContent
	}

	fmt.Printf("Adding route %s %s\n", method, path)
	i.e.Add(method, path, func(c echo.Context) error {
		data := map[string]interface{}{}
		bodyBytes, err := io.ReadAll(c.Request().Body)
		if err != nil {
			fmt.Println(err)
		}

		switch bodyType {
		case "json":
			err = json.Unmarshal(bodyBytes, &data)
		case "xml":
			err = xml.Unmarshal(bodyBytes, &data)
		case "raw":
			break
		}

		if err != nil {
			fmt.Println(err)
		}

		data["body_type"] = bodyType
		data["body_raw"] = string(bodyBytes)
		data["url"] = c.Request().URL.String()
		data["headers"] = c.Request().Header

		for _, out := range outputs {
			out(input, data)
		}

		return c.NoContent(int(returnCode))
	})
}

func (i *InputHTTP) Initialise() {
	// Ignore subsequent initialisations
	if i.e != nil {
		return
	}

	i.e = echo.New()

	go func() {
		i.e.Logger.Fatal(i.e.Start(":1323"))
	}()
}
