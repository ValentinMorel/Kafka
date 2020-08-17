package main

import (
	"kafka/apihandlers"

	"github.com/labstack/echo"
)

func main() {

	// Define the routes where we will perform POST requests.
	APIinstance := echo.New()
	APIinstance.POST("/producer/sync", apihandlers.SendSyncMsg)

	// Start server on port 8888 and listen.
	APIinstance.Logger.Fatal(APIinstance.Start(":8888"))
}
