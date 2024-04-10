package main

import (
    "bytes"
    "fmt"
    "github.com/labstack/echo/v4"
    "gopkg.in/yaml.v3"
    "io"
    "net/http"
    "os"
)

type WebhookTransformer struct {
    webhooks []WebhookConfig
    client   http.Client
}

func NewWebhookTransformer() WebhookTransformer {
    yamlFile, err := os.ReadFile("config.yaml")
    if err != nil {
        fmt.Println(err)
    }
    yamlConfig := Config{}
    err = yaml.Unmarshal(yamlFile, &yamlConfig)
    if err != nil {
        fmt.Println(err)
    }
    return WebhookTransformer{
        webhooks: yamlConfig.Webhooks,
    }
}

func (wt *WebhookTransformer) SetupEcho(e *echo.Echo) {

    for _, config := range wt.webhooks {
        fmt.Printf("Adding route %s %s\n", config.Input.Method, config.Route)
        if config.Input.ReturnCode == 0 {
            config.Input.ReturnCode = 204
        }
        e.Add(config.Input.Method, config.Route, func(c echo.Context) error {
            var body io.Reader

            if config.Output.Body == "" {
                body = c.Request().Body
            } else {
                body = bytes.NewBuffer([]byte(config.Output.Body))
            }

            request, err := http.NewRequest(config.Output.Method, config.Output.URL, body)
            if err != nil {
                fmt.Println(err)
                return c.JSON(500, err)
            }

            if config.Output.Body == "" {
                request.Header.Add("content-type", c.Request().Header.Get("content-type"))
            }

            res, err := wt.client.Do(request)

            output, _ := io.ReadAll(res.Body)
            fmt.Println(string(output))

            if config.Input.ReturnBody == "" {
                return c.NoContent(config.Input.ReturnCode)
            }

            return c.String(config.Input.ReturnCode, config.Input.ReturnBody)
        })
    }

}

func main() {

    webhookTransformer := NewWebhookTransformer()

    e := echo.New()

    webhookTransformer.SetupEcho(e)

    e.Logger.Fatal(e.Start(":1323"))
}
