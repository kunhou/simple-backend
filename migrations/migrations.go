package migrations

import (
	"context"
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"gorm.io/gorm"

	"github/kunhou/simple-backend/pkg/data"
)

//go:embed sql/postgres/*.sql
var fs embed.FS

type Migrator struct {
	ctx context.Context
	db  *gorm.DB
}

func NewMigrator(ctx context.Context, db *gorm.DB) *Migrator {
	return &Migrator{
		ctx: ctx,
		db:  db,
	}
}

func (m *Migrator) Migrate() error {
	db, err := m.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get db connection: %w", err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to get postgres driver: %w", err)
	}
	d, err := iofs.New(fs, "sql/postgres")
	if err != nil {
		return fmt.Errorf("failed to get iofs: %w", err)
	}

	mdb, err := migrate.NewWithInstance(
		"iofs", d,
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to get migrate instance: %w", err)
	}

	if err := mdb.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to migrate up: %w", err)
	}

	return nil
}

func Migrate(debug bool, dbConf *data.DatabaseConf) error {
	engine := data.NewDB(debug, dbConf)
	m := NewMigrator(context.Background(), engine)
	return m.Migrate()
}
