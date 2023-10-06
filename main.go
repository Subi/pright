package main

import (
	"log"
	api "pright/API"
	database "pright/Database"
	handler "pright/Handler"
	monitor "pright/Monitor"
)

func main() {
	newDB := database.NewDb()
	handler := handler.New(newDB)
	api := api.NewApi()
	monitor := monitor.NewMonitor(api, handler)
	log.Println("Initializing...")
	monitor.Init()
}
