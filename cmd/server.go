/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nintran52/mypetgo/cmd/internal/api"
	"github.com/nintran52/mypetgo/cmd/internal/api/router"
	"github.com/nintran52/mypetgo/cmd/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	probeFlag   string = "probe"
	migrateFlag string = "migrate"
	seeFlag     string = "seed"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Long: `Start the stateless RESTful JSON server
Requite configuration through ENV and
a fully migrated PostgresSQL database.`,
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
		})
		e.Logger.Fatal(e.Start(":1323"))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func runServer() {
	config := config.DefaultServerConfigFromEnv()
	zerolog.TimeFieldFormat = time.RFC3339Nano
	s := api.NewServer(config)
	router.Init(s)

	go func() {
		if err := s.Start(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Info().Msg("Server closed")
			} else {
				log.Fatal().Err(err).Msg("Failed to start server")
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
