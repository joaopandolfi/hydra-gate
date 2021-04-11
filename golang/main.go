package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hydra_gate/config"
	"hydra_gate/utils/logger"
	"hydra_gate/web/router"
	"hydra_gate/web/server"
	"hydra_gate/web/socket"

	"github.com/gorilla/mux"
)

func configInit() {
	config.Load(config.LoadFromFile(""))
}

func resilient() {
	logger.Info("[SERVER] - Shutdown")

	if err := recover(); err != nil {
		logger.Error("[SERVER] - Returning from the dark", err)
		main()
	}
}

func gracefullShutdown() {

}

func welcome() {
	// https://patorjk.com/software/taag/#p=display&f=Slant&t=ms%20-%20calendar
	fmt.Println("     __              __                          __     ")
	fmt.Println("    / /_  __  ______/ /________ _   ____ _____ _/ /____ ")
	fmt.Println("   / __ \\/ / / / __  / ___/ __ `/  / __ `/ __ `/ __/ _ \\")
	fmt.Println("  / / / / /_/ / /_/ / /  / /_/ /  / /_/ / /_/ / /_/  __/")
	fmt.Println(" /_/ /_/\\__, /\\__,_/_/   \\__,_/   \\__, /\\__,_/\\__/\\___/ ")
	fmt.Println("       /____/                    /____/                 ")

	fmt.Println("")
}

func main() {
	defer resilient()

	welcome()

	//Init
	configInit()

	// Initialize Mux Router
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	srv := server.New(r, config.Get())
	skt := socket.New()
	nr := router.New(srv)
	nr.HandleSocket(skt)

	nr.Setup()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go srv.Start()
	//skt.Start()
	<-done
	logger.Info("[SERVER] Gracefully shutdown")
	gracefullShutdown()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server Shutdown Failed", err.Error())
	}
	cancel()
}
