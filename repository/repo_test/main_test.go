package repotest

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"gorm.io/gorm"

	"github/kunhou/simple-backend/migrations"
	"github/kunhou/simple-backend/pkg/data"
)

type TestDBSetting struct {
	Driver       string
	ImageName    string
	ImageVersion string
	ENV          []string
	PortID       string
	Connection   string
}

var (
	postgresDBSetting = TestDBSetting{
		Driver:       "postgres",
		ImageName:    "postgres",
		ImageVersion: "16.0-alpine",
		ENV:          []string{"POSTGRES_USER=root", "POSTGRES_PASSWORD=root", "POSTGRES_DB=simple", "LISTEN_ADDRESSES='*'"},
		PortID:       "5432/tcp",
		Connection:   "host=localhost port=%s user=root password=root dbname=simple sslmode=disable",
	}
)

var testDB *gorm.DB

func TestMain(t *testing.M) {
	cleanup := initTestData(postgresDBSetting)
	defer func() {
		if cleanup != nil {
			cleanup()
		}
	}()

	t.Run()
}

func initTestData(dbSetting TestDBSetting) (cleanup func()) {
	connection, cleanup, err := initDatabaseImage(dbSetting)
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	config := &data.DatabaseConf{}
	config.SetConnection(connection)
	testDB = data.NewDB(true, config)

	if err := initDatabase(testDB); err != nil {
		log.Fatalf("Could not run migrations: %s", err)
	}

	return cleanup
}

func initDatabaseImage(dbSetting TestDBSetting) (connection string, cleanup func(), err error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: postgresDBSetting.ImageName,
		Tag:        postgresDBSetting.ImageVersion,
		Env:        postgresDBSetting.ENV,
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		return "", nil, fmt.Errorf("could not start resource: %s", err)
	}

	connection = fmt.Sprintf(dbSetting.Connection, resource.GetPort(dbSetting.PortID))
	if err := pool.Retry(func() error {
		db, err := sql.Open(dbSetting.Driver, connection)
		if err != nil {
			return err
		}

		defer db.Close()
		if err := db.Ping(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return "", nil, fmt.Errorf("could not connect to database: %s", err)
	}

	return connection, func() { _ = pool.Purge(resource) }, nil
}

func initDatabase(db *gorm.DB) error {
	return migrations.NewMigrator(context.Background(), db).Migrate()
}
