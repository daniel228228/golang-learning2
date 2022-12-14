package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"syscall"

	"http_sample/internal/app"
	"http_sample/internal/config"
	"http_sample/internal/errset"
	"http_sample/internal/logger"
)

func main() {
	envFiles := flag.String("conf", "../.env", "env files separated by semicolons for configuration setup")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config, ignoredEnvFileList, err := config.LoadConfig(*envFiles)
	if err != nil {
		panic(fmt.Sprintf("loading config error: %s", err.Error()))
	}

	var log logger.Logger
	var logFile *os.File

	if config.Log.WriteFile {
		logFile = logger.InitLogFile(config.Log.File)
		defer logFile.Close()
	}

	log = logger.NewLogger(logger.GetConfigParams(config), logFile)

	log.Print("starting app")
	defer log.Print("stopped app")

	if ignoredEnvFileList != nil {
		log.Debug("the following configuration files were ignored: ", strings.Join(ignoredEnvFileList, "; "))
	}

	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("panic: %v\n\n%s", r, string(debug.Stack()))
			cancel()
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	go func() {
		var info any

		select {
		case <-ctx.Done():
			info = ctx.Err()
		case s := <-sig:
			info = s.String()
		case <-errset.Ch:
			info = "error"
		}

		log.Printf("%v received: stopping app", info)

		cancel()
	}()

	app.Run(ctx, log, config)

	if err := errset.Error(); err != nil {
		log.Error(err)
	}
}
