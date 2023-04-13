package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/DV-Lab/zuni-backend/api"
	db "github.com/DV-Lab/zuni-backend/db/sqlc"
	"github.com/DV-Lab/zuni-backend/util"
	_ "github.com/lib/pq"
)

func main() {
	// fmt.Println("Hello world")
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}


	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the db...:", err)
	}

	defer conn.Close()

	runDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot initialize the server...", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot connect to the server...", err)
	}
}


func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up")
	}

	fmt.Println("db migrated successfully")
}
