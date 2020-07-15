package main

import (
	"goldnoti/api"
	"goldnoti/service"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

func main() {
	service.SetupConfig("./config")
	service.SetupConfigParams()

	http.HandleFunc("/api/health", api.Health)
	http.HandleFunc("/api/today", api.GetTodayPrice)

	port := os.Getenv("PORT")
	if port == "" {
		port = viper.GetString("Listening.Port")
	}
	http.ListenAndServe(":"+port, nil)
}
