package main

import (
	"fmt"

	"github.com/Tambarie/eCommmerceApp/internal/core/helper"
	"github.com/Tambarie/eCommmerceApp/internal/core/shared"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	helper.InitializeLogDir()

	//mongoURL := "mongodb://localhost:27017"
	service_address, service_port, _, _, _, _, _, _ := helper.LoadConfig()

	router := gin.Default()
	router.Use(helper.LogRequest)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404,
			helper.PrintErrorMessage("404", shared.NoResourceFound))
	})

	fmt.Println("service running on " + service_address + ":" + service_port)
	helper.LogEvent("Info", fmt.Sprintf("started platform service application on "+service_address+":"+service_port+" in "+time.Since(time.Now()).String()))
	_ = router.Run(":" + service_port)

}
