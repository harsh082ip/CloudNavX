package main

import (
	"github.com/harsh082ip/CloudNavX/internal/config"
	"github.com/harsh082ip/CloudNavX/internal/router"
)

const (
	WEBPORT = ":8000"
)

func main() {
	config.LoadConfig()
	r := router.SetupRouter()
	r.Run(WEBPORT)
}
