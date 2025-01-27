package api

import (
	"context"
	"fmt"
	"net/http"
)

type HelloInput struct {
	Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
}
type HelloOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func Hello(ctx context.Context, input *HelloInput) (*HelloOutput, error) {
	resp := &HelloOutput{}
	resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
	return resp, nil
}

func Docs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
    <!doctype html>
    <html>
      <head>
        <title>API Reference</title>
			  <link rel="icon" href="public/favicon.ico"/>
        <meta charset="utf-8" />
        <meta
          name="viewport"
          content="width=device-width, initial-scale=1" />
      </head>
      <body>
        <script
          id="api-reference"
          data-url="/openapi.json"></script>
        <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
      </body>
    </html>
  `),
	)
}
