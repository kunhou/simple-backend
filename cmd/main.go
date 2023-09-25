package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github/kunhou/simple-backend/migrations"
	"github/kunhou/simple-backend/pkg/config"
	"github/kunhou/simple-backend/pkg/constant"
	"github/kunhou/simple-backend/pkg/servmanager"
	gserv "github/kunhou/simple-backend/pkg/servmanager/grpc"
	"github/kunhou/simple-backend/pkg/servmanager/http"
)

var (
	Name = "simple-backend"
	// Version is the version of the project
	Version = "0.0.0"
	// GitCommitSha is the git short commit revision number
	GitCommitSha = "-"
	// BuildDate is the build time of the project
	BuildDate = "-"
)

func Main() {
	cliApp.Run(os.Args)
}

func runApp() {
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	app, cleanup, err := initApplication(
		cfg.Debug, &cfg.Server, &cfg.Data.Database)
	if err != nil {
		panic(err)
	}
	constant.Version = Version
	constant.GitCommitSha = GitCommitSha
	constant.BuildDate = BuildDate

	fmt.Println("app Version:", constant.Version, " GitCommitSha:", constant.GitCommitSha, " BuildDate:", constant.BuildDate)

	if err := migrations.Migrate(cfg.Debug, &cfg.Data.Database); err != nil {
		panic(err)
	}

	defer cleanup()
	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}

func newApplication(serverConf *config.Server, engine *gin.Engine, gServer *grpc.Server) *servmanager.Application {
	return servmanager.NewApp(
		servmanager.WithServer(http.NewServer(engine, &serverConf.HTTP)),
		servmanager.WithServer(gserv.NewServer(gServer, &serverConf.GRPC)),
	)
}
