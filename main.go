package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/akrylysov/algnhsa"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/xonha/gatth/api"
	"github.com/xonha/gatth/app"
)

//go:embed all:public
var embeddedFiles embed.FS

func main() {
	config := huma.DefaultConfig("Docs Example", "1.0.0")
	config.DocsPath = ""

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	humaApi := humachi.New(router, config)

	assets, _ := fs.Sub(embeddedFiles, "public")
	router.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.FS(assets))))

	router.Get("/docs", api.Docs)
	router.Get("/", templ.Handler(app.Page()).ServeHTTP)

	huma.Get(humaApi, "/hello/{name}", api.Hello)

	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		algnhsa.ListenAndServe(router, nil)
	} else {
		fmt.Println("Server starting locally on port 3000...")
		http.ListenAndServe(":3000", router)
	}
}
