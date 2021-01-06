package main

import (
	"context"
	"flag"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"grocery/config"
	"grocery/internal/db"
	"grocery/internal/server"
	"grocery/internal/server/api"
	"grocery/pkg/version"
	"os"
	"os/signal"
	"syscall"
)

var (
	fGenerate = flag.String("generate", "", "generate all default configs to specified path")
	fMigrate  = flag.Bool("migrate", false, "Run db migration(with gorm's AutoMigrate)")
)

func fatalIf(err error) {
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}

func initDb(c config.Conf) (dbManager *db.Manager) {
	dbManager = &db.Manager{}
	err := dbManager.Connect(c)
	fatalIf(err)
	return dbManager
}

func initHandler(c config.Conf, dbManager *db.Manager) (*api.Handler, error) {
	handler, err := api.NewHandler(dbManager, c.Debug)
	if err != nil {
		return nil, err
	}
	return handler, nil
}

func main() {
	flag.Parse()

	if *fGenerate != "" {
		config.GenerateConfigs(*fGenerate)
		return
	}

	version.Init()

	appConf, err := config.AppConfig()
	fatalIf(err)

	initLog(appConf)

	dbManager := initDb(appConf)

	if *fMigrate {
		err := dbManager.Migrate()
		fatalIf(err)
		return
	}

	handler, err := initHandler(appConf, dbManager)
	fatalIf(err)
	defer handler.WaitAndClose()

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		quit := make(chan os.Signal, 16)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-quit
		cancel()
	}()

	conf := server.Config{Addr: appConf.ListenAddr, Debug: appConf.Debug}
	log.Info().Bool("debug", appConf.Debug).Msg("Grocery start")
	log.Fatal().Err(server.ListenAndServe(ctx, conf, handler)).Msg("app exits")
}

func initLog(c config.Conf) {
	log.Logger = log.Output(os.Stdout)

	if c.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
