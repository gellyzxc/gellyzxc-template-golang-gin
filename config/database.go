package config

import (
	"context"
	"database/sql"
	"log"
	"os"

	"gellyzxc-template-golang-gin/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var DB *bun.DB

func ConnectDB() {
	dsn := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") +
		"@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?sslmode=disable"

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	DB = bun.NewDB(sqldb, pgdialect.New())

	// Проверка соединения
	if err := DB.Ping(); err != nil {
		log.Fatalf("connect db error: %v", err)
	}

	log.Println("db ok")

	Migrate()
}

func Migrate() {
	ctx := context.Background()

	m := []interface{}{
		(*models.User)(nil),
		(*models.Post)(nil),
	}

	for _, model := range m {
		_, err := DB.NewCreateTable().Model(model).IfNotExists().Exec(ctx)
		if err != nil {
			log.Fatalf("migration error: %v", err)
		}
	}
	log.Println("done migrations")
}
