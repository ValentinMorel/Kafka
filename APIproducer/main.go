package main

import (
	"kafka/APIhandlers"

	"github.com/labstack/echo"
)

func main() {

	APIinstance := echo.New()
	APIinstance.POST("/producer/sync", APIhandlers.SendSyncMsg)
	APIinstance.Logger.Fatal(APIinstance.Start(":8888"))
}
