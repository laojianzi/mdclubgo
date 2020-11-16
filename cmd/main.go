package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"github.com/laojianzi/mdclubgo/api"
	"github.com/laojianzi/mdclubgo/cache"
	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/db"
	"github.com/laojianzi/mdclubgo/email"
	"github.com/laojianzi/mdclubgo/log"
)

var customConf string

func main() {
	rootCmd := &cobra.Command{Use: "mdclubgo"}
	apiCmd := &cobra.Command{
		Use:   "api",
		Short: "Run http server for api",
		Run:   run,
	}

	apiCmd.Flags().StringVarP(&customConf, "config-file", "f",
		filepath.Join(conf.CustomDir(), "conf", "app.ini"),
		"input your custom config file path")

	rootCmd.AddCommand(apiCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(fmt.Errorf("api cmd execute: %w", err).Error())
	}
}

func run(cmd *cobra.Command, args []string) {
	if err := conf.Init(customConf); err != nil {
		log.Fatal(fmt.Sprintf("failed to initialize application: %v", err))
	}

	db.Init()
	cache.Init()
	email.Init()

	addr := fmt.Sprintf("%s:%s", conf.Server.HTTPAddr, conf.Server.HTTPPort)
	go func() {
		if err := api.Server().Start(addr); err != nil && !errors.Is(err,
			http.ErrServerClosed) {
			log.Fatal("api start error: %s", err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := api.Server().Shutdown(ctx); err != nil {
		log.Fatal("api server shutdown error: %s", err.Error())
	}
	cancel()

	cache.Close()
	db.Close()
	log.Close()
}
