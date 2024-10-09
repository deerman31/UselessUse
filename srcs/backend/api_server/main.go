package main

import (
	"backendUselessUse/routing"
	"os"
)

func main() {
	router := routing.NewRouter()

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8000" // デフォルトポート
	}
	router.Logger.Fatal(router.Start(":" + port))
}
