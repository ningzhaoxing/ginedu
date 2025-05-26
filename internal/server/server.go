package server

import (
	"context"
	"fmt"
	"gindeu/initialize"
	"gindeu/pkg/globals"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	initialize.AppInit()
	fmt.Printf("App init Evn :%s \r\n", globals.Env)
	InitRouter(globals.E)

	svr := &http.Server{
		Addr:    fmt.Sprint(globals.C.App.Host, ":", globals.C.App.Port),
		Handler: globals.E,
	}

	go func() {
		if err := svr.ListenAndServe(); err != nil {
			fmt.Println("server-->", err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	<-sigs
	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := svr.Shutdown(ctx); err != nil {
		fmt.Printf("stutdown err %v \n", err)
	}
	fmt.Println("shutdown-->ok")
}
