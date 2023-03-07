package main

import (
	"fmt"
	"sync"

	"github.com/AdanJSuarez/maxmind/internal/app"
	"github.com/AdanJSuarez/maxmind/internal/configuration"
)

func main() {
	config := configuration.New()
	if err := config.CheckConfiguration(); err != nil {
		return
	}
	var wg sync.WaitGroup

	app, err := app.New(&wg, config)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer app.Close()

	fmt.Println("==> Start <==")
	app.Start()
	fmt.Println("==> Finished <== ")
}
